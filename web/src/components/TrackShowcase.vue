<script lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { getDevPrefix } from '../main'

export default {
    props: ['trackId'],
    async setup(props) {
        const req = await fetch(getDevPrefix() + "/api/track/" + props.trackId)
        const data = await req.json()

        console.log(data)
        return {
            data
        }
    }
}
</script>

<template>
    <div class="showcase">
        <div class="cover-area">
            <img :src="'http://s3.rillo.internal:8333' + data.CoverUrl">
        </div>
        <div class="meta">

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

        </div>

    </div>
</template>

<style scoped></style>