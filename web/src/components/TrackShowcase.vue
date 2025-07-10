<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import Showcase from './Showcase.vue';
import EditableText from './EditableText.vue';
import { computed, ref } from 'vue';
import { type Analysis, type Music } from '../types';
import { mapGetters } from 'vuex';
import AssignedArtistsPicker from './AssignedArtistsPicker.vue';
import { router } from '../router';
import { useStore } from 'vuex';

function saveAnalysisToDisk() {

    // Create a blob of the data
    var file = new Blob([JSON.stringify(analysisData.value)], {
        type: 'application/json'
    });

    let exportUrl = URL.createObjectURL(file);
    window.location.assign(exportUrl);
    URL.revokeObjectURL(exportUrl);
    // Save the file

}

async function fetchData(trackId): Promise<Music> {

    const req = await fetch("/api/track/" + trackId)
    if (req.ok) {
    } else {
        showError.value = true;
    }
    return req.json();
}
async function save(event: SubmitEvent) {
    const fdata = new FormData(event.target as HTMLFormElement);

    const data = {
        Tracktitle: fdata.get("Tracktitle").toString(),
        ReleaseDate: fdata.get("ReleaseDate").toString(),
        ArtistId: fdata.get("ArtistId").toString()
    }

    console.log(data);

    const req = await fetch("/api/track/" + props.trackId, {
        method: "PUT",
        credentials: "include",
        body: JSON.stringify(data)
    })
}

async function deleteTrack(event) {
    const req = await fetch("/api/track/" + props.trackId, {
        method: "DELETE",
        credentials: "include",
    })

    if (req.ok) {
        router.push("/catalog/tracks")
    }
}

async function getAnalysisData() {
    const req = await fetch("/api/track/" + props.trackId + "/analysis", {
        method: "GET",
    })

    return await req.json();
}

const props = defineProps({
    "trackId": String,
    "edit": Boolean
});
const edit = ref(false)
const showError = ref(false)
const data = ref<Music>();
const analysisData = ref<Analysis>();
analysisData.value = await getAnalysisData()

const store = useStore();
const isLoggedIn = computed(() => store.getters.isLoggedIn)

data.value = await fetchData(props.trackId)


console.log(data)
</script>

<template>
    <form @submit.prevent="save">

        <Showcase>
            <template #left>
                <img class="cover" :src="$store.getters.s3Host + data.CoverUrl">
            </template>
            <template #default>
                <span class="title">
                    <EditableText :reverse-style="false" :edit="edit" :value="data.Tracktitle" name="Tracktitle"
                        type="text" />
                    <div class="delete-tools">
                        <a href="#" v-if="edit" @click.prevent="deleteTrack">
                            <img class="icon" src="../assets/trash.svg">
                        </a>
                    </div>
                </span>
                <strong>Artist:</strong>
                <RouterLink :to="'/catalog/artists/' + data.ArtistId" v-if="!edit">{{ data.Artist }}</RouterLink>
                <AssignedArtistsPicker v-else />
                <strong>Length:</strong>
                <span>{{ data.Length }}</span>

                <strong>Track Releasedate:</strong>
                <EditableText :reverse-style="false" :edit="edit" :value="data.ReleaseDate" type="date"
                    name="ReleaseDate"></EditableText>
                <strong>Release-Code: </strong>
                <a v-if="data.ReleaseId != null" class="release-code" :href="'/catalog/releases/' + data.ReleaseId">{{
                    data.CatalogNo }}</a>
                <small v-else>no catalog release</small>
                <p>Play: </p>
                <p>
                    <audio v-if="data.PublicUrl != null" controls :src="$store.getters.s3host + data.PublicUrl"></audio>
                    <small v-else>stream not available</small>
                </p>

                <div class="actionbuttons">

                    <button type="button" @click.prevent="edit = !edit" v-if="isLoggedIn">{{ edit ? "Cancel" : "Edit"
                    }}</button>
                    <input type="submit" v-if="edit && isLoggedIn">
                </div>

            </template>
        </Showcase>

        <div class="a-title">
            <span>Analysis Data</span>
             <a :href="'/api/track/' + props.trackId + '/analysis?type=json'" >Save</a>
        </div>
        <div class="table">
            <div class="a-row">
                <span>Tempo</span>
                <span>{{ analysisData.Tempo }}</span>
            </div>
            <div class="a-row">
                <span>Zero Crossing Rate</span>
                <span>{{ analysisData.ZCR }}</span>
            </div>
            <div class="a-row">
                <span>Centroid</span>
                <span>{{ analysisData.Centroid }}</span>
            </div>
            <div class="a-row">
                <span>Flatness</span>
                <span>{{ analysisData.Flatness }}</span>
            </div>
            <div class="a-row">
                <span>RMS</span>
                <span>{{ analysisData.RMS }}</span>
            </div>

            <div class="a-row">
                <span>Rolloff</span>
                <span>{{ analysisData.Rolloff }}</span>
            </div>

            <div class="a-row">
                <span>Duration</span>
                <span>{{ analysisData.Duration }}</span>
            </div>
        </div>
    </form>

</template>

<style scoped>
.table {
    display: flex;
    flex-direction: column;
    font-family: "Bytesized";
}

.table .a-row {
    display: flex;
    justify-content: space-between;
}

.a-title {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    font-size: 2em;
}
.a-title a {
    font-size: 1em;
}

.actionbuttons {
    display: flex;
    gap: 1em;
}

.title {
    display: flex;
    justify-content: space-between;
    font-size: 2em;
    font-family: "Silkscreen";
}

.icon {
    filter: invert(1);
    width: 24px;
    height: 24px;
}

.cover {
    border: 1px solid white;
}
</style>