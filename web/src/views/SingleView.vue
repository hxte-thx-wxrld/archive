<script lang="ts" setup>
import TrackShowcase from "../components/TrackShowcase.vue";
import ReleaseShowcase from "../components/ReleaseShowcase.vue";
import ArtistShowcase from "../components/ArtistShowcase.vue";
</script>
<script lang="ts">

export default {
    methods: {
        getSubmode() {
            if (!this.$route.params.mode) {
                return "tracks";
            } else return this.$route.params.mode;
        },
        getObjectId() {
            return this.$route.params.id
        }
    }
}
</script>

<template>
    <div v-if="getSubmode() == 'tracks'">
        <div class="mode-title-container">

            <h1 class="mode-title inverted icon tracks">Track</h1>
            <img class="back-icon" src="../assets/arrow.png">
            <RouterLink to="/catalog/tracks">Back to Catalog</RouterLink>
        </div>
        <Suspense>
            <TrackShowcase :track-id="getObjectId()" />
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
    <div v-if="getSubmode() == 'releases'">
        <div class="mode-title-container">

            <h1 class="mode-title inverted icon releases">Release</h1>
            <img class="back-icon" src="../assets/arrow.png">
            <RouterLink to="/catalog/releases">Back to Catalog</RouterLink>
        </div>
        <Suspense>
            <ReleaseShowcase :release-id="getObjectId()" />
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
    <div v-if="getSubmode() == 'artists'">
        <div class="mode-title-container">
            <h1 class="mode-title inverted icon artists">Artist</h1>
            <img class="back-icon" src="../assets/arrow.png">
            <RouterLink to="/catalog/artists">Back to Catalog</RouterLink>
        </div>
        <Suspense>
            <ArtistShowcase :artist-id="getObjectId()" />
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
    padding-bottom: 2em;
    margin-bottom: 2em;
    border-bottom: 2px dashed white;
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
