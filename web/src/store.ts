import type { InjectionKey } from "vue"
import { getDevPrefix } from "./main"
import { createStore, Store } from "vuex"
import type { Artist } from "./types"

export interface State {
  UserId: null | string
  Username: null | string
  Admin: boolean,
  AssignedArtists: Artist[]
}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    state: {
        UserId: null,
        Username: null,
        Admin: false,
        AssignedArtists: []
    },
    mutations: {
        setUserdata(state, data) {
            console.log(data)
            state.UserId = data.UserId
            state.Username = data.Username
            state.Admin = data.Admin
            state.AssignedArtists = data.AssignedArtists
        },
        removeUserdata(state) {
            state.UserId = null
            state.Username = null
            state.Admin = false
            state.AssignedArtists = []
        }
    },
    getters: {
        isLoggedIn(state) {
            return state.UserId != null
        },
        isAdmin(state) {
            return state.UserId != null && state.Admin
        }
    },
    actions: {
        async login({ dispatch }, fdata) {
            const req = await fetch(getDevPrefix() + "/api/login", {
                method: "POST",
                credentials: 'include',
                body: JSON.stringify({
                    username: fdata.get("username"),
                    password: fdata.get("password"),
                })
            })

            if (req.ok) {
                return dispatch('fetchSelfUser')
            } else {
                const j = await req.json()
                return Promise.reject(j)
            }

        },
        async logout({ commit }) {
            const req = await fetch(getDevPrefix() + "/api/logout", {
                method: "POST",
                credentials: 'include'
            })

            if (req.ok) {
                commit('removeUserdata')
            }
        },
        async fetchSelfUser({ commit }) {

            //const json = await req.json();

            const whoami = await fetch(getDevPrefix() + "/api/me", {
                method: "GET",
                credentials: 'include'
            })

            const data = await whoami.json();
            commit('setUserdata', data)
        },
    }
})