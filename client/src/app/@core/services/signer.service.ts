import { ChangeDetectorRef, Injectable } from '@angular/core';
import Web3 from 'web3';
import {
  BehaviorSubject,
  catchError,
  defer,
  from,
  fromEvent,
  map,
  Observable,
  of,
  Subject,
  switchMap,
  tap,
  timer,
} from 'rxjs';
import { EthAccounts } from '../model/common.model';

const ONE_ETH_TO_GWEI = 1e18;
@Injectable({
  providedIn: 'root',
})
export class SignerService {
  private _web3!: Web3;

  private _isReady = false;
  private readonly _address$ = new BehaviorSubject<string>('');

  constructor() {
    this.setup();
  }

  private setup(): void {
    window.addEventListener('load', async () => {
      const accounts$ = defer(
        () => window.ethereum.send('eth_accounts') as Promise<EthAccounts>
      );
      if (!window.ethereum) return;
      this._web3 = new Web3(window.ethereum);

      timer(0, 1000)
        .pipe(switchMap(() => accounts$.pipe(map((resp) => resp.result))))
        .subscribe((accounts) => {
          if (Array.isArray(accounts) && accounts.length) {
            this._address$.next(accounts[0]);
          }
        });
    });
  }

  get isReady() {
    return this._isReady;
  }

  get address$() {
    return this._address$.asObservable();
  }

  get signer$() {
    return this._web3;
  }

  connectWallet() {
    const account$ = from<Promise<EthAccounts>>(
      window.ethereum.send('eth_requestAccounts')
    );

    return account$.pipe(
      map((resp) => resp.result),
      tap((accounts) => {
        if (Array.isArray(accounts) && accounts.length) {
          this._address$.next(accounts[0]);
        }
      }),
      catchError(() => of(null))
    );
  }
}

declare global {
  interface Window {
    ethereum: any;
    web3: Web3;
  }
}
