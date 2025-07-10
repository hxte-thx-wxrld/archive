<script setup lang="ts">

</script>

<script lang="ts">
export default {
    data() {
        return {}
    },
    methods: {
        async create(event: SubmitEvent) {
            const body = new FormData(event.target as HTMLFormElement)
            const req = await fetch("/api/release", {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({
                    "Name": body.get("name")
                })
            })

        }
    }
}
</script>

<template>
    <form @submit.prevent="create">
        <p class="title">Create new Release</p>

        <label for="name">Release Name</label>
        <input type="text" name="name">

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