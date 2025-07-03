<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue'

</script>
<script lang="ts">
export default {

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
          <RouterLink to="/track">Tracks</RouterLink>
          <RouterLink to="/release">Releases</RouterLink>
          <RouterLink to="/artist">Artists</RouterLink>
        </nav>
      </div>
    </header>
    <main>
      <RouterView />
    </main>
    <footer>
      TESTING
    </footer>
  </div>
</template>

<style scoped>
.top .logo {

  width: 50%;
}

.sub nav {
  padding: 1em 0;
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
}

.sub nav a {
  justify-self: center;
}

header .top {
  display: grid;
  grid-template-columns: 1fr 1fr;
}
header .top .action {
  align-self:center;
  justify-self: end;
}
footer {
  padding: 1em 0;
}
</style>
