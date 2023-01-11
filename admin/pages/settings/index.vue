<template>
  <div>
    <custom-container
      :hide="true"
      @save="save()"
      @edit="toggleEdit"
      @cancel="cancel"
      :hide1="false"
      :hideCreate="true"
    >
      <custom-box :title="'Social Media'">
        <b-field
          label="Facebook Url"
          custom-class="is-small"
          :type="errors.has('facebook url') ? 'is-danger' : ''"
          :message="errors.first('facebook url')"
        >
          <b-input
            expanded
            v-model="form.attributes.facebook"
            name="facebook url"
            :disabled="EditMode"
            placeholder="Facebook Url"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Linkedin Url"
          custom-class="is-small"
          :type="errors.has('linkedin url') ? 'is-danger' : ''"
          :message="errors.first('linkedin url')"
        >
          <b-input
            expanded
            v-model="form.attributes.linkedin"
            name="linkedin url"
            :disabled="EditMode"
            placeholder="Linkedin Url"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Instagram Url"
          custom-class="is-small"
          :type="errors.has('instagram url') ? 'is-danger' : ''"
          :message="errors.first('instagram url')"
        >
          <b-input
            expanded
            v-model="form.attributes.instagram"
            name="instagram url"
            :disabled="EditMode"
            placeholder="Instagram Url"
            v-validate="'required'"
          />
        </b-field>
      </custom-box>

      <custom-box :title="'Contact Information'">
        <b-field
          label="Contact Address"
          custom-class="is-small"
          :type="errors.has('Contact Address') ? 'is-danger' : ''"
          :message="errors.first('Contact Address')"
        >
          <b-input
            expanded
            v-model="form.attributes.address"
            name="Contact Address"
            :disabled="EditMode"
            placeholder="Contact Address"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Email Address"
          custom-class="is-small"
          :type="errors.has('email address') ? 'is-danger' : ''"
          :message="errors.first('email address')"
        >
          <b-input
            expanded
            v-model="form.attributes.email"
            name="email address"
            :disabled="EditMode"
            placeholder="Email Address"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Mobile Number"
          custom-class="is-small"
          :type="errors.has('mobile number') ? 'is-danger' : ''"
          :message="errors.first('mobile number')"
        >
          <b-input
            expanded
            v-model="form.attributes.mobile"
            name="mobile number"
            :disabled="EditMode"
            placeholder="Mobile Number"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Mobile Number 2"
          custom-class="is-small"
          :type="errors.has('mobile number 2') ? 'is-danger' : ''"
          :message="errors.first('mobile number 2')"
        >
          <b-input
            expanded
            v-model="form.attributes.mobile2"
            name="mobile number 2"
            :disabled="EditMode"
            placeholder="Mobile Number 2"
            v-validate="'required'"
          />
        </b-field>
      </custom-box>
    </custom-container>
  </div>
</template>

<script>
import { mapState } from 'vuex'

import CustomContainer from '@/components/widgets/CustomContainer'
import CustomBox from '@/components/widgets/CustomBox'
import { copy, Notification } from '@/utils/helpers'

import PagesMix from '@/mixins/page.index'

export default {
  async asyncData({ route, store }) {
    // boa8j0drmtiktn3tkovg
    let pgname = route.path.split('/')[1]
    let croute = pgname + '/create'
    store.commit('commitpageTitle', pgname)

    let param = {
      url: `${pgname}/1`
    }
    let page = await store.dispatch('Get', param)

    if (page.done) {
      let pgstore = copy(page.store)
      return { form: pgstore, apiRoute: `${pgname}/1` }
    }

    return { form: { attributes: {} }, apiRoute: pgname }
  },
  components: {
    CustomContainer,
    CustomBox
  },
  computed: {
    ...mapState(['EditMode'])
  },
  methods: {
    cancel() {
      this.$router.push(`${this.apiRoute}`)
    },

    toggleEdit() {
      this.$store.commit('commitEdit', !this.EditMode)
    },
    async save() {
      let result = await this.$validator.validateAll()
      if (result) {
        let param = {
          url: this.apiRoute,
          data: { attributes: this.form.attributes }
        }
        let req = await this.$store.dispatch('Postdata', param)
        if (req.done) {
          Notification(this, 'Operation completed successfully.', 'is-success')
          // setTimeout(() => this.$router.push(`/${this.apiRout}`), 1000)
        } else {
          Notification(this, 'An error occured.', 'is-danger')
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
</style>