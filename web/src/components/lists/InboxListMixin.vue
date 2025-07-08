<script setup lang="ts">
import { inject } from 'vue';
import type { InboxItem } from '../../types';

const Rows = inject<InboxItem[]>("Rows");
const emit = defineEmits<{
    'trackSelect': [string]
}>()
</script>

<script lang="ts">
export default {
}
</script>

<template>
    <div v-for="item in Rows">
        <p>{{ item.Trackname }}</p>
        <a href="#" @click.prevent="emit('trackSelect', item.UploadId)">accept</a>
    </div>
    

</template>