<script setup lang="ts">
import InboxListMixin from '../components/lists/InboxListMixin.vue';
import TrackListMixin from '../components/lists/TrackListMixin.vue';
import DatabaseList from '../components/mixins/DatabaseList.vue';

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
    <h1>About</h1>


</template>

<style scoped></style>
