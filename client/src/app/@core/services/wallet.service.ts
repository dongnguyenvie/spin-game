import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, catchError, of, switchMap, tap, timer } from 'rxjs';
import { API } from '../constant/common.constant';
import { WalletData, WalletRespose } from '../model/common.model';

@Injectable({
  providedIn: 'root',
})
export class WalletService {
  private _wallet$ = new BehaviorSubject<WalletData>(
    null as unknown as WalletData
  );

  constructor(private http: HttpClient) {
    this._setup();
  }

  private _setup() {
    timer(0, 10000).subscribe((resp) => {
      this.fetchMyWallet().subscribe();
    });
  }

  get wallet$() {
    return this._wallet$.asObservable();
  }

  fetchMyWallet() {
    return this.http.get<WalletRespose>(API.myWallet).pipe(
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
