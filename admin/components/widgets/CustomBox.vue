<template>
  <div class="_customm-box" :class="EditMode ? '_customm-box-disabled' : ''">
    <div class="header">
      <div class="__left">{{title}}</div>
      <div class="__right">
        <slot name="button"></slot>
      </div>
    </div>

    <slot></slot>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      default: () => ''
    },
    EditMode: {
      type: Boolean,
      default: () => false
    }
  }
}
</script>

<style lang="scss" scoped>
</style>