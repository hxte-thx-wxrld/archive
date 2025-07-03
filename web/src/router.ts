import { createRouter, createWebHistory } from 'vue-router'

import HomeView from './views/HomeView.vue'
import AboutView from './views/AboutView.vue'
import UploadView from './views/UploadView.vue'
import CatalogView from './views/CatalogView.vue'
import SingleView from './views/SingleView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/about', component: AboutView },
  { path: '/upload', component: UploadView },
  { path: '/catalog/:mode', component: CatalogView, alias: ["/catalog"] },
  { path: '/catalog/:mode/:id', component: SingleView },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})