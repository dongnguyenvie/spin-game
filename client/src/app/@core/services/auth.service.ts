import { Injectable } from '@angular/core';
import { SignerService } from './signer.service';
import { lastValueFrom, of, switchMap } from 'rxjs';
import { API } from '../constant/common.constant';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(private signer: SignerService, private http: HttpClient) {}

  getNonce() {
    return this.signer.addressAsObs.pipe(
      switchMap((addr) => {
        if (!addr) return of(null);
        return this.http.get(API.nonce.replace(':address', addr));
      })
    );
  }
}
