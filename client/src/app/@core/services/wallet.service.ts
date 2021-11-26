import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import {
  BehaviorSubject,
  catchError,
  from,
  map,
  of,
  switchMap,
  take,
  tap,
  timer,
} from 'rxjs';
import { API } from '../constant/common.constant';
import {
  WalletData,
  MyWalletRespose,
  MyPackageRespose,
  MyPackageData,
} from '../model/common.model';
import { toPlainString } from '../utils/util';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class WalletService {
  private _wallet$ = new BehaviorSubject<WalletData>(
    null as unknown as WalletData
  );

  constructor(private http: HttpClient, private authSvc: AuthService) {
    this._setup();
  }

  private _setup() {
    const token$ = this.authSvc.token$.pipe(take(1));
    timer(0, 10000)
      .pipe(switchMap(() => token$))
      .subscribe((token) => {
        if (!token) return;
        this.fetchMyWallet$().subscribe();
      });
  }

  get wallet$() {
    return this._wallet$.asObservable();
  }

  fetchMyWallet$() {
    return this.http.get<MyWalletRespose>(API.myWallet).pipe(
      tap((resp) => {
        if (!resp) return;
        this._wallet$.next(resp.data);
      }),
      catchError((err) => {
        return of(null);
      })
    );
  }
}
