<script lang="ts" setup>
import type { UploadedTrackResponse } from '../../types';
import AssignedArtistsPicker from '../AssignedArtistsPicker.vue';
import EditableText from '../EditableText.vue';
import UploadComp from '../UploadComp.vue';
</script>

<script lang="ts">
export default {
    methods: {
        async uploadTrack(event: SubmitEvent) {
            //const req = await fetch("")
            const fdata = new FormData((event.target as HTMLFormElement))
            /*const d = {
                "TrackTitle": fdata.get("TrackTitle"),
                "ArtistId": fdata.get("ArtistId"),
                "PublicUrl": fdata.get("PublicUrl"),
            };*/


            const track = await fetch("/api/track/", {
                method: "POST",
                credentials: "include",
                body: fdata

                //body: JSON.stringify(d)
            })

            if (track.ok) {
                const r = await track.json() as UploadedTrackResponse
                console.log(r)
                //this.$router.push("/catalog/tracks/" + r.TrackId)
            }
        }
    },
    computed: {

    }
}
</script>

<template>
    <form @submit.prevent="uploadTrack">

        <h2 class="full-width">Create new Track</h2>

        <label for="TrackTitle">Trackname:</label>
        <input type="text" class="reverse" name="TrackTitle">

        <label for="ReleaseDate">Release Date:</label>
        <EditableText :reverse-style="true" :edit="true" :value="String(new Date())" type="date" name="ReleaseDate"></EditableText>

        <label for="ArtistId">Artist:</label>
        <AssignedArtistsPicker />
        
        <label for="AudioFile">
            Audio File:
        </label>
        <input type="file" name="AudioFile">

<!--
    <div class="full-width">
        <UploadComp endpoint="/api/track/presign" mime="audio/x-wav" />
    </div>
    -->


        <input type="submit">
    </form>
</template>

<style scoped>
form {
    display: grid;
    grid-template-columns: 1fr 1fr;
}
form .full-width {
    grid-column-start: 1;
    grid-column-end: 3;
}
</style>