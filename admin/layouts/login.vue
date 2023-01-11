<template>
  <div class="login-page">
    <!-- <custom-load /> -->
    <div class="columns is-gapless">
      <div class="column hero is-fullheight left-content"></div>
      <div class="column hero is-fullheight">
        <div class="right-column-content">
          <div class="logo has-text-centered">
            <img src="/img/logo.png" alt="'Motorcycle by kesaryvamshi from the Noun Project'" />
            <div class="admin-title is-blue">Prudent Domiciliary Care Admin</div>
          </div>
          <div class="login-form">
            <nuxt />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import CustomLoad from '@/components/CustomLoad'
import Mix from '@/mixins/index'
export default {
  mixins: [Mix],
  components: {
    CustomLoad
  }
}
</script>

<style lang="scss">
</style>