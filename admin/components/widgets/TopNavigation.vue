<template>
  <div class="topnavigation">
    <div class="left">
      <div class="pg-title">{{ pageTitle }}</div>
    </div>
    <div class="right">
      <!-- <i
        class="admin-icon mdi mdi-account"
        title="Edit Profile"
        @click="editProfile"
      />-->
      <i class="admin-icon mdi mdi-logout" title="Log Out" @click="logout()" />
    </div>
  </div>
</template>
<script>
import { mapState } from 'vuex'
import LogoutModal from '@/components/widgets/LogoutModal'

export default {
  components: {
    LogoutModal
  },
  computed: {
    ...mapState(['pageTitle'])
  },
  methods: {
    handleMenu() {
      this.$store.commit('savetogglemenu')
    },
    editProfile() {},
    logout(id) {
      let that = this
      this.$buefy.modal.open({
        parent: this,
        component: LogoutModal,
        hasModalCard: true,
        customClass: '',
        trapFocus: true,
        events: {
          logout(event) {
            that.logAction()
          }
        }
      })
    },
    async logAction() {
      let param = {
        url: `auth/logout`
      }
      let page = await this.$store.dispatch('Get', param)
      if (page.done) {
        this.$router.push('/')
      }
    }
  }
}
</script>

<style></style>
