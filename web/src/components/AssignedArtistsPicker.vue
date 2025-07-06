<script lang="ts" setup>

</script>

<script lang="ts">

</script>

<template>
        <select name="ArtistId">
        <option v-for="(artist, id) in $store.state.AssignedArtists" :value="artist.ArtistId">
            {{ artist.Name }}
        </option>
    </select>
</template>

<style scoped>
    select {
        border: 0;
        background-color: inherit;
        -webkit-appearance: none;
        color: inherit;
        font-family: inherit;
        font-size: inherit;
    }
</style>