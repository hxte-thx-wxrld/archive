<script lang="ts">
import { getDevPrefix } from '../main'
import MockTracks from '../mock/all_tracks.json'

export default {
    async setup() {
        const req = await fetch(getDevPrefix() + "/api/release/")
        const data = await req.json()

        console.log(data)

        return {
            data
        }
    },
    data() {
        return {}
    }
}
</script>

<template>
    <div class="browse-list">
        <div class="row" v-for="(item, index) in data" v-if="data != null">
            <div class="cover-area">
                <img :src="'http://s3.rillo.internal:8333' + item.CoverUrl">
            </div>
            
            <div class="name-area">
                    <RouterLink class="open-track" :to="'/catalog/releases/' + item.ReleaseId">
                    <div class="tracktitle">
                        <strong>
                            {{ item.Name }}
                        </strong>
                    </div>
                    <div class="artist">{{ item.Artist }}</div>
                    <div v-if="item.CatalogId" class="release-code">{{ item.CatalogId }}</div>
                    <small class="catalogId" v-else>No catalog release</small>
                </RouterLink>
            </div>
        </div>
        <div v-else>No Data</div>
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