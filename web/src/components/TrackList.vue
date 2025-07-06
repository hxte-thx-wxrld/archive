<script setup lang="ts">
import { mapGetters } from 'vuex'
import { getDevPrefix } from '../main'
import CreateTrackDialog from './CreateTrackDialog.vue'
import type { MusicRow } from '../types'
import { ref } from 'vue';

const data = ref<
    {
        Rows: [MusicRow],
        FullLength: Number
    }>()

async function reloadList() {
    const req = await fetch(getDevPrefix() + "/api/track/")
    return await req.json()
}

data.value = await reloadList()

</script>
<script lang="ts">


export default {
    props: ["page"],

    methods: {
        openNew(event) {
            console.log(event);
            (this.$refs.new_track as HTMLDialogElement).showModal();
        },
        closeNew(event) {
            (this.$refs.new_track as HTMLDialogElement).close();
        },
    },
    computed: {
        ...mapGetters(["isLoggedIn"])
    },
    data() {
        return {}
    }
}
</script>

<template>
    <a href="#" @click.prevent="openNew" v-if="isLoggedIn">Add New</a>
    <div class="browse-list">
        <div class="row" v-for="(item, index) in data.Rows" v-if="data != null">
            <div class="cover-area">
                <img :src="'http://s3.rillo.internal:8333' + item.CoverUrl">
            </div>

            <div class="name-area">
                <RouterLink class="open-track" :to="'/catalog/tracks/' + item.TrackId">
                    <div class="tracktitle">
                        <strong>
                            {{ item.Tracktitle }}
                        </strong>
                    </div>
                    <div class="artist">{{ item.Artist }}</div>
                    <div v-if="item.CatalogNo" class="catalogId">{{ item.CatalogNo }}</div>
                    <small class="catalogId" v-else>No catalog release</small>
                </RouterLink>
            </div>
        </div>
        <div v-else>No Data</div>
    </div>
    <dialog v-if="isLoggedIn" ref="new_track">
        <CreateTrackDialog />
    </dialog>
</template>

<style lang="css" scoped>
a:any-link {
    color: unset;
    text-decoration-line: unset;
}

.browse-list {
    display: grid;
    justify-items: stretch;
    gap: .5em;
}

.browse-list .row .cover-area img {
    width: 64px;
    height: 64px;
}

.browse-list .row {
    border: 1px solid white;
    display: grid;
    grid-template-columns: 64px 1fr;
    grid-template-rows: 64px;
    padding: 1em;
}

.browse-list .row .open-button {
    border: 1px solid red;
}

.browse-list .row .name-area {
    align-self: center;
    margin: .5em;
}
</style>