<script setup lang="ts">
</script>

<script lang="ts">

export default {
    data() {
        return {
            err: null
        }
    },
    methods: {
        async login(event: SubmitEvent) {
            this.err = null
            const data = new FormData(event.target as HTMLFormElement);

            this.$store.dispatch("login", data).then(() => {
                this.$emit("close")
            }).catch((err) => {
                this.err = "login unsuccessful"
            })
        }
    }
}
</script>

<template>
    <form @submit.prevent="login">
        <strong class="login-label whole-width">Login</strong>
        <p class="whole-width" v-if="err != null">{{ err }}</p>
        <label for="username">Username:</label>
        <input class="reverse" type="text" name="username">
        <label for="password">Password:</label>
        <input class="reverse"     type="password" name="password">

        <input type="submit">
    </form>
</template>

<style scoped>
.login-label {
    font-size: 1.5em;
    margin: 0 0 1em 0;
}

.whole-width {
    grid-column-start: 1;
    grid-column-end: 3;
}

form input[type="submit"] {
    margin: 1em 0 0 0;
    grid-column-start: 2;
    grid-column-end: 3;
}

form {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 5px;
}
</style>