<template>
  <div class="about">
    <custom-container @save="save" @edit="toggleEdit" @cancel="cancel" :hideCreate="true">
      <custom-box :title="'Section: Page Header'">
        <image-upload
          label="Header Image"
          :readonly="EditMode"
          v-model="pageData.attributes.header.image"
          @image="isImageValid($event)"
        />
      </custom-box>
      <custom-box :title="'Section: Middle section'">
        <template v-slot:button>
          <!-- <b-tag class="button is-blue" :disabled="!EditMode" @click="addBlock">add new block</b-tag> -->
          <button class="button is-blue" :disabled="EditMode" @click="addBlock">add new block</button>
        </template>
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
    delSection(obj) {
      let catalog = this.pageData.attributes.content.List
      let delobj = _.find(catalog, obj)
      catalog.splice(catalog.indexOf(delobj), 1)
    },
    addBlock() {
      if (!this.EditMode) {
        let form = this.pageData.attributes.content.List
        let block = { title: '', image: {}, content: '' }
        form.push(block)
      }
    },
    cancel() {
      this.$router.push('/page')
    },
    editService() {
      let that = this
      this.$buefy.modal.open({
        parent: this,
        component: ServiceEditor,
        hasModalCard: true,
        trapFocus: true,
        events: {
          input(val) {}
        }
      })
    },
    async save() {
      let ref = this.$refs.custom_dropdown
      let validated = true

      for (let a = 0; a < ref.length; a++) {
        let x = await ref[a].validate()
        validated = x && validated
      }

      let result = await this.$validator.validateAll()
      if (!this.isImage) {
        Notification(this, 'Please upload image.', 'is-danger')
        return
      }

      if (result) {
        if (validated) {
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
}
</script>

<style lang="scss" scoped></style>
