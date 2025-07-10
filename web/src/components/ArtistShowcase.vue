<script setup lang="ts">
import { ref } from 'vue';
import { type PaginatedMusicLookup, type Artist } from '../types';
import Showcase from './Showcase.vue'
import Paginator from './Paginator.vue';
import { useRouter } from 'vue-router';

async function getTracksByArtist(artistId, page) {
    const req = await fetch("/api/artist/" + artistId + "/tracks?offset=" + page)
    const data = await req.json()
    if (req.ok) {
        console.log(data)
        return data;
    } else {
        console.error(data)
    }
}

async function switchMusicCollection(p) {
    page.value = p
    music.value = await getTracksByArtist(props.artistId, p)
}

const props = defineProps<{
    artistId: string
}>();

const router = useRouter()


const page = ref(0);
const req = await fetch("/api/artist/" + props.artistId)
const data = ref<Artist>(await req.json())

const areq = await getTracksByArtist(props.artistId, page.value)
const music = ref<PaginatedMusicLookup>(areq)


console.log(data)

</script>

<template>
    <div>
        <Showcase>
            <template #left>
                <img class="cover" :src="$store.getters.s3Host + data.ArtistPicture">
            </template>
            <template #default>

                <h1 class="title">{{ data.Name }}</h1>
                <div class="title container">
                    {{ data.Description }}
                </div>
            </template>
        </Showcase>

        <h2>Tracks by Artist</h2>
        <div class="list">

            <div class="row" v-for="(item, index) in music.Rows">
                <a href="#" @click.prevent="router.push('/catalog/tracks/' + item.TrackId)">
                    <div><strong>{{ item.Tracktitle }}</strong></div>
                    <div>{{ item.Length ?? "00:00:00" }}</div>
                </a>
            </div>

            <Paginator :page="page" :pagecount="music.FullLength" @switch-page="switchMusicCollection"></Paginator>
        </div>
    </div>
</template>

<style scoped>
.list {
    display: flex;
    flex-direction: column;
    gap: .5em;
}
.row {
    border: 1px solid white;
    padding: 1em;
}
</style>