<script setup lang="ts">
import InboxListMixin from '../components/lists/InboxListMixin.vue';
import TrackListMixin from '../components/lists/TrackListMixin.vue';

async function acceptInboxItem(item) {
    console.log(item)
    const req = await fetch("/api/inbox/" + item + "/accept", {
        credentials: "include"
    })


    console.log(req)
}
</script>

<script lang="ts">
export default {

}
</script>

<template>
    <h1>About / Typography Test</h1>

    <h1>Hello World</h1>
    <h2>Hello World</h2>
    <h3>Hello World</h3>
</template>

<style scoped></style>
