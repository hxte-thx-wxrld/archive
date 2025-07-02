<script setup lang="ts">
//import UploadComp from './components/UploadComp.vue'
import ArtistList from './components/ArtistList.vue';
import BrowseList from './components/TrackList.vue'
</script>

<script lang="ts">
export default {

}
</script>

<template>
    <h1>Artists</h1>
    <Suspense>
        <ArtistList />
        <template #fallback>
            Loading...
        </template>
    </Suspense>
    <!-- <UploadComp /> -->
</template>

<style>

</style>
