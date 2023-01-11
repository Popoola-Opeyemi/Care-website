<template>
  <div>
    <div class="card">
      <div class="card-image">
        <figure class="image is-4by3">
          <img :src="image" />
        </figure>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    image: {
      type: String,
      default: ''
    }
  }
}
</script>

<style>
</style>