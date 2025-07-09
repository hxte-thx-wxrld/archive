<script setup lang="ts">
//import UploadComp from './components/UploadComp.vue'
import { mapGetters } from 'vuex';
import { ref } from 'vue';
import ArtistList from '../components/ArtistList.vue';
import ReleaseList from '../components/ReleaseList.vue';
import TrackList from '../components/TrackList.vue';
import TrackUploadDialog from '../components/dialogs/TrackUploadDialog.vue';
import ArtistCreateDialog from '../components/dialogs/ArtistCreateDialog.vue';
import ReleaseCreateDialog from '../components/dialogs/ReleaseCreateDialog.vue';
import DatabaseList from '../components/mixins/DatabaseList.vue';
import DatabaseListItem from '../components/lists/InboxListMixin.vue';
import TrackListMixin from '../components/lists/TrackListMixin.vue';
import InboxIndicator from '../components/InboxIndicator.vue';

const showSearchbar = ref(false)
</script>

<script lang="ts">
export default {
    computed: {
        ...mapGetters(["isLoggedIn", "isAdmin"])
    },
    methods: {
        getSubmode() {
            if (!this.$route.params.mode) {
                return "tracks";
            } else return this.$route.params.mode;
        },
        openCreateTrack(event) {
            console.log(event);
            (this.$refs.new_track as HTMLDialogElement).showModal();
        },
        closeCreateTrack(event) {
            (this.$refs.new_track as HTMLDialogElement).close();
        },
        openCreateArtist() {
            (this.$refs.new_artist as HTMLDialogElement).showModal()
        },
        openCreateRelease() {
            (this.$refs.new_release as HTMLDialogElement).showModal()

        },
        openTrack(item) {
            console.log(item)
            this.$router.push('/catalog/tracks/' + item.TrackId)
            //:to="'/catalog/tracks/' + item.TrackId"
        },
        openRelease(item) {
            console.log(item)
            this.$router.push('/catalog/releases/' + item.ReleaseId)
            //:to="'/catalog/tracks/' + item.TrackId"
        },
        openArtist(item) {
            console.log(item)
            this.$router.push('/catalog/artists/' + item.ArtistId)
            //:to="'/catalog/tracks/' + item.TrackId"
        },
        openInbox() {
            alert("inbox")
        }
    }
}
</script>

<template>
    <div class="top">

        <h1>Catalog</h1>

        <nav class="mode">
            <RouterLink class="icon tracks" :class="getSubmode() == 'tracks' ? 'active' : ''" to="/catalog/tracks">
                Tracks</RouterLink>
            <RouterLink class="icon releases" :class="getSubmode() == 'releases' ? 'active' : ''"
                to="/catalog/releases">Releases</RouterLink>
            <RouterLink class="icon artists" :class="getSubmode() == 'artists' ? 'active' : ''" to="/catalog/artists">
                Artists</RouterLink>
            <a href="#" @click.prevent="showSearchbar = !showSearchbar">{{ showSearchbar ? "-" : "+" }}</a>
        </nav>
        <div class="tools" :class="showSearchbar ? '' : 'hide'">
            <input class="search" type="text">
            <input class="go" type="submit">
        </div>
    </div>

    <div v-if="getSubmode() == 'tracks'">
        <div class="catalog-toolbar" v-if="isLoggedIn">
            <a class="add-action" href="#" @click.prevent="openCreateTrack" v-if="isLoggedIn">Add New</a>
            <!-- <a class="open-inbox" href="#" data-count="1" @click.prevent="$router.push('/catalog/track-inbox')" v-if="isAdmin">Inbox</a> -->
            <Suspense v-if="isAdmin">
                <InboxIndicator></InboxIndicator>
            </Suspense>
        </div>
        <Suspense>
            <TrackList :small="false" :show-cover="true" @trackSelect="openTrack" />

            <template #fallback>
                Loading...
            </template>
        </Suspense>

    </div>
    <div v-if="getSubmode() == 'track-inbox'">
        <h1>Inbox</h1>
        <RouterLink class="icon tracks" to="/catalog/tracks" v-if="isAdmin">Go back</RouterLink>
        <Suspense>
            <DatabaseList api-endpoint="/api/inbox/" :small="true">
                <DatabaseListItem></DatabaseListItem>
                <!--<div class="row" v-for="(item, index) in $refs.inboxlist.Rows" v-if="data != null">
                    {{ item }}
                </div>-->

            </DatabaseList>
        </Suspense>
    </div>
    <div v-else-if="getSubmode() == 'releases'">
        <div class="catalog-toolbar">

            <a class="add-action" href="#" @click.prevent="openCreateRelease" v-if="isLoggedIn">Add New</a>
        </div>
        <Suspense>
            <ReleaseList @releaseSelect="openRelease" />
            <template #fallback>
                Loading...
            </template>
        </Suspense>
    </div>
    <div v-else-if="getSubmode() == 'artists'">
        <div class="catalog-toolbar">
            <a href="#" @click.prevent="openCreateArtist" v-if="isLoggedIn">Add New</a>
        </div>
        <Suspense>
            <ArtistList @artistSelect="openArtist" />
            <template #fallback>
                Loading...
            </template>
        </Suspense>
    </div>
    <dialog v-if="isLoggedIn" ref="new_track">
        <TrackUploadDialog />
    </dialog>

    <dialog v-if="isLoggedIn" ref="new_release">
        <ReleaseCreateDialog />
    </dialog>

    <dialog v-if="isLoggedIn" ref="new_artist">
        <ArtistCreateDialog />
    </dialog>
    <!-- <UploadComp /> -->
</template>

<style scoped>
.catalog-toolbar {
    display: flex;
    padding: 1em 0;
    justify-content: space-between;
}

h1 {
    padding: 1em 0;
}

.hide {
    display: none !important;
}

nav {
    display: grid;
}

.tools {
    border: 1px solid white;
    padding: 1em;
    margin: 0 0 1em 0;
    display: grid;
    grid-template-columns: 1fr 100px;
}

.tools .search {
    border: unset;
    border-bottom: 3px solid white;
    background-color: transparent;
    color: white;
    font-family: "Apricot200L";
}

.tools .go:active {
    background-color: black;
    color: white;
}

.tools .go {
    justify-self: stretch;
    border: unset;
    background-color: white;
    color: black;
    font-family: "Apricot200L";
    padding: 1em;
    margin: 0 .5em;
}

.top {
    position: sticky;
    top: 0;
    background-color: black;
}

nav.mode {
    padding: 1em 0;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 32px;
}

nav.mode a {
    justify-self: stretch;
}

nav.mode a.active {
    background-color: grey;
    color: white;
}

nav.mode a.active::before {
    filter: invert(1)
}

nav.mode>* {
    background-color: white;
    color: black;
    padding: .5em;
    font-family: "Apricot200L";
}
</style>
