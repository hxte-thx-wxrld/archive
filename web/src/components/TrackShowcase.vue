<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { getDevPrefix } from '../main'
import Showcase from './Showcase.vue';
import EditableText from './EditableText.vue';
import { ref } from 'vue';
import type { MusicRow, TrackEditRequest } from '../types';
import { mapGetters } from 'vuex';
import AssignedArtistsPicker from './AssignedArtistsPicker.vue';
import { router } from '../router';

async function fetchData(trackId): Promise<MusicRow> {

    const req = await fetch(getDevPrefix() + "/api/track/" + trackId)
    return req.json()
}
async function save(event: SubmitEvent) {
    const fdata = new FormData(event.target as HTMLFormElement);

    const data: TrackEditRequest = {
        Tracktitle: fdata.get("Tracktitle").toString(),
        ReleaseDate: fdata.get("ReleaseDate").toString(),
        ArtistId: fdata.get("ArtistId").toString()
    }

    console.log(data);

    const req = await fetch(getDevPrefix() + "/api/track/" + props.trackId, {
        method: "PUT",
        credentials: "include",
        body: JSON.stringify(data)
    })
}

const props = defineProps({
    "trackId": String,
    "edit": Boolean
});
const edit = ref(false)
const data = ref<MusicRow>();

data.value = await fetchData(props.trackId)


console.log(data)
</script>
<script lang="ts">
export default {
    methods: {

    },
    computed: {
        ...mapGetters(["isLoggedIn"])
    }
}
</script>

<template>
    <form @submit.prevent="save">

        <Showcase>
            <template #left>
                <img class="cover" :src="'http://s3.rillo.internal:8333' + data.CoverUrl">
            </template>
            <template #default>
                <h1 class="title">
                    <EditableText :edit="edit" :value="data.Tracktitle" name="Tracktitle" type="text" />
                </h1>
                <strong>Artist:</strong>
                <RouterLink :to="'/catalog/artists/' + data.ArtistId" v-if="!edit">{{ data.Artist }}</RouterLink>
                <AssignedArtistsPicker v-else />
                <strong>Length:</strong>
                <span>{{ data.Length }}</span>

                <strong>Track Releasedate:</strong>
                <EditableText :edit="edit" :value="data.ReleaseDate" type="date" name="ReleaseDate"></EditableText>
                <strong>Release-Code: </strong>
                <a v-if="data.ReleaseId != null" class="release-code" :href="'/catalog/releases/' + data.ReleaseId">{{
                    data.CatalogNo }}</a>
                <small v-else>no catalog release</small>
                <p>Play: </p>
                <p>
                    <audio v-if="data.PublicUrl != null" controls :src="data.PublicUrl"></audio>
                    <small v-else>stream not available</small>
                </p>

                <div class="actionbuttons">

                    <button type="button" @click.prevent="edit = !edit" v-if="isLoggedIn">{{ edit ? "Cancel" : "Edit"
                    }}</button>
                    <input type="submit" v-if="edit && isLoggedIn">
                </div>

            </template>
        </Showcase>
    </form>

</template>

<style scoped>
.actionbuttons {
    display: flex;
    gap: 1em;
}
</style>