<script setup lang="ts">
import { mapGetters } from 'vuex';
import Showcase from './Showcase.vue';
import EditableText from './EditableText.vue';

import AssignTracksReleaseDialog from './dialogs/AssignTracksReleaseDialog.vue';
import { computed, ref } from 'vue';
import type { Music, Release } from '../types';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

const router = useRouter()

function openTrackSelectModal() {
    assigndialog.value.showModal()
}
function closeTrackSelectModal() {
    assigndialog.value.close()
}

async function fetchData(releaseId): Promise<Release> {
    const req = await fetch("/api/release/" + releaseId)
    return req.json()
}

function applySelection(tracks: Music[]) {
    //temporary.value.push(...tracks)
    for (const t of tracks) {

        data.value.RelatedMusic.push({
            ArtistName: t.Artist,
            Name: t.Tracktitle,
            Position: 999, //irrelevant
            TrackId: t.TrackId
        })
    }
}

async function save(submitEvent: SubmitEvent) {
    const fdata = new FormData(submitEvent.target as HTMLFormElement)
    //console.log(Object.fromEntries(fdata))

    const req = await fetch("/api/release/" + data.value.ReleaseId, {
        credentials: "include",
        method: "PUT",
        body: fdata
    })
    edit.value = false;
    data.value = await(fetchData(props.releaseId))
    closeTrackSelectModal()
}

const props = defineProps({
    "releaseId": String
});

const edit = ref(false);
const assigndialog = ref<HTMLDialogElement>()
const data = ref<Release>(await fetchData(props.releaseId));

const store = useStore();
const isLoggedIn = computed(() => store.getters.isLoggedIn);

console.log(data.value)

</script>



<template>
    <form @submit.prevent="save">
    <Showcase>
        <template #left>
            <img class="cover" :src="$store.getters.s3Host + data.CoverUrl">
            <div class="barcode">
                <vue-barcode :value="data.CatalogId" :options="{ displayValue: true }" tag="svg"></vue-barcode>
            </div>
        </template>

        <template #default>

                <h1 class="title">
                    <EditableText :reverse-style="false" :edit="edit" :value="data.Name" type="text" name="Name">
                    </EditableText>
                </h1>
                <strong>Release Date:</strong>
                <EditableText :reverse-style="false" :edit="edit" :value="data.ReleaseDate" type="date" name="ReleaseDate">
                </EditableText>
                <strong>Release-Code:</strong><span class="release-code">{{ data.CatalogId }}</span>
                
                <strong>ISRC-Code:</strong>
                <EditableText :reverse-style="false" :edit="edit" :value="data.Isrc ?? '-'" type="text" name="Isrc">
                </EditableText>
                
                <div class="tracklist">
                    <h2>Track List</h2>
                    <div v-for="(music, index) in data.RelatedMusic">

                    </div>
<!-- 
    <ol class="tracklist">
        <li v-for="(music, index) in data.RelatedMusic" >
            <a :href="'/catalog/tracks/' + music.TrackId" v-if="!edit">
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
            <div v-else>
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
            </div>
            <input type="hidden" name="tracklist" :value="music.TrackId">
        </li>
    </ol>
        -->
        <a href="#" @click.prevent="openTrackSelectModal" v-if="isLoggedIn && edit" class="add">
                            Add Track to Release
                        </a>
                    
                </div>
                
                <div class="actionbuttons">
                    
                    <button type="button" @click.prevent="edit = !edit" v-if="isLoggedIn">{{ edit ? "Cancel" : "Edit"
                    }}</button>
                <input type="submit" v-if="edit && isLoggedIn">
            </div>
        </template>
    </Showcase>
</form>
    <dialog id="add_tracks_to_release" ref="assigndialog">
        <AssignTracksReleaseDialog @save="applySelection" @close="assigndialog.close()" />
    </dialog>
</template>

<style scoped>
.tracklist ol {
    margin-block-start: 2em;
}
.actionbuttons {
    display: flex;
    gap: 1em;
}

#add_tracks_to_release {
    width: 50%;
}

.tracklist .add {
    display: block;
    padding: 1em 0;
    border-top: 5px dotted white;
    list-style-type: circle;
}

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
    padding: 1em 0 0 0;
    grid-column-start: 1;
    grid-column-end: 3;
}

.tracklist h2 {}

.tracklist li {}
</style>