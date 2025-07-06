<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { getDevPrefix } from '../main'
import Showcase from './Showcase.vue';

const props = defineProps({
    "trackId": String
});
const req = await fetch(getDevPrefix() + "/api/track/" + props.trackId)
const data = await req.json()

console.log(data)
</script>
<script lang="ts">
export default {

}
</script>

<template>
    <Showcase>
        <template #left>
            <img class="cover" :src="'http://s3.rillo.internal:8333' + data.CoverUrl">
        </template>
        <template #default>
            <h1 class="title">{{ data.Tracktitle }}</h1>
            <strong>Artist:</strong>
            <RouterLink :to="'/catalog/artist/' + data.ArtistId">{{ data.Artist }}</RouterLink>
            <strong>Length:</strong>
            <span>{{ data.Length }}</span>
            <strong>Release-Code: </strong>
            <a v-if="data.ReleaseId != null" class="release-code" :href="'/catalog/releases/' + data.ReleaseId">{{ data.CatalogNo }}</a>
            <small v-else>no catalog release</small>
            <p>Play: </p>
            <p>
                <audio v-if="data.PublicUrl != null" controls :src="'http://10.0.0.4:8333' + data.PublicUrl"></audio>
                <small v-else>stream not available</small>
            </p>
        </template>
    </Showcase>
</template>

<style scoped></style>