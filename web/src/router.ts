import { createRouter, createWebHistory } from 'vue-router'
import AboutView from './views/AboutView.vue'
import UploadView from './views/UploadView.vue'
import CatalogView from './views/CatalogView.vue'
import SingleView from './views/SingleView.vue'

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