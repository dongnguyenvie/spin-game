import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import {
  BehaviorSubject,
  catchError,
  map,
  of,
  switchMap,
  take,
  tap,
  timer,
} from 'rxjs';
import { API } from '../constant/common.constant';
import {
  AwardOptionResponse,
  MyPackageData,
  MyPackageRespose,
  PlayGameResponse,
} from '../model/common.model';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class GameService {
  private _package$ = new BehaviorSubject<MyPackageData>(
    null as unknown as MyPackageData
  );
  private _processing = false;

  constructor(private http: HttpClient, private authSvc: AuthService) {
    this._setup();
  }

  private _setup() {
    const token$ = this.authSvc.token$.pipe(take(1));
    timer(0, 10000)
      .pipe(switchMap(() => token$))
      .subscribe((token) => {
        if (!token || this.processing) return;
        this.fetchMyPackage$().subscribe();
      });
  }

  get package$() {
    return this._package$.asObservable();
  }

  get processing() {
    return this._processing;
  }

  fetchPlayGame$() {
    return this.http.post<PlayGameResponse>(API.playGame, {}).pipe(
      map((resp) => resp.data),
      catchError(() => of(null))
    );
  }

  fetchAwardOptions$() {
    return this.http.get<AwardOptionResponse>(API.awards, {}).pipe(
      map((resp) => resp.data),
      catchError(() => of(null))
    );
  }

  fetchMyPackage$() {
    return this.http.get<MyPackageRespose>(API.myPackage).pipe(
      map((resp) => resp.data),
      tap((data) => {
        if (!data) return;
        this._package$.next(data);
      }),
      catchError(() => of(null))
    );
  }

  fetchBuyPackage$(num: number) {
    this._processing = true;
    return this.http
      .post(API.buyPackage, {
        quantity: num,
      })
      .pipe(
        tap(() => {
          this.fetchMyPackage$().subscribe();
          this._processing = false;
        })
      );
  }

  startProcessing(): void {
    this._processing = true;
  }

  stopProcessing(): void {
    this._processing = false;
  }
}
