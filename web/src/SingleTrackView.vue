<script lang="ts" setup>
import TrackShowcase from "./components/TrackShowcase.vue";
</script>
<script lang="ts">

export default {

}
</script>

<template>
    <div>
        <h1>Track</h1>
        <Suspense>
            <TrackShowcase :track-id="$route.params.id"/>
            <template #loading>
                Loading...
            </template>
        </Suspense>
    </div>
</template>

<style scoped></style>
