<script setup lang="ts">
//import UploadComp from './components/UploadComp.vue'
import { mapGetters } from 'vuex';
import { computed, ref, useTemplateRef } from 'vue';
import ArtistList from '../components/ArtistList.vue';
import ReleaseList from '../components/ReleaseList.vue';
import TrackList from '../components/TrackList.vue';
import TrackUploadDialog from '../components/dialogs/TrackUploadDialog.vue';
import ArtistCreateDialog from '../components/dialogs/ArtistCreateDialog.vue';
import ReleaseCreateDialog from '../components/dialogs/ReleaseCreateDialog.vue';
import InboxList from '../components/InboxList.vue';
import InboxIndicator from '../components/InboxIndicator.vue';
import { useStore } from 'vuex';
import { useRoute, useRouter } from 'vue-router';

const showSearchbar = ref(false)

const store = useStore()
const isLoggedIn = computed(() => store.getters.isLoggedIn)
const isAdmin = computed(() => store.getters.isAdmin)
const route = useRoute()
const router = useRouter()
const new_track = useTemplateRef("new_track")
const new_artist = useTemplateRef("new_artist")
const new_release = useTemplateRef("new_release")

function getSubmode() {
    if (!route.params.mode) {
        return "tracks";
    } else return route.params.mode;
}
function openCreateTrack(event) {
    console.log(event);
    (new_track.value as HTMLDialogElement).showModal();
}
function closeCreateTrack(event) {
    (new_track.value as HTMLDialogElement).close();
}
function openCreateArtist() {
    (new_artist.value as HTMLDialogElement).showModal()
}
function openCreateRelease() {
    (new_release.value as HTMLDialogElement).showModal()

}
function openTrack(item) {
    console.log(item)
    router.push('/catalog/tracks/' + item.TrackId)
    //:to="'/catalog/tracks/' + item.TrackId"
}
function openRelease(item) {
    console.log(item)
    router.push('/catalog/releases/' + item.ReleaseId)
    //:to="'/catalog/tracks/' + item.TrackId"
}
function openArtist(item) {
    console.log(item)
    router.push('/catalog/artists/' + item.ArtistId)
    //:to="'/catalog/tracks/' + item.TrackId"
}
function openInbox() {
    alert("inbox")
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
            <!--
                <a href="#" @click.prevent="showSearchbar = !showSearchbar">{{ showSearchbar ? "-" : "+" }}</a>
                -->
        </nav>
        <!--
            <div class="tools" :class="showSearchbar ? '' : 'hide'">
                <input class="search" type="text">
                <input class="go" type="submit">
            </div>
            -->
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
        <div class="inbox-header">
            <span class="title">Inbox</span>
            <RouterLink class="icon tracks" to="/catalog/tracks" v-if="isAdmin">Go back</RouterLink>
        </div>

        <Suspense>
            <!--
            <DatabaseList api-endpoint="/api/inbox/" :small="true">
                <DatabaseListItem></DatabaseListItem>
                <div class="row" v-for="(item, index) in $refs.inboxlist.Rows" v-if="data != null">
                    {{ item }}
                </div>
                
            </DatabaseList>
            -->
            <InboxList></InboxList>
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
.inbox-header {
    display: flex;
    justify-content: space-between;
}

.inbox-header .title {
    font-size: 2em;
}

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
    font-family: "PixelifySans";
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
    font-family: "Silkscreen";
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
    grid-template-columns: 1fr 1fr 1fr;
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
    font-family: "Silkscreen";
}
</style>
