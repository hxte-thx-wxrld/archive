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
        <h1>Track</h1>
        <Suspense>
            <TrackShowcase :track-id="getObjectId()" />
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
    <div v-if="getSubmode() == 'releases'">
        <h1>Release</h1>
        <Suspense>
            <ReleaseShowcase :release-id="getObjectId()" />
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
    <div v-if="getSubmode() == 'artists'">
        <h1>Artist</h1>
        <Suspense>
            <ArtistShowcase :artist-id="getObjectId()" />
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
</template>

<style scoped></style>
