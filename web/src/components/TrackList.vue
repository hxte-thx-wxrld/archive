<script setup lang="ts">
import { mapGetters } from 'vuex'
import { getDevPrefix, getS3Host } from '../main'
import Paginator from './Paginator.vue'
import type { Music } from '../types'
import { ref } from 'vue';

const data = ref<
    {
        Rows: Music[],
        FullLength: number,
    }>();

const page = ref<number>(0);

const props = defineProps<{
    showCover: boolean,
    small: boolean,
}>();

const emit = defineEmits<{
    'trackSelect': [Music]
}>()

async function reloadList(page) {
    const req = await fetch(getDevPrefix() + "/api/track/?offset=" + page)
    const j = await req.json();
    console.table(j.Rows);
    return j;
}

function switchPage(i: number) {
    console.log(i)
    page.value = i
    reloadList(page.value).then(d => {
        data.value = d
    })
}

data.value = await reloadList(page.value)

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
    <div class="browse-list" :class="{ 'small': small }">
        <div class="row" v-for="(item, index) in data.Rows" v-if="data != null">
            <div class="cover-area" v-if="showCover">
                <img :src="getS3Host() + item.CoverUrl">
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
        <div v-else>No Data</div>
        <Paginator :pagecount="data.FullLength" :page="page" @switchPage="switchPage" />
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