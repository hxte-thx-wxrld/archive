<script setup lang="ts">
//import UploadComp from './components/UploadComp.vue'
import ReleaseList from './components/ReleaseList.vue'
</script>

<script lang="ts">
export default {

}
</script>

<template>
    <h1>Releases</h1>
    <Suspense>
        <ReleaseList />
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
