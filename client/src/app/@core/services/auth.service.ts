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
} from 'rxjs';
import { API } from '../constant/common.constant';
import { HttpClient } from '@angular/common/http';
import { NonceRespose } from '../model/common.model';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(private signer: SignerService, private http: HttpClient) {}

  getNonce(addr: string) {
    return this.http
      .get<NonceRespose>(API.nonce.replace(':address', addr))
      .pipe(map((resp) => resp.data.nonce));
  }

  login(signature: string, addr: string) {
    return this.http.post(API.login, {
      signature,
      walletAddress: addr,
    });
  }

  login$() {
    const lastAddresses$ = defer(() =>
      of(null).pipe(
        withLatestFrom(this.signer.address$),
        map(([_, lastAddr]) => lastAddr)
      )
    );

    return lastAddresses$.pipe(
      switchMap((addr) =>
        this.getNonce(addr).pipe(
          concatMap((nonce) => {
            if (!nonce) return of(null);
            return this.signer.signer.eth.personal.sign(
              String(nonce),
              addr,
              ''
            );
          }),
          concatMap((signature) => {
            if (!signature) return of(null);
            return this.login(signature, addr);
          })
        )
      )
    );
  }
}
