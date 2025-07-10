<script setup lang="ts">
import { mapGetters } from 'vuex'
import CreateTrackDialog from '../dialogs/TrackUploadDialog.vue'

import { inject, ref } from 'vue';
import type { Music } from '../../types';

const props = defineProps<{
    small: boolean
    showCover: boolean
}>()

const Rows = inject<Music[]>("Rows")
const emit = defineEmits<{
    'trackSelect': Music[]
}>()

</script>

<script lang="ts">
export default {
    methods: {

    },
    computed: {
        ...mapGetters(["isLoggedIn"])
    },
    data() {
        return {}
    }
}
</script>

<template>


        <ol>

            <li v-for="(item, index) in Rows">
                {{ item.Tracktitle }}
            </li>
        </ol>

</template>

<style lang="css" scoped>
a:any-link {
    color: unset;
    text-decoration-line: unset;
}

a:hover {
    text-decoration-line: underline;
}


.add-action {}
</style>