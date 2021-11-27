import { ChangeDetectorRef, Injectable, NgZone } from '@angular/core';
import Web3 from 'web3';
import {
  BehaviorSubject,
  catchError,
  defer,
  distinctUntilChanged,
  from,
  fromEvent,
  map,
  Observable,
  of,
  Subject,
  switchMap,
  take,
  tap,
  timer,
  withLatestFrom,
} from 'rxjs';
import { EthAccounts } from '../model/common.model';
import { setupContract } from '../libs/contracts/contract';
import { SpinProccessor } from '../libs/contracts/SpinProccessor';
import { SPIN_PROCESSER } from '../constant/common.constant';

const ONE_ETH_TO_GWEI = 1e18;
@Injectable({
  providedIn: 'root',
})
export class SignerService {
  private _web3!: Web3;

  private _isReady = false;
  private readonly _address$ = new BehaviorSubject<string>('');
  private _contract: SpinProccessor;

  constructor(private ngZone: NgZone) {
    this._setup();
  }

  private _setup(): void {
    window.addEventListener('load', async () => {
      const accounts$ = defer(
        () => window.ethereum.send('eth_accounts') as Promise<EthAccounts>
      );
      if (!window.ethereum) return;
      this._web3 = new Web3(window.ethereum);

      this._contract = setupContract(this._web3, SPIN_PROCESSER);

      const lastAddr$ = this._address$.pipe(take(1));
      timer(0, 1000)
        .pipe(switchMap(() => accounts$.pipe(map((resp) => resp.result))))
        .pipe(
          switchMap((accounts) =>
            lastAddr$.pipe(map((addr) => [addr, accounts]))
          )
        )
        .subscribe(([addr, accounts]) => {
          if (
            Array.isArray(accounts) &&
            accounts.length &&
            addr !== accounts[0]
          ) {
            this._address$.next(accounts[0]);
          }
        });
    });
  }

  get isReady() {
    return this._isReady;
  }

  get address$() {
    return this._address$.pipe(distinctUntilChanged());
  }

  get signer() {
    return this._web3;
  }

  get contract() {
    return this._contract;
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
