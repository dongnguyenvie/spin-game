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
} from 'rxjs';
import { API } from '../constant/common.constant';
import { HttpClient } from '@angular/common/http';
import { NonceRespose } from '../model/common.model';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(private signer: SignerService, private http: HttpClient) {}

  getNonce() {
    return this.signer.addressAsObs.pipe(
      switchMap((addr) => {
        if (!addr) return of(null);
        return this.http
          .get<NonceRespose>(API.nonce.replace(':address', addr))
          .pipe(map((resp) => resp.data.nonce));
      })
    );
  }

  login() {
    const lastAddresses$ = defer(() =>
      this.signer.addressAsObs.pipe(takeLast(0))
    );

    return this.getNonce().pipe(
      concatMap((nonce) => {
        if (!nonce) return of(null);
        return lastAddresses$.pipe(
          concatMap((addresses) =>
            this.signer.signer.eth.personal.sign(String(nonce), addresses, '')
          )
        );
      }),
      concatMap((signature) => {
        if (!signature) return of(null);
        return lastAddresses$.pipe(
          concatMap((addresses) =>
            this.http.post(API.login, {
              signature,
              walletAddress: addresses,
            })
          )
        );
      })
    );
  }
}
