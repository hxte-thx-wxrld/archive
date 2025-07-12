<script setup lang="ts">

</script>

<template>
    <header>
        <div class="top">
            <RouterLink to="/">
                <img class="logo" src="../assets/logo.png">
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
</style>