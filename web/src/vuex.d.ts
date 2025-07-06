import { Store } from 'vuex'

declare module 'vue' {
  // declare your own store states

  // provide typings for `this.$store`
  interface ComponentCustomProperties {
    $store: Store<State>
  }
}