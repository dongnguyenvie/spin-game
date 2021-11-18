import { Routes, RouterModule } from '@angular/router';
import { AppLayoutComponent } from '../_layout/app-layout/app-layout.component';

const routes: Routes = [
  {
    path: '',
    component: AppLayoutComponent,
    children: [
      {
        path: '',
        loadChildren: () =>
          import('./home-page/home-page.module').then((m) => m.HomePageModule),
      },
      {
        path: 'sign-in',
        loadChildren: () =>
          import('./signin-page/signin-page.module').then(
            (m) => m.SigninPageModule
          ),
      },
      {
        path: 'sign-up',
        loadChildren: () =>
          import('./signup-page/signup-page.module').then(
            (m) => m.SignupPageModule
          ),
      },
    ],
  },
  {
    path: '**',
    loadChildren: () =>
      import('./not-found-page/not-found-page.module').then(
        (m) => m.NotFoundPageModule
      ),
  },
];

export const PageRoutes = RouterModule.forChild(routes);
