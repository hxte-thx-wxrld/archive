<script setup lang="ts">
import { ref } from 'vue';
import { getS3Host } from '../main'
import type { Artist } from '../types';
import Showcase from './Showcase.vue'

const props = defineProps<{
    artistId: string
}>();

const req = await fetch("/api/artist/" + props.artistId)
const data = ref<Artist>(await req.json())


console.log(data)

</script>

<template>
    <div>
        <Showcase>
            <template #left>
                <img class="cover" :src="getS3Host() + data.ArtistPicture">
            </template>
            <template #default>

                <h1 class="title">{{ data.Name }}</h1>
                <div class="title container">
                    {{ data.Description }}
                </div>
            </template>
        </Showcase>
    </div>
</template>