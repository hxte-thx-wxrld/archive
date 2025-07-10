<script setup lang="ts">
import { computed, ref } from 'vue';
import type { PaginatedInboxItems } from '../types';
import Paginator from './Paginator.vue';


async function reloadList(page): Promise<PaginatedInboxItems> {
    const req = await fetch("/api/inbox/?offset=" + page)
    const j = await req.json();
    console.table(j.Rows);
    return j;
}

async function acceptItem(id: string) {
    const req = await fetch("/api/inbox/" + id + "/accept", {
        method: "POST"
    })
    if (req.ok) {
        data.value = await reloadList(page.value);
    }
}

async function deleteItem(id: string) {
    const req = await fetch("/api/inbox/" + id + "/delete", {
        method: "POST"
    })
    if (req.ok) {
        data.value = await reloadList(page.value);
    } else {
        const j = await req.json()
        console.error(j)
    }
}

const page = ref<number>(0);
const data = ref(await reloadList(page.value))
console.log(data)

function switchPage(i: number) {
    console.log(i)
    page.value = i
    reloadList(page.value).then(d => {
        data.value = d
    })
}

const hasData = computed(() => data.value.Rows.length > 0)
</script>


<template>
    
    <ul>
        <li v-for="d in data.Rows">
            <div>{{ d.Trackname }}</div>
            <div class="tools">
                <a href="#" @click.prevent="acceptItem(d.UploadId)">Accept</a>
                <a href="#" @click.prevent="deleteItem(d.UploadId)">Delete</a>
            </div>
        </li>

        <p v-if="!hasData">Your Inbox is Empty</p>
    </ul>
    <Paginator :page="page" :pagecount="data.PageCount" @switch-page="switchPage"></Paginator>
</template>

<style scoped>
li {
    display: flex;
    justify-content: space-between;
}

ul {
    display: flex;
    flex-direction: column;
    gap: 1em;
}

.tools {
    display: flex;
    gap: 1em;
}

.tools a {
    padding: .5em;
    background-color: white;
    color: black;
}
</style>