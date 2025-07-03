<script setup lang="ts">
//import UploadComp from './components/UploadComp.vue'
import BrowseList from './components/TrackList.vue'
</script>

<script lang="ts">
export default {
    methods: {
        
    }
}
</script>

<template>
    <h1>Tracks</h1>
    <div class=""><button type="button" @click.prevent="">+</button></div>
    <Suspense>
        <BrowseList />
        <template #fallback>
            Loading...
        </template>
    </Suspense>
    <!-- <UploadComp /> -->
</template>

<style>
h1 {
    padding: 1em 0;
}
</style>
