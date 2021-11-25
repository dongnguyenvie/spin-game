import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SharedModule } from '../shared/shared.module';
import { AppHeaderComponent } from './app-header/app-header.component';
import { AppLayoutComponent } from './app-layout/app-layout.component';
import { AppNavigationComponent } from './app-navigation/app-navigation.component';
import { AppRoutingModule } from '../app-routing.module';
import { ProfileModule } from '../featues/profile/profile.module';

@NgModule({
  imports: [AppRoutingModule, CommonModule, SharedModule, ProfileModule],
  declarations: [
    AppHeaderComponent,
    AppLayoutComponent,
    AppNavigationComponent,
  ],
  exports: [AppHeaderComponent, AppNavigationComponent],
})
export class LayoutModule {}
