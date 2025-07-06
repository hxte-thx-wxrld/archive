<script lang="ts" setup>
import TrackShowcase from "../components/TrackShowcase.vue";
import ReleaseShowcase from "../components/ReleaseShowcase.vue";
import ArtistShowcase from "../components/ArtistShowcase.vue";
import { ref } from "vue";

const editmode = ref(false);
</script>
<script lang="ts">

export default {
    computed: {
        submode() {
            if (!this.$route.params.mode) {
                return "tracks";
            } else return this.$route.params.mode;
        },
        objectId() {
            return this.$route.params.id as string
        }
    },
}
</script>

<template>
    <div v-if="submode == 'tracks'">
        <div class="mode-title-container">

            <h1 class="mode-title inverted icon tracks">Track</h1>
            <img class="back-icon" src="../assets/arrow.png">
            <RouterLink to="/catalog/tracks">Back to Catalog</RouterLink>
        </div>
        <Suspense>
            <TrackShowcase :edit="editmode" :track-id="objectId" />
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
    <div v-if="submode == 'releases'">
        <div class="mode-title-container">

            <h1 class="mode-title inverted icon releases">Release</h1>
            <img class="back-icon" src="../assets/arrow.png">
            <RouterLink to="/catalog/releases">Back to Catalog</RouterLink>
        </div>
        <Suspense>
            <ReleaseShowcase :release-id="objectId" />
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
    <div v-if="submode == 'artists'">
        <div class="mode-title-container">
            <h1 class="mode-title inverted icon artists">Artist</h1>
            <img class="back-icon" src="../assets/arrow.png">
            <RouterLink to="/catalog/artists">Back to Catalog</RouterLink>
        </div>
        <Suspense>
            <ArtistShowcase :artist-id="objectId" />
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
</template>

<style scoped>
.mode-title {
    padding: 0
}

.mode-title-container {
    position: sticky;
    top: 0;
    padding-bottom: 2em;
    margin-bottom: 2em;
    border-bottom: 2px dashed white;
    background-color: black;
}

.back-icon {
    filter: invert(1);
    display: inline-block;
    height: 1.5em;
    rotate: 180deg;
    padding: 0 .5em;

    vertical-align: sub;

}
</style>
