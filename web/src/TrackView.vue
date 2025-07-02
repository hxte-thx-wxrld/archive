<script setup lang="ts">
//import UploadComp from './components/UploadComp.vue'
import BrowseList from './components/TrackList.vue'
</script>

<script lang="ts">
export default {

}
</script>

<template>
    <h1>Tracks</h1>
    <Suspense>
        <BrowseList />
        <template #fallback>
            Loading...
        </template>
    </Suspense>
    <!-- <UploadComp /> -->
</template>

<style>

</style>
