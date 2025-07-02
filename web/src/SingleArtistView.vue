<script lang="ts" setup>
//import TrackShowcase from "./components/TrackShowcase.vue";
import ArtistShowcase from "./components/ArtistShowcase.vue";
</script>
<script lang="ts">

export default {

}
</script>

<template>
    <div>
        <h1>Artist</h1>
        <Suspense>
            <ArtistShowcase :artist-id="$route.params.id"/>
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
</template>

<style scoped></style>
