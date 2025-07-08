<script setup lang="ts">
import { mapGetters } from 'vuex'
import { getDevPrefix } from '../../main'
import CreateTrackDialog from '../dialogs/TrackUploadDialog.vue'

import { inject, ref } from 'vue';
import type { MusicRow } from '../../types';

const Rows = inject<MusicRow[]>("Rows")
const emit = defineEmits<{
    'trackSelect': [MusicRow]
}>()

</script>

<script lang="ts">
export default {
    methods: {

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

        <div class="row" v-for="(item, index) in Rows">
            <div class="cover-area" v-if="showCover">
                <img :src="'http://s3.rillo.internal:8333' + item.CoverUrl">
            </div>

            <div class="name-area">

                <a href="#" class="open-track" @click.prevent="emit('trackSelect', item)" v-if="small">
                    <strong>{{ item.Tracktitle }}</strong> <span>-</span> <span>
                        {{ item.Artist }}
                    </span>
                </a>

                <a href="#" class="open-track" @click="emit('trackSelect', item)" v-else>

                    <div class="tracktitle">
                        <strong>
                            {{ item.Tracktitle }}
                        </strong>
                    </div>
                    <div class="artist">{{ item.Artist }}</div>
                    <div v-if="item.CatalogNo != null" class="catalogId">{{ item.CatalogNo }}</div>
                    <small class="catalogId" v-else-if="item.CatalogNo == null">No catalog release</small>
                </a>

            </div>
        </div>

</template>

<style lang="css" scoped>
a:any-link {
    color: unset;
    text-decoration-line: unset;
}

a:hover {
    text-decoration-line: underline;
}

.browse-list:not(.small) {
    display: grid;
    justify-items: stretch;
    gap: .5em;
}

.browse-list .row .cover-area img {
    width: 64px;
    height: 64px;
}

.browse-list:not(.small) .row {
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

.add-action {}
</style>