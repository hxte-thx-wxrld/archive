import { createRouter, createWebHistory } from 'vue-router'

import HomeView from './HomeView.vue'
import AboutView from './AboutView.vue'
import SingleTrackView from './SingleTrackView.vue'
import SingleReleasesView from './SingleReleasesView.vue'
import SingleArtistView from './SingleArtistView.vue'
import TrackView from './TrackView.vue'
import ReleasesView from './ReleasesView.vue'
import ArtistView from './ArtistView.vue'
import UploadView from './UploadView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/about', component: AboutView },
  { path: '/track/:id', component: SingleTrackView },
  { path: '/track', component: TrackView },
  { path: '/release', component: ReleasesView },
  { path: '/release/:id', component: SingleReleasesView },
  { path: '/artist', component: ArtistView },
  { path: '/artist/:id', component: SingleArtistView },
  { path: '/upload', component: UploadView },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})