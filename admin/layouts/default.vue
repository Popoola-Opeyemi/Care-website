<template>
  <div class="admin-dashboard">
    <custom-load v-if="loading" />
    <side-bar />
    <section class="admin-main">
      <top-navigation />
      <div class="pages-content">
        <nuxt />
      </div>
    </section>
    <!-- <div class="pages-content">
    </div>-->
  </div>
</template>

<script>
import SideBar from '@/components/widgets/SideBar'
import TopNavigation from '@/components/widgets/TopNavigation'
import CustomLoad from '@/components/CustomLoad'
import { mapState } from 'vuex'

export default {
  middleware: ['auth_status'],
  components: {
    SideBar,
    TopNavigation,
    CustomLoad
  },
  computed: {
    ...mapState(['loading'])
  },
  data() {
    return {}
  }

  // async mounted() {
  //   this.$store.dispatch('api/Initialize')
  // }
}
</script>
