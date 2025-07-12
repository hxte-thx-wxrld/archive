<script lang="ts" setup>
import { ref, useTemplateRef } from 'vue';
import AssignedArtistsPicker from '../AssignedArtistsPicker.vue';
import EditableText from '../EditableText.vue';
import UploadComp from '../UploadComp.vue';

const showfinishedMessage = ref<boolean>(false);
const form = useTemplateRef("track_upload_form")

const uploadTrack = async (event: SubmitEvent) => {
    const fdata = new FormData((event.target as HTMLFormElement))


    const track = await fetch("/api/inbox", {
        method: "POST",
        credentials: "include",
        body: fdata

        //body: JSON.stringify(d)
    })

    if (track.ok) {
        const r = await track.json()
        console.log(r)
        form.value.reset()
        showfinishedMessage.value = true;

        setTimeout(() => {
            showfinishedMessage.value = false;
        }, 5000)
    }
}

</script>

<template>
    <form v-if="showfinishedMessage">
        <span class="full-width">Upload Successful!</span>
    </form>
    <form @submit.prevent="uploadTrack" ref="track_upload_form" v-else>

        <h2 class="full-width">Create new Track</h2>


        <label for="Trackname">Trackname:</label>
        <input type="text" class="reverse" name="Trackname">

        <label for="ReleaseDate">Release Date:</label>
        <EditableText :reverse-style="true" :edit="true" :value="String(new Date())" type="date" name="ReleaseDate">
        </EditableText>

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