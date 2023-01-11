<template>
  <div>
    <form @submit.prevent="userLogin">
      <div class="input-form">
        <b-field
          label="Email"
          :type="errors.has('email address') ? 'is-danger' : ''"
          :message="errors.has('email address') ? errors.first('email address') : ''"
          class="field"
        >
          <b-input
            name="email address"
            v-validate="'required'"
            v-model="form.email"
            placeholder="Enter Email Address"
          />
        </b-field>

        <b-field
          label="Password"
          :type="errors.has('password') ? 'is-danger' : ''"
          :message="errors.has('password') ? errors.first('password') : ''"
          class="loginpage-label"
        >
          <b-input
            type="password"
            name="password"
            v-validate="'required'"
            password-reveal
            v-model="form.password"
            placeholder="Enter Password"
          />
        </b-field>
      </div>

      <div class="login-button-container">
        <button
          class="button is-fullwidth is-bg-red has-text-white"
          :class="authstatus.class"
          :disabled="authstatus.toggle"
        >Login</button>
      </div>
    </form>
  </div>
</template>

<script>
import { Notification, Notify } from '@/utils/helpers'
import { mapState } from 'vuex'

export default {
  computed: {
    ...mapState(['loading'])
  },
  layout: 'login',
  data() {
    return {
      form: {},
      authstatus: { class: '', toggle: false },
      authmsg: ''
    }
  },
  methods: {
    async userLogin() {
      let result = await this.$validator.validateAll()
      if (result) {
        let param = {
          url: 'auth/login',
          data: this.form
        }
        let req = await this.$store.dispatch('Postdata', param)
        if (req.store) {
          this.$router.push('/page')
          Notification(this, 'Authenticated. Redirecting...', 'is-success', 800)
        }
        if (!req.store) {
          Notification(this, req.errors.message, 'is-danger', 10000)
        }
      }
    }
  },
  watch: {
    loading(newValue, oldValue) {
      if (newValue) {
        this.authstatus = { class: 'is-loading', toggle: true }
      }
      if (newValue) {
        this.authstatus = { class: '', toggle: false }
      }
    }
  }
}
</script>

<style lang="scss" scoped></style>
