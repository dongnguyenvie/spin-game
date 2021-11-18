import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomePageComponent } from './home-page.component';
import { SpinModule } from 'src/app/featues/spin/spin.module';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    component: HomePageComponent,
  },
];

@NgModule({
  imports: [CommonModule, SpinModule, RouterModule.forChild(routes)],
  declarations: [HomePageComponent],
})
export class HomePageModule {}
