import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SpinGameWorkspaceComponent } from './components/spin-game-workspace/spin-game-workspace.component';
import { SpinMonitoringComponent } from './components/spin-monitoring/spin-monitoring.component';

@NgModule({
  imports: [CommonModule],
  declarations: [SpinGameWorkspaceComponent, SpinMonitoringComponent],
  exports: [SpinGameWorkspaceComponent, SpinMonitoringComponent],
})
export class SpinModule {}
