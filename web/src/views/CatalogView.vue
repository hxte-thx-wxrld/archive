<script setup lang="ts">
//import UploadComp from './components/UploadComp.vue'
import ArtistList from '../components/ArtistList.vue';
import ReleaseList from '../components/ReleaseList.vue';
import TrackList from '../components/TrackList.vue';
</script>

<script lang="ts">
export default {
    data() {
        return {
            showSearchbar: false
        }
    },
    methods: {
        getSubmode() {
            if (!this.$route.params.mode) {
                return "tracks";
            } else return this.$route.params.mode;
        }
    }
}
</script>

<template>
    <div class="top">

        <h1>Catalog</h1>

        <nav class="mode">
            <RouterLink class="icon tracks" :class="getSubmode() == 'tracks' ? 'active' : ''" to="/catalog/tracks">Tracks</RouterLink>
            <RouterLink class="icon releases" :class="getSubmode() == 'releases' ? 'active' : ''" to="/catalog/releases">Releases</RouterLink>
            <RouterLink class="icon artists" :class="getSubmode() == 'artists' ? 'active' : ''" to="/catalog/artists">Artists</RouterLink>
            <a href="#" @click.prevent="showSearchbar = !showSearchbar">{{ showSearchbar ? "-" : "+" }}</a>
        </nav>
        <div class="tools" :class="showSearchbar ? '' : 'hide'">
            <input class="search" type="text">
            <input class="go" type="submit">
        </div>
    </div>


    <div v-if="getSubmode() == 'tracks'">
        <Suspense>
            <TrackList />
            <template #fallback>
                Loading...
            </template>
        </Suspense>

    </div>
    <div v-else-if="getSubmode() == 'releases'">

        <Suspense>
            <ReleaseList />
            <template #fallback>
                Loading...
            </template>
        </Suspense>
    </div>
    <div v-else-if="getSubmode() == 'artists'">

        <Suspense>
            <ArtistList />
            <template #fallback>
                Loading...
            </template>
        </Suspense>
    </div>
    <!-- <UploadComp /> -->
</template>

<style scoped>



h1 {
    padding: 1em 0;
}

.hide {
    display: none !important;
}

nav {
    display: grid;
}

.tools {
    border: 1px solid white;
    padding: 1em;
    margin: 0 0 1em 0;
    display: grid;
    grid-template-columns: 1fr 100px;
}

.tools .search {
    border: unset;
    border-bottom: 3px solid white;
    background-color: transparent;
    color: white;
    font-family: "Apricot200L";
}

.tools .go:active {
    background-color: black;
    color: white;
}
.tools .go {
    justify-self: stretch;
    border: unset;
    background-color: white;
    color: black;
    font-family: "Apricot200L";
    padding: 1em;
    margin: 0 .5em;
}

.top {
    position: sticky;
    top: 0;
    background-color: black;
}

nav.mode {
    padding: 1em 0;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 32px;
}

nav.mode a {
    justify-self: stretch;
}

nav.mode a.active {
    background-color: grey;
    color: white;
}

nav.mode a.active::before {
    filter: invert(1)
}

nav.mode>* {
    background-color: white;
    color: black;
    padding: .5em;
    font-family: "Apricot200L";
}
</style>
