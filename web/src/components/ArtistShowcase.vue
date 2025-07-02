<script lang="ts">
export default {
    props: ['artistId'],
    async setup(props) {
        const req = await fetch("http://localhost:8080/api/artist/" + props.artistId)
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