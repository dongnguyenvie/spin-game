import { Injectable } from '@angular/core';
import Web3 from 'web3';
import { BehaviorSubject, fromEvent, Observable, of, Subject, tap } from 'rxjs';

const ONE_ETH_TO_GWEI = 1e18;
@Injectable({
  providedIn: 'root',
})
export class SignerService {
  private _web3!: Web3;

  private _isReady = false;
  private _isLogon = new BehaviorSubject(false);
  private _address$ = new BehaviorSubject<string>('');

  constructor() {
    this.preSetup();
  }

  async preSetup() {
    window.addEventListener('load', async () => {
      if (!window.ethereum) return;
      this._web3 = new Web3(window.ethereum);
      try {
        const account = await window.ethereum.send('eth_accounts');
        if (account.result.length) {
          this._address$.next(account.result[0]);
          this._isReady = true;
        }
      } catch (error) {}
      window.ethereum.on('accountsChanged', (accounts: string[]) => {
        this._address$.next(accounts[0]);
      });
    });
  }

  get isReady() {
    return this._isReady;
  }

  get isLogon$() {
    return this._isLogon.asObservable();
  }

  get address$() {
    return this._address$.asObservable();
  }

  get signer() {
    return this._web3;
  }

  async connectWallet() {
    try {
      const account = await window.ethereum.send('eth_requestAccounts');
      if (account.result.length) {
        this._address$.next(account.result[0]);
        this._isReady = true;
      }
    } catch (error) {}
  }
}

declare global {
  interface Window {
    ethereum: any;
    web3: Web3;
  }
}
