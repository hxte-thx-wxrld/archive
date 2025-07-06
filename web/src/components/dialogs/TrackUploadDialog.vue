<script lang="ts" setup>
import AssignedArtistsPicker from '../AssignedArtistsPicker.vue';
import UploadComp from '../UploadComp.vue';
</script>

<script lang="ts">
export default {
    methods: {
        async uploadTrack(event: SubmitEvent) {
            //const req = await fetch("")
            const fdata = new FormData((event.target as HTMLFormElement))
            const d = {
                "TrackTitle": fdata.get("TrackTitle"),
                "ArtistId": fdata.get("ArtistId"),
                "PublicUrl": fdata.get("PublicUrl"),
            };

            console.log(d);

            const track = await fetch("/api/track/" + fdata.get("TrackId"), {
                method: "POST",
                credentials: "include",
                body: JSON.stringify(d)
            })

            if (track.ok) {
                this.$router.push("/catalog/tracks/" + fdata.get("TrackId"))
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
        <input type="text" name="TrackTitle">

        <label for="ReleaseDate">Release Date:</label>
        <input type="date" name="ReleaseDate" :value="new Date()">

        <label for="ArtistId">Artist:</label>
        <AssignedArtistsPicker />

        <div class="full-width">
            <UploadComp endpoint="/api/track/presign" mime="audio/x-wav" />
        </div>

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