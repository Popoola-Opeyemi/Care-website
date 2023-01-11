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

      <custom-box :title="'Section: Contacts'">
        <!-- <template v-slot:button>
          <button
            class="button is-blue"
            :disabled="EditMode"
            @click="addBlock"
            v-if="pageData.attributes.content.List.length < 2"
          >add new block</button>
        </template>-->
        <custom-dropdown
          ref="custom_dropdown"
          v-for="(form, idx) in pageData.attributes.content.List"
          :key="idx"
          :index="idx"
          @deleteSection="delSection($event)"
          :form="form"
          :EditMode="EditMode"
        ></custom-dropdown>
      </custom-box>
    </custom-container>
  </div>
</template>

<script>
import PagesMix from '@/mixins/admin_pages'
import CustomEditor from '@/components/widgets/CustomEditor'
import CustomDropdown from '@/components/widgets/contacts/Dropdown'

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
    delSection(obj) {
      let catalog = this.pageData.attributes.content.List
      let delobj = _.find(catalog, obj)
      catalog.splice(catalog.indexOf(delobj), 1)
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
