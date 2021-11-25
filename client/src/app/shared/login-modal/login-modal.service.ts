import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class LoginModalService {
  private _isShow$ = new BehaviorSubject(false);

  constructor() {}

  public open() {
    this._isShow$.next(true);
  }

  public close() {
    this._isShow$.next(false);
  }

  get isShow$() {
    return this._isShow$.asObservable();
  }
}
