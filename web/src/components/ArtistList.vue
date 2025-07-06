<script setup lang="ts">
import { mapGetters } from 'vuex'
import { getDevPrefix } from '../main'
import CreateArtistDialog from './CreateArtistDialog.vue'
import { ref } from 'vue';
import type { Artist } from '../types'

const data = ref<[Artist]>(null);
const req = await fetch(getDevPrefix() + "/api/artist/")
data.value = await req.json()

console.log(data.value)
</script>
<script lang="ts">
export default {
    components: {
        CreateArtistDialog: CreateArtistDialog
    },
    methods: {
        openNew() {
            (this.$refs.new_artist as HTMLDialogElement).showModal()
        }
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
        <div class="row" v-for="(item, index) in data">
            <div class="cover-area">
                <img :src="'http://s3.rillo.internal:8333' + item.ArtistPicture">
            </div>

            <div class="name-area">
                <RouterLink class="open-track" :to="'/catalog/artists/' + item.ArtistId">
                    <div class="tracktitle">
                        <strong>
                            {{ item.Name }}
                        </strong>
                    </div>
                </RouterLink>
            </div>
        </div>
    </div>
    <dialog v-if="isLoggedIn" ref="new_artist">
        <CreateArtistDialog />
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