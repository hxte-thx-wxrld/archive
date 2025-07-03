<script lang="ts">
export default {
    data() {
        return {
            active: false,
            files: []
        }
    },
    methods: {
        setActive(event: DragEvent) {
            console.log(event)

            this.active = true
        },
        setInactive(_event: DragEvent) {
            //console.log(event)
            this.active = false
        },
        onDrop(event: DragEvent) {
            if (event.dataTransfer) {
                this.files.push(...event.dataTransfer.files)
            }
            this.active = false
        }
    }
}
</script>

<template>
    <div class="upload-box" :data-active="active" @dragenter.prevent="setActive" @dragover.prevent="setActive" @dragleave.prevent="setInactive" @drop.prevent="onDrop">
        
    </div>
    <ul>
        <li v-for="(file, index) in files">{{ file.name }}</li>
    </ul>
    <form>
        
    </form>
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
