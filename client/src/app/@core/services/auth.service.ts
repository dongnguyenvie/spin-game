import { Injectable } from '@angular/core';
import { SignerService } from './signer.service';
import {
  lastValueFrom,
  map,
  of,
  switchMap,
  from,
  firstValueFrom,
  concatMap,
  defer,
  take,
  takeLast,
  withLatestFrom,
  last,
  catchError,
  tap,
  BehaviorSubject,
  Observable,
} from 'rxjs';
import { API } from '../constant/common.constant';
import { HttpClient } from '@angular/common/http';
import { IUser, LogonRespose, NonceRespose } from '../model/common.model';
import { LoginModalService } from 'src/app/shared/login-modal/login-modal.service';
import { parseJwt } from '../utils/util';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly _token$ = new BehaviorSubject<string>('');

  constructor(
    private signer: SignerService,
    private http: HttpClient,
    private loginModal: LoginModalService
  ) {}

  get token$() {
    return this._token$.asObservable();
  }

  get user$() {
    return this._token$.pipe(
      map((token) => {
        if (!token) return null;
        const parse = parseJwt(token);
        return parse.payload as IUser;
      })
    );
  }

  fetchNonce$(addr: string) {
    return this.http
      .get<NonceRespose>(API.nonce.replace(':address', addr))
      .pipe(map((resp) => resp.data.nonce));
  }

  fetchLogin$({ signature, addr }: { signature: string; addr: string }) {
    return this.http.post<LogonRespose>(API.login, {
      signature,
      walletAddress: addr,
    });
  }

  register({
    signature,
    email,
    password,
    addr,
    messageSign,
  }: {
    signature: string;
    email: string;
    password: string;
    addr: string;
    messageSign: string;
  }) {
    return this.http.post(API.register, {
      signature: signature,
      email: email,
      password: password,
      walletAddress: addr,
      messageSign: messageSign,
    });
  }

  register$({
    email,
    password,
    addr,
    messageSign,
  }: {
    email: string;
    password: string;
    addr: string;
    messageSign: string;
  }) {
    const signature$ = defer(() =>
      this.signer.signer$.eth.personal.sign(messageSign, addr, '')
    );

    return signature$.pipe(
      switchMap((signature) => {
        return this.register({
          signature,
          email,
          password,
          addr,
          messageSign,
        });
      }),
      catchError(() => of(null))
    );
  }

  login$() {
    const lastAddresses$ = defer(() =>
      of(null).pipe(
        withLatestFrom(this.signer.address$),
        map(([_, lastAddr]) => lastAddr)
      )
    );

    return lastAddresses$
      .pipe(
        switchMap((addr) =>
          this.fetchNonce$(addr).pipe(
            concatMap((nonce) => {
              if (!nonce) return of(null);
              return this.signer.signer$.eth.personal.sign(
                String(nonce),
                addr,
                ''
              );
            }),
            concatMap((signature) => {
              if (!signature) return of(null);
              return this.fetchLogin$({
                signature,
                addr,
              });
            }),
            catchError((err) => {
              if (err.status === 404) {
                this.loginModal.onOpen();
              }
              return of(null);
            })
          )
        )
      )
      .pipe(
        tap((result) => {
          if (!result) return;
          console.warn('login$', result);
          this._token$.next(result.data.token);
        })
      );
  }
}
