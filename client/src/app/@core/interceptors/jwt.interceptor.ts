import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor,
} from '@angular/common/http';
import { defer, map, Observable, of, switchMap, withLatestFrom } from 'rxjs';

import { AuthService } from '../services/auth.service';
import { ENDPOINT } from '../constant/common.constant';

@Injectable()
export class JwtInterceptor implements HttpInterceptor {
  constructor(private authSvc: AuthService) {}

  intercept(
    request: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    const token$ = defer(() =>
      of(null).pipe(
        withLatestFrom(this.authSvc.token$),
        map(([_, token]) => token)
      )
    );

    return token$.pipe(
      switchMap((token) => {
        const isApiUrl = request.url.startsWith(ENDPOINT);
        const isLoggedIn = token;
        if (isLoggedIn && isApiUrl) {
          request = request.clone({
            setHeaders: { Authorization: `Bearer ${token}` },
          });
        }
        return next.handle(request);
      })
    );
  }
}
