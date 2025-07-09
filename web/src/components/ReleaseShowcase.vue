<script setup lang="ts">
import { mapGetters } from 'vuex';
import { getDevPrefix, getS3Host } from '../main'
import Showcase from './Showcase.vue';
import EditableText from './EditableText.vue';

import AssignTracksReleaseDialog from './dialogs/AssignTracksReleaseDialog.vue';
import { ref } from 'vue';
import type { Music, Release } from '../types';

function addTrackToRelease(event: MouseEvent) {
    assigndialog.value.showModal()
}

async function fetchData(releaseId): Promise<Release> {
    const req = await fetch(getDevPrefix() + "/api/release/" + releaseId)
    return req.json()
}

const props = defineProps({
    "releaseId": String
});

const edit = ref(false);
const assigndialog = ref<HTMLDialogElement>()
const data = ref<Release>(await fetchData(props.releaseId));

console.log(data.value)

</script>

<script lang="ts">
export default {
    computed: {
        ...mapGetters(["isLoggedIn"])
    },
    methods: {
    }
}
</script>

<template>
    <Showcase>
        <template #left>
            <img class="cover" :src="getS3Host() + data.CoverUrl">
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
                <ol class="tracklist">
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
                    <a href="#" @click.prevent="addTrackToRelease" v-if="isLoggedIn && edit" class="add">
                        Add Track to Release
                    </a>
                </ol>
                
            </div>

            <div class="actionbuttons">

                <button type="button" @click.prevent="edit = !edit" v-if="isLoggedIn">{{ edit ? "Cancel" : "Edit"
                    }}</button>
                <input type="submit" v-if="edit && isLoggedIn">
            </div>
        </template>
    </Showcase>
    <dialog id="add_tracks_to_release" ref="assigndialog">
        <AssignTracksReleaseDialog @save="console.log" @close="assigndialog.close()" />
    </dialog>
</template>

<style scoped>
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