<script setup lang="ts">
const props = defineProps<{
    page: number,
    pagecount: number
}>();

const emit = defineEmits<{
    'switchPage': [page: number]
    
}>();
function switchPage(i: number) {

    emit('switchPage', i - 1)
}
</script>
<script lang="ts">
export default {
    methods: {

    }
}
</script>

<template>
    <div class="p-container">
        <a href="#" :class="(i-1) == page ? 'selected' : ''" v-for="i in pagecount+1" @click.prevent="(event) => switchPage(i)">{{ i }}</a>
    </div>
</template>

<style scoped>
.p-container {
    display: flex;
    flex-direction: row;
    gap: .2em;
}

.p-container a {
    background-color: white;
    color: black;
    padding: .2em;
}

.selected {
    background-color: black !important;
    color: white !important;
}
</style>