<script setup lang="ts">
import { getDevPrefix } from '../main'
import Showcase from './Showcase.vue';

const props = defineProps({
    "releaseId": String
});

const req = await fetch(getDevPrefix() + "/api/release/" + props.releaseId)
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
            <div class="barcode">
                <vue-barcode :value="data.CatalogId" :options="{ displayValue: true }" tag="svg"></vue-barcode>
            </div>
        </template>

        <template #default>
            <h1 class="title">{{ data.Name }}</h1>
            <strong>Release Date:</strong><span>{{ data.ReleaseDate }}</span>
            <strong>Release-Code:</strong><span class="release-code">{{ data.CatalogId }}</span>

            <strong>ISRC-Code:</strong><span>{{ data.Isrc ?? "n. A." }}</span>
            <span>Play: </span><a v-if="data.PublicUrl != null">Play</a><small v-else>stream not available</small>

            <div class="tracklist">

                <h2>Track List</h2>
                <ol>
                    <li v-for="(music, index) in data.RelatedMusic">
                        <a :href="'/catalog/tracks/' + music.TrackId">
                            <div>
                                <p>
                                    <strong>
                                        {{ music.Name }}
                                    </strong>
                                </p>
                            </div>
                            <div>
                                <p>
                                    {{ music.ArtistName }}
                                </p>
                            </div>
                        </a>
                    </li>
                </ol>
            </div>
        </template>
    </Showcase>
</template>

<style scoped>
.barcode {
    z-index: -1;
}

.barcode svg {
    width: 100%;
}

img.cover {
  width: 256px;
  height: 256px;
  border: 1px solid white;
}

li {
    padding: 0 0 0 0;
}

.tracklist {
    padding: 2em 0;
    grid-column-start: 1;
    grid-column-end: 3;
}

.tracklist h2 {
    margin: 1em 0;
}

.tracklist li {
    margin: .5em;
}
</style>