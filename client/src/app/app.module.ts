import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { JwtInterceptor } from './@core/interceptors/jwt.interceptor';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ProfileModule } from './featues/profile/profile.module';
import { ImageComponent } from './shared/image/image.component';
import { LoginModalModule } from './shared/login-modal/login-modal.module';
import { LayoutModule } from './_layout/layout.module';

const DefaultModule = [
  HttpClientModule,
  BrowserModule,
  // FormsModule,
  // ReactiveFormsModule,
];
@NgModule({
  imports: [
    ...DefaultModule,
    AppRoutingModule,
    LayoutModule,
    LoginModalModule,
    ProfileModule,
  ],
  declarations: [AppComponent],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },
  ],
  bootstrap: [AppComponent],
  exports: [],
})
export class AppModule {}
