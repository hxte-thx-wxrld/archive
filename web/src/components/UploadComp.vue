<script setup lang="ts">
import { hash } from '../hash';
const props = defineProps<{
    endpoint: string,
    mime: string
}>()
</script>

<script lang="ts">
export default {
    data() {
        return {
            active: false,
            file: null,
            presigned_url: null
        }
    },
    methods: {
        setActive(event: DragEvent) {


            this.active = true
        },
        setInactive(_event: DragEvent) {
            //console.log(event)
            this.active = false
        },
        async onDrop(event: DragEvent) {
            /*if (event.dataTransfer) {
                this.files.push(...event.dataTransfer.files)
            }*/
            const file = event.dataTransfer.files[0];
            if (file.type != this.mime) {
                throw new Error(`Only ${this.mime} files are supported`)
            }

            const f = await file.bytes()

            //const res = await this.createPresignedUrl(h, f.length);


            /*const upload = await fetch(res.Url, {
                credentials: "same-origin",
                method: "PUT",
                body: file,
                headers: {
                    "ContentLength": String(f.length)
                }
            })

            if (upload.ok) {

                this.file = res;
                this.active = false
            }*/

        },
        async createPresignedUrl(h, len) {
            const req = await fetch(this.$props.endpoint + "?hash=" + btoa(h), {
                method: "POST",
                credentials: "include",
            });

            return await req.json();
        }
    }
}
</script>

<template>
    <div class="upload-box" :data-active="active" @dragenter.prevent="setActive" @dragover.prevent="setActive"
        @dragleave.prevent="setInactive" @drop.prevent="onDrop">
        <input type="hidden" name="PublicUrl" :value="file.Url" v-if="file != null">
        <input type="hidden" name="TrackId" :value="file.TrackId" v-if="file != null">
    </div>
</template>

<style scoped>
.upload-box {
    width: 100%;
    height: 8em;
    background-color: whitesmoke;
}

[data-active=true] {
    background-color: yellow;
}
</style>
