<script setup lang="ts">
const props = defineProps<{
    edit: boolean
    reverseStyle: boolean
    value: string
    type: string
    name: string
}> ();
</script>

<script lang="ts">
export default {
    methods: {

    }
}
</script>

<template>
<span v-if="!edit">{{ value }}</span>
<input :class="{'reverse': reverseStyle}" :name="name" :type="type" :value="value" v-else>
</template>

<style scoped>

</style>
