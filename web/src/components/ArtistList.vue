<script setup lang="ts">
import { mapGetters } from 'vuex'
import CreateArtistDialog from './dialogs/ArtistCreateDialog.vue'
import { computed, ref } from 'vue';
import type { Artist, PaginatedArtistLookup } from '../types'
import { useStore } from 'vuex';

const data = ref<PaginatedArtistLookup>();
const req = await fetch("/api/artist?offset=0")
data.value = await req.json()
const store = useStore();
const isLoggedIn = computed(() => store.getters.isLoggedIn)
console.log(data.value)

</script>

<template>

    <div class="browse-list">
        <div class="row" v-for="(item, index) in data.Rows">
            <div class="cover-area">
                <img :src="$store.getters.s3Host + item.ArtistPicture">
            </div>

            <div class="name-area">

                    <a class="open-track" href="#" @click="$emit('artistSelect', item)">

                        <div class="tracktitle">
                            <strong>
                                {{ item.Name }}
                            </strong>
                        </div>
                    </a>

            </div>
        </div>
    </div>

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