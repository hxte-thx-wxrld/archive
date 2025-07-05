<script lang="ts">
import { getDevPrefix } from '../main';

export default {
    methods: {
        async login(event: SubmitEvent) {
            const data = new FormData(event.target as HTMLFormElement);

            const req = await fetch(getDevPrefix() + "/api/login", {
                method: "POST",
                credentials: 'include',
                body: JSON.stringify({
                    username: data.get("username"),
                    password: data.get("password"),
                })
            })

            if (req.ok) {
                console.log(req.ok)
                //const json = await req.json();

                const whoami = await fetch(getDevPrefix() + "/api/me", {
                    method: "GET",
                    credentials: 'same-origin'
                })

                console.log(await whoami.json());
            }
        }
    }
}
</script>

<template>
    <form @submit.prevent="login">
        <strong class="login-label">Login</strong>
        <label for="username">Username:</label>
        <input type="text" name="username">
        <label for="password">Password:</label>
        <input type="password" name="password">

        <input type="submit">
    </form>
</template>

<style>
.login-label {
    grid-column-start: 1;
    grid-column-end: 3;
    font-size: 1.5em;
    margin: 0 0 1em 0;
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