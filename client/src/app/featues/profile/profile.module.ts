import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProfileModalComponent } from './components/profile-modal/profile-modal.component';
import { ProfileComponent } from './components/profile/profile.component';
import { ProfileButtonComponent } from './components/profile-button/profile-button.component';
import { ProfileTabComponent } from './components/profile-tab/profile-tab.component';
import { VarDirective } from 'src/app/@core/directives/ng-var.directive';
import { ProfileMonitorComponent } from './components/profile-monitor/profile-monitor.component';

@NgModule({
  imports: [CommonModule],
  providers: [],
  declarations: [
    ProfileModalComponent,
    ProfileComponent,
    ProfileButtonComponent,
    ProfileTabComponent,
    ProfileMonitorComponent,
  ],
  exports: [ProfileComponent, ProfileButtonComponent, ProfileMonitorComponent],
})
export class ProfileModule {}
