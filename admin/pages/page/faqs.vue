<template>
  <div>
    <custom-container @save="save" @edit="toggleEdit" @cancel="cancel" :hideCreate="true">
      <custom-box :title="'Section: Page Header'">
        <image-upload
          label="Header Image"
          :readonly="EditMode"
          v-model="pageData.attributes.header.image"
          @image="isImageValid($event)"
        />
        <b-field
          label="Title"
          custom-class="is-small"
          :type="errors.has('title') ? 'is-danger': ''"
          :message="errors.first('title')"
        >
          <b-input
            expanded
            v-model="pageData.attributes.header.title"
            name="title"
            :disabled="EditMode"
            placeholder="Title"
            v-validate="'required'"
          />
        </b-field>
      </custom-box>
    </custom-container>
  </div>
</template>

<script>
import PagesMix from '@/mixins/admin_pages'
import CustomEditor from '@/components/widgets/CustomEditor'
import CustomDropdown from '@/components/widgets/process/Dropdown'

import ServiceEditor from '@/components/widgets/services/EditorBlock'
import { copy, clone, Notification } from '@/utils/helpers'

export default {
  mixins: [PagesMix],
  components: {
    CustomEditor,
    ServiceEditor,
    CustomDropdown
  },
  methods: {
    cancel() {
      this.$router.push('/page')
    },
    async save() {
      let result = await this.$validator.validateAll()
      if (result) {
        let form = copy(this.pageData)
        let param = {
          url: `pages/${this.uuid}`,
          data: form
        }
        let req = await this.$store.dispatch('Postdata', param)
        if (req.done) {
          Notification(this, 'Changes saved successfully', 'is-success')
          this.$router.push('/page')
        } else {
          Notification(this, 'an error occured', 'is-danger')
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped></style>
