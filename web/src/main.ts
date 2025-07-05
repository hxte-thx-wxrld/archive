import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { router } from './router'
import VueBarcode from '@chenfengyuan/vue-barcode';
import { createStore } from 'vuex'

console.log(import.meta.env.MODE)

const store = createStore({
    state: {
        userid: null
    },
    mutations: {
        setUserdata(state, data) {
            console.log(data)
        }
    },
    actions: {
        async login({ dispatch }, fdata) {
            await fetch(getDevPrefix() + "/api/login", {
                method: "POST",
                credentials: 'include',
                body: JSON.stringify({
                    username: fdata.get("username"),
                    password: fdata.get("password"),
                })
            })

            dispatch('fetchSelfUser')
        },
        async logout(state) { },
        async fetchSelfUser({ commit }) {

            //const json = await req.json();

            const whoami = await fetch(getDevPrefix() + "/api/me", {
                method: "GET",
                credentials: 'same-origin'
            })

            const data = await whoami.json();
            commit('setUserdata', data)
        },
    }
})

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
