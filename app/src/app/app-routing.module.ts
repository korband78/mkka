import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: 'intro', data: {animation: 'intro_page'},
    loadChildren: () => import('./route/intro/index.module').then( m => m.IndexPageModule)
  },
  {
    path: '**', data: {animation: 'index_page'},
    loadChildren: () => import('./route/index/index.module').then( m => m.IndexPageModule)
  },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, { preloadingStrategy: PreloadAllModules, anchorScrolling: "enabled" })
  ],
  exports: [RouterModule]
})
export class AppRoutingModule { }
