import { createApp, type InjectionKey } from 'vue'
import App from './App.vue'
import { router } from './router'
import VueBarcode from '@chenfengyuan/vue-barcode';
import { store } from './store';

console.log(import.meta.env.MODE)


export function getDevPrefix() {
    if (import.meta.env.MODE == "development") {
        return "http://localhost:8080"
    } else return "";
}

export function getS3Host() {
    if (import.meta.env.VITE_S3_ENDPOINT) {
        return import.meta.env.VITE_S3_ENDPOINT
    }
}
createApp(App)
    .use(router)
    .use(store)
    .component(VueBarcode.name, VueBarcode)
    //.component("VPaginator", VuePaginator)
    .mount('#app')
