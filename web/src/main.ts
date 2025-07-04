import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { router } from './router'
import VueBarcode from '@chenfengyuan/vue-barcode';

console.log(import.meta.env.MODE)
export function getDevPrefix() {
    if (import.meta.env.MODE == "development") {
        return "http://localhost:8080"
    } else return "";
}
createApp(App)
    .use(router)
    .component(VueBarcode.name, VueBarcode)
    //.component("VPaginator", VuePaginator)
    .mount('#app')
