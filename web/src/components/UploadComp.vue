<script lang="ts">
export default {
    data() {
        return {
            active: false
        }
    },
    methods: {
        setActive(event: DragEvent) {
            console.log(event)
            debugger
            this.active = true
        },
        setInactive(_event: DragEvent) {
            //console.log(event)
            this.active = false
        },
        onDrop(event: DragEvent) {
            if (event.dataTransfer) {
                console.log(event.dataTransfer.files)
            }
            this.active = false
        }
    }
}
</script>

<template>
    <div class="upload-box" :data-active="active" @dragenter.prevent="setActive" @dragover.prevent="setActive" @dragleave.prevent="setInactive" @drop.prevent="onDrop">
        
    </div>
</template>

<style scoped>
.upload-box {
    width: 100px;
    height: 100px;
    background-color: whitesmoke;
}

[data-active=true] {
    background-color: yellow;
}
</style>
