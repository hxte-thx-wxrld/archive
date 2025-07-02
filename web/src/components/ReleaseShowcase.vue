<script lang="ts">

export default {
    props: ['releaseId'],
    async setup(props) {
        console.log(props.releaseId)
        const req = await fetch("http://localhost:8080/api/release/" + props.releaseId)
        const data = await req.json()

        console.log(data)
        return {
            data
        }
    }
}
</script>

<template>
    <div class="showcase">
        <div class="cover-area">
            <img :src="'http://s3.rillo.internal:8333' + data.CoverUrl">
        </div>
        <div class="meta">
            <h1>{{ data.Name }}</h1>
            <div></div>
            <strong>Release Date:</strong><span>{{ data.ReleaseDate }}</span>
            <strong>Release-Code:</strong><span class="release-code">{{ data.CatalogId }}</span>
            <strong>ISRC-Code:</strong><span>{{ data.Isrc }}</span>
            <span>Play: </span><a v-if="data.PublicUrl != null">Play</a><small v-else>stream not available</small>
            <h2>Track List</h2>
            <ol>
                <li v-for="(music, index) in data.RelatedMusic">
                    <div>
                        <p>
                            <strong>
                                {{ music.Name }}
                            </strong>
                        </p>
                    </div>
                    <div>
                        <p>
                            {{ music.ArtistName }}
                        </p>
                    </div>
                    </li>
            </ol>

        </div>

    </div>
</template>

<style scoped>
.showcase .cover-area img {}

.showcase .meta {
    display: grid;
    align-self: flex-start;
}

.showcase .meta li {
    padding: 0 0 0 0;
  }
</style>