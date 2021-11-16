import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PageRoutes } from './page.routing';
import { PagesComponent } from './pages.component';

@NgModule({
  imports: [CommonModule, PageRoutes],
  declarations: [PagesComponent],
})
export class PagesModule {}
