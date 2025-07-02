<script lang="ts" setup>
import ReleaseShowcase from './components/ReleaseShowcase.vue';
</script>
<script lang="ts">

export default {

}
</script>

<template>
    <h1>Release</h1>
    <Suspense>
        <ReleaseShowcase :release-id="$route.params.id"/>
        <template #loading>
            Loading...
        </template>
    </Suspense>
    </template>

<style scoped></style>
