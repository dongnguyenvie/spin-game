import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SpinGameWorkspaceComponent } from './components/spin-game-workspace/spin-game-workspace.component';
import { SpinMonitoringComponent } from './components/spin-monitoring/spin-monitoring.component';
import { ProfileModule } from '../profile/profile.module';

@NgModule({
  imports: [CommonModule, ProfileModule],
  declarations: [SpinGameWorkspaceComponent, SpinMonitoringComponent],
  exports: [SpinGameWorkspaceComponent, SpinMonitoringComponent],
})
export class SpinModule {}
