<script setup lang="ts">
//import UploadComp from './components/UploadComp.vue'
import TrackList from './components/TrackList.vue';
</script>

<script lang="ts">
export default {

}
</script>

<template>
    <h1>Home</h1>
    <Suspense>
        <TrackList />
        <template #fallback>
            Loading...
        </template>
    </Suspense>
    <!-- <UploadComp /> -->
</template>

<style scoped>
h1 {
    padding: 1em 0;
}
</style>
