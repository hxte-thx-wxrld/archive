<script setup lang="ts">
import { ref } from 'vue';
import EditableText from '../EditableText.vue';
import { useRouter } from 'vue-router';

const name = ref("")
const router = useRouter();

async function create(event: SubmitEvent) {
    const body = new FormData(event.target as HTMLFormElement)
    console.log(body.get("Name"))
    const req = await fetch("/api/artist", {
        method: "POST",
        body: JSON.stringify({
            "Name": body.get("Name")
        })
    })
    
    const res = await req.json()
    console.log(res)

    router.push("/catalog/artists/" + res.ArtistId)
}

</script>

<template>
    <form @submit.prevent="create">
        <p class="title">Create new Artist</p>

        <label for="name">Artist Name</label>
        <EditableText :edit="true" :reverse-style="true" :value="name" type="text" name="Name"></EditableText>

        <input type="submit">
    </form>
</template>

<style scoped>
form {
    display: grid;
    grid-template-columns: 1fr 1fr;
}

form .title {
    grid-column-start: 1;
    grid-column-end: 3;
}
</style>