<template>
  <div class="modal-card">
    <div class="modal-card-head">
      <h1 class="modal-info-heading">
        Edit Admin Details
      </h1>
    </div>

    <div class="modal-card-body">
      <div class="columns">
        <div class="column">
          <b-field
            label="First Name"
            :type="errors.has('name') ? 'is-danger' : ''"
            :message="errors.first('name')"
          >
            <b-input
              v-model="admin.first_name"
              placeholder="First Name"
              v-validate="'required'"
              name="name"
            />
          </b-field>

          <b-field
            label="Email"
            :type="errors.has('email') ? 'is-danger' : ''"
            :message="errors.first('email')"
          >
            <b-input
              v-model="admin.email"
              placeholder="Email"
              type="email"
              v-validate="'required'"
              name="email"
            />
          </b-field>
        </div>

        <div class="column">
          <b-field
            label="Surname"
            :type="errors.has('surname') ? 'is-danger' : ''"
            :message="errors.first('surname')"
          >
            <b-input
              v-model="admin.surname"
              placeholder="Surame"
              v-validate="'required'"
              name="surname"
            />
          </b-field>

          <b-field
            label="Password"
            :type="errors.has('password') ? 'is-danger' : ''"
            :message="errors.collect('password')"
          >
            <b-input
              name="password"
              type="password"
              v-validate="'required'"
              placeholder="Enter Password"
              password-reveal
              v-model="admin.password"
            />
          </b-field>
        </div>
      </div>
    </div>

    <div class="modal-card-foot">
      <button class="button is-success is-fullwidth" @click="send">
        Save
      </button>

      <button class="button is-danger is-fullwidth" @click="cancel">
        Cancel
      </button>
    </div>
  </div>
</template>

<script>
import { copy } from '@/utils/helpers'

export default {
  props: {
    value: {
      type: Object,
      default() {
        return {}
      }
    }
  },

  computed: {
    admin() {
      let data = copy(this.value)
      return data
    }
  },

  methods: {
    cancel() {
      this.$parent.close()
    },

    async send() {
      let result = await this.$validator.validateAll()
      if (result) {
        let data = copy(this.admin)
        this.$emit('send', data)
        this.$parent.close()
      }
    }
  }
}
</script>

<style></style>
