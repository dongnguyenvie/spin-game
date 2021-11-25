import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ProfileService {
  private readonly _isShow$ = new BehaviorSubject(false);

  constructor() {}

  get show$() {
    return this._isShow$.asObservable();
  }

  onOpen(): void {
    this._isShow$.next(true);
  }

  onClose(): void {
    this._isShow$.next(false);
  }
}
