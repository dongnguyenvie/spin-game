import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Signer } from 'crypto';
import {
  BehaviorSubject,
  catchError,
  defer,
  from,
  map,
  mergeMap,
  of,
  Subject,
  switchMap,
  take,
  takeUntil,
  tap,
  timer,
} from 'rxjs';
import { API } from '../constant/common.constant';
import {
  WalletData,
  MyWalletRespose,
  MyPackageRespose,
  MyPackageData,
  WithDrawResponse,
} from '../model/common.model';
import { toPlainString } from '../utils/util';
import { AuthService } from './auth.service';
import { GameService } from './game.service';
import { SignerService } from './signer.service';

@Injectable({
  providedIn: 'root',
})
export class WalletService {
  private _wallet$ = new BehaviorSubject<WalletData>(
    null as unknown as WalletData
  );
  private readonly _killTrigger$ = new Subject();

  constructor(
    private http: HttpClient,
    private authSvc: AuthService,
    private signer: SignerService,
    private gameSvc: GameService
  ) {
    this._setup();
  }

  private _setup() {
    const token$ = this.authSvc.token$.pipe(take(1));
    timer(0, 10000)
      .pipe(switchMap(() => token$))
      .subscribe((token) => {
        if (!token || this.gameSvc.processing) return;
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

  fetchWithdraw$(amount: number) {
    return this.http
      .post<WithDrawResponse>(API.withdraw, {
        amount,
      })
      .pipe(
        map((resp) => resp.data),
        tap((resp) => {
          if (!resp) return;
          this.fetchMyWallet$().subscribe();
        }),
        catchError((err) => {
          return of(null);
        })
      );
  }
}
