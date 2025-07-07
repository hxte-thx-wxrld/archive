import { createRouter, createWebHistory, type RouteRecordInfo } from 'vue-router'
import AboutView from './views/AboutView.vue'
import UploadView from './views/UploadView.vue'
import CatalogView from './views/CatalogView.vue'
import SingleView from './views/SingleView.vue'

/*export interface RouteNamedMap {
  about: RouteRecordInfo<
    'about',
    '/about',
    Record<never, never>,
    Record<never, never>,
    never
  >

  catalog: RouteRecordInfo<
    'catalog',
    '/catalog/:mode',
    { mode: string },
    { mode: string },
    string
  >

  singleview: RouteRecordInfo<
    'catalog',
    '/catalog/:mode/:id',
    { mode: string, id: string },
    { mode: string, id: string },
    string
  >
}

// Last, you will need to augment the Vue Router types with this map of routes
declare module 'vue-router' {
  interface TypesConfig {
    RouteNamedMap: RouteNamedMap
  }
}*/

const routes = [
  { path: '/about', component: AboutView },
  { path: '/upload', component: UploadView },
  { path: '/catalog/:mode', component: CatalogView, alias: ["/catalog", "/"] },
  { path: '/catalog/:mode/:id', component: SingleView },
  //{ path: '/action/edit_track', component: SingleView },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})