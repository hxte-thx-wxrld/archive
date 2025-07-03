import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { router } from './router'

console.log(import.meta.env.MODE)
export function getDevPrefix() {
    if (import.meta.env.MODE == "development") {
        return "http://localhost:8080"
    } else return "";
}
createApp(App)
    .use(router)
    .mount('#app')
