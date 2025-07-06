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