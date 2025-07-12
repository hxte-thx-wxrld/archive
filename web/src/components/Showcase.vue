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

<style scoped>
/* iphone */
@media only screen and (max-width: 768px) {
  .showcase {
    grid-template-columns: 1fr;
    justify-items: center;
  }

  .showcase .meta {
    display: grid;
    grid-template-columns: 1fr;
    margin: 1em;
    justify-self: left;
  }

  .showcase .left img {
    width: 100%;
    height: auto;
    border: unset;
  }

  .showcase .meta :nth-child(2n) {
    padding: 0 0 1em 0;
  }

  .root {
    padding: 1em;
  }

  .top .logo {
    width: 100%;
  }

}

/*desktop*/
@media only screen and (min-width: 768px) {
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
}
</style>