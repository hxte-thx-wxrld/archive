<script setup lang="ts">
import { useRouter } from 'vue-router';
import EditableText from '../EditableText.vue';
import { ref } from 'vue';


const router = useRouter()

const name = ref("");

async function create(event: SubmitEvent) {
    const body = new FormData(event.target as HTMLFormElement)
    const req = await fetch("/api/release", {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({
            "Name": body.get("Name")
        })
    })

    if (req.ok) {
        const res = await req.json()
        router.push("/catalog/releases/" + res.ReleaseId)
    }
}

</script>

<template>
    <form @submit.prevent="create">
        <p class="title">Create new Release</p>

        <label for="name">Release Name</label>
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