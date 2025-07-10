<script setup lang="ts">
import { mapGetters } from 'vuex';
import HelloWorld from './components/HelloWorld.vue'
import LoginDialog from './components/dialogs/LoginDialog.vue'
import { mapState } from 'vuex';

</script>
<script lang="ts">
export default {
  mounted() {
    this.$store.dispatch("fetchConfig")
    this.$store.dispatch("fetchSelfUser")
  },
  methods: {
    showLoginModal(event) {
      (this.$refs.loginDialog as HTMLDialogElement).showModal();
    },
    hideLoginModal(event) {
      (this.$refs.loginDialog as HTMLDialogElement).close();
    },
  },
  data() {
    return {

    }
  },
  computed: {
    ...mapGetters(["isLoggedIn"])
  }
}
</script>

<template>
  <div class="root">
    <header>
      <div class="top">
        <RouterLink to="/">
          <img class="logo" src="./assets/logo.png">
        </RouterLink>
        <!--
          <RouterLink class="action" to="/upload">
            upload
          </RouterLink>
          -->
      </div>
      <div class="sub">
        <nav>
          <RouterLink :class="$route.name == 'tracks'" to="/catalog">Catalog</RouterLink>
          <!-- <RouterLink :class="$route.name == 'artists'" to="/artist">Art Gallery</RouterLink> -->
          <RouterLink :class="$route.name == 'releases'" to="/about">About</RouterLink>
        </nav>
      </div>
    </header>
    <main>
      <RouterView />
    </main>
    <footer>
      <div v-if="!isLoggedIn">
        <a href="#" @click.prevent="showLoginModal">Login</a>


        <dialog ref="loginDialog">
          <LoginDialog @close="hideLoginModal" />
        </dialog>
      </div>
      <div v-else>
        <a href="#" @click.prevent="$store.dispatch('logout')">Logout</a>
        <span>{{ $store.Username }}</span>
      </div>
    </footer>
  </div>
</template>

<style scoped>
header {
  padding: 1em 0;
}

.sub nav * {
  background-color: white;
  color: black;
  padding: .5em;
  font-family: "Silkscreen";
}

.selected {
  background-color: grey;
}

.sub nav {
  padding: 1em 0;
  display: grid;
  grid-template-columns: 1fr 1fr;
}

.sub nav a {
  justify-self: center;
}

header .top {
  display: grid;
  grid-template-columns: 1fr;
}

header .top .action {
  align-self: center;
  justify-self: end;
}

footer {
  padding: 1em 0;
}
</style>

<style>
@font-face {
  font-family: "Apricot200L";
  src: url("./assets/fonts/Web437_Apricot_200L.woff") format("woff");
}

@font-face {
  font-family: "ApricotPortable";
  src: url("./assets/fonts/Web437_ApricotPortable.woff") format("woff");
}

@font-face {
  font-family: "ApricotXenC";
  src: url("./assets/fonts/Web437_ApricotXenC.woff") format("woff");
}

@font-face {
  font-family: "Bytesized";
  src: url("./assets/fonts/Bytesized-Regular.ttf")
}

@font-face {
  font-family: "DotGothic16";
  src: url("./assets/fonts/DotGothic16-Regular.ttf")
}

@font-face {
  font-family: "PixelifySans";
  src: url("./assets/fonts/PixelifySans-Regular.ttf");
  font-weight: normal;
  font-style: normal;
}

@font-face {
  font-family: "PixelifySans";
  src: url("./assets/fonts/PixelifySans-Bold.ttf");
  font-weight: bold;
  font-style: normal;
}

@font-face {
  font-family: "Silkscreen";
  src: url("./assets/fonts/Silkscreen-Regular.ttf");
  font-weight: normal;
  font-style: normal;
}

@media (resolution: 1dppx),
(resolution: 2dppx),
(resolution: 3dppx),
(resolution: 4dppx) {
  body {
    font-smooth: never;
  }
}

body {
  text-size-adjust: none;

  -moz-osx-font-smoothing: none;
  -webkit-text-size-adjust: none;
  text-size-adjust: none;
}

h1,
h2 {
  font-family: "Silkscreen";
}

p,
div,
form {
  /*font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;*/
  /*font-family: "biosfont";*/
  font-family: "DotGothic16";

}

/*h1 {
  font-size: 2em;
  margin: 1em 0;
}

h2 {
  font-size: 1.5em;
  margin: 1em 0;
}

body {
  font-size: .75em;
}*/
body {
  margin: 0;
}

a {
  color: white;
  text-decoration: none;
  text-decoration-thickness: 3px;
}

a:hover {
  text-decoration: underline;
}

a.active {
  text-decoration: underline;
}

body {
  background-color: black;
  color: white;
}

.release-code {
  font-family: 'Courier New', Courier, monospace;
  color: white;
}

.icon::before {
  display: inline-block;
  vertical-align: sub;
  padding: 0 .5em;
  width: 1em;
  height: 1em;
}

.icon.tracks::before {
  content: url(./assets/music.svg);
}

.icon.releases::before {
  content: url(./assets/cd.svg);
}

.icon.artists::before {
  content: url(./assets/artist.svg);
}

.icon.inverted::before {
  filter: invert(1)
}

button,
input[type=submit] {
  -webkit-appearance: none;
  background-color: white;
  border: 0;
  font: inherit;
  color: black;
  padding: .5em;
  cursor: pointer;
}

input[type=text],
input[type=password],
input[type=date]
 {
  font-family: inherit;
  font-size: inherit;
  color: white;
  background-color: black;
  border: 0;
  border-bottom: 3px solid white;
}

input[type=text].reverse,
input[type=password].reverse,
input[type=date].reverse
 {
  background-color: white;
  border-bottom: 3px solid black;
  color: black;
}

button:active {
  background-color: grey;
}

/* iphone */
@media only screen and (max-width: 600px) {
  .showcase {
    grid-template-columns: 1fr;
    justify-items: center;
  }

  .showcase .meta {
    display: grid;
    grid-template-columns: 1fr;
    margin: 1em;
    justify-self: left;
  }

  .showcase .cover-area img {
    width: 100%;
    height: auto;
    border: unset;
  }

  .showcase .meta :nth-child(2n) {
    padding: 0 0 1em 0;
  }

  .root {
    padding: 1em;
  }

  .top .logo {
    width: 100%;
  }

}

/* tablets */
@media only screen and (min-width: 600px) {
  .showcase {
    grid-template-columns: 1fr;
  }

  .root {
    padding: 0 5em;
  }

  .top .logo {
    width: 100%;
  }

}

/*desktop*/
@media only screen and (min-width: 768px) {
  .showcase {
    grid-template-columns: 256px 1fr;
    gap: 4em;
  }

  .root {
    padding: 0 10em;
  }

  .top .logo {
    width: 50%;
  }

  header .top a {
    text-align: center;
  }
}
</style>