<script lang="ts">
import { getDevPrefix } from '../main'

export default {
    props: ['artistId'],
    async setup(props) {
        const req = await fetch(getDevPrefix() + "/api/artist/" + props.artistId)
        const data = await req.json()

        console.log(data)
        return {
            data
        }
    }
}
</script>

<template>
    <div>
        <h1>{{ data.Name }}</h1>
    </div>
</template>