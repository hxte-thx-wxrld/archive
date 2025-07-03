<script lang="ts">
import { getDevPrefix } from '../main'


export default {
    props: ['releaseId'],
    async setup(props) {
        console.log(props.releaseId)
        const req = await fetch(getDevPrefix() + "/api/release/" + props.releaseId)
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
            <h1 class="title">{{ data.Name }}</h1>
            <strong>Release Date:</strong><span>{{ data.ReleaseDate }}</span>
            <strong>Release-Code:</strong><span class="release-code">{{ data.CatalogId }}</span>

            <strong>ISRC-Code:</strong><span>{{ data.Isrc ?? "n. A." }}</span>
            <span>Play: </span><a v-if="data.PublicUrl != null">Play</a><small v-else>stream not available</small>

            <div class="tracklist">

                <h2>Track List</h2>
                <ol>
                    <li v-for="(music, index) in data.RelatedMusic">
                        <a :href="'/track/' + music.TrackId">
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
                        </a>
                    </li>
                </ol>
            </div>
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

.tracklist {
    padding: 2em 0;
    grid-column-start: 1;
    grid-column-end: 3;
}
.tracklist h2 {
    margin: 1em 0;
}

.tracklist li {
    margin: .5em;
}
</style>