<script>
export default {

}
</script>

<template>
    <div class="showcase">
        <div class="left">
            <slot name="left"></slot>
        </div>
        <div class="meta">
            <slot></slot>
        </div>
    </div>
</template>

<style>
.showcase {
  display: grid;
  grid-template-columns: 1fr 1fr;
}

.showcase:has(.left .cover) {
  grid-template-columns: 256px 1fr;
}

.showcase .left {
  display: flex;
  flex-direction: column;
}

.showcase .meta {
  display: grid;
  grid-template-columns: 1fr 1fr;
  align-self: flex-start;
}

.showcase .meta .title {
  grid-column-start: 1;
  grid-column-end: 3;
    padding: 0 0 1em 0;
}
</style>