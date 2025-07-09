<script lang="ts" setup>
import { ref } from 'vue';
import type { Music } from '../../types';
import TrackList from '../TrackList.vue';

const selection = ref<Music[]>([])
const assignform = ref<HTMLFormElement>();

const emit = defineEmits<{
    'save': [string[]],
    'close': []
}>()

function addToSelection(item) {
    selection.value.push(item)
}

function save(event: MouseEvent) {
    const fdata = new FormData(assignform.value as HTMLFormElement);
    const ids = fdata.getAll("TrackId").map((v, i) => v.toString());
    emit('save', ids)

    //emit('save', { data: selection.value })
}

function cancel() {
    emit('close')
}
</script>
<script lang="ts"></script>

<template>
    <div class="container">
        <div class="left">
            <form @submit.prevent="save" ref="assignform">
                <h2>Selection</h2>
                <ul>
                    <li v-for="(item, index) in selection.filter((v, index) => {

                        // implement pagination here
                        return true
                    })" :key="index">
                        <div>
                            <p>
                                <strong>
                                    {{ item.Tracktitle }}
                                </strong>
                            </p>
                            <p>
                                {{ item.Artist }}
                            </p>
                            <input type="hidden" name="TrackId" :value="item.TrackId">
                        </div>
                    </li>
                </ul>
            </form>
        </div>
        <div class="right">
            <h2>Available Tracks</h2>
            <TrackList :show-cover="false" :small="true" @trackSelect="addToSelection"></TrackList>
        </div>
        
        <div class="actionbuttons">
            <button type="button" @click="save">Save</button>
            <button type="button" @click="cancel">Cancel</button>
        </div>
    </div>
</template>

<style scoped>
.actionbuttons {
    display: flex;
    gap: 1em;
}
.container {
    padding: 1em;
    display: grid;
    grid-template-columns: 1fr 1fr;
}
button {
    background-color: black;
    color: white
}
</style>