import { createApp, type InjectionKey } from 'vue'
import App from './App.vue'
import { router } from './router'
import VueBarcode from '@chenfengyuan/vue-barcode';
import { store } from './store';

console.log(import.meta.env.MODE)

createApp(App)
    .use(router)
    .use(store)
    .component(VueBarcode.name, VueBarcode)
    //.component("VPaginator", VuePaginator)
    .mount('#app')
