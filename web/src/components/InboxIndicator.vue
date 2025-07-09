<script setup lang="ts">
import { ref } from 'vue';
import type { PaginatedInboxItems } from '../types';

const count = ref(0)
const r = await fetch("/api/inbox/count", {
    credentials: 'include'
})
const req = await r.json() as PaginatedInboxItems
count.value = req.Count

</script>

<template>
    <a class="open-inbox" href="#" :data-count="count" @click.prevent="$router.push('/catalog/track-inbox')">Inbox</a>
</template>

<style scoped>
a.open-inbox::before {
    content: "(" attr(data-count) ")";
    font-size: .5em;
}
</style>