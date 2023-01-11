<template>
  <div>
    <custom-container
      :croute="backlink"
      @save="save"
      :hideCreate="true"
      @edit="toggleEdit"
      @cancel="cancel"
    >
      <!-- {{pageData}} -->
      <custom-box :title="'Section: Page Header'">
        <b-field
          label="Title"
          custom-class="is-small"
          :type="errors.has('title') ? 'is-danger': ''"
          :message="errors.first('title')"
        >
          <b-input
            expanded
            v-model="pageData.attributes.header.content.title"
            name="title"
            :disabled="EditMode"
            placeholder="Title"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Subtitle"
          custom-class="is-small"
          :type="errors.has('subtitle') ? 'is-danger': ''"
          :message="errors.first('subtitle')"
        >
          <b-input
            expanded
            v-model="pageData.attributes.header.content.subtitle"
            name="subtitle"
            :disabled="EditMode"
            placeholder="Subtitle"
            v-validate="'required'"
          />
        </b-field>
        <div class="columns">
          <div class="column">
            <image-upload
              label="Header Background Image"
              :readonly="EditMode"
              v-model="pageData.attributes.header.image"
              @image="isImageValid($event)"
            />
          </div>
          <div class="column">
            <image-upload
              label="Header Logo "
              :readonly="EditMode"
              v-model="pageData.attributes.header.content.image"
              @image="isImageValid($event)"
            />
          </div>
        </div>
      </custom-box>
      <custom-box :title="'Section: Testimonials'">
        <template v-slot:button>
          <button
            class="button is-blue"
            :disabled="EditMode"
            @click="addBlock"
            v-if="pageData.attributes.content.List.length < 3"
          >add new block</button>
        </template>
        <!-- <b-field
          label="Title"
          custom-class="is-small"
          :type="errors.has('title') ? 'is-danger': ''"
          :message="errors.first('title')"
        >
          <b-input
            expanded
            v-model="pageData.attributes.header.content.title"
            name="title"
            :disabled="EditMode"
            placeholder="Title"
            v-validate="'required'"
          />
        </b-field>-->
        <custom-dropdown
          ref="custom_dropdown"
          v-for="(form, idx) in pageData.attributes.content.List"
          :key="idx"
          :index="idx"
          @deleteSection="delSection($event)"
          :form="form"
          :pagesList="pagesList"
          :EditMode="EditMode"
        ></custom-dropdown>
      </custom-box>
    </custom-container>
  </div>
</template>

<script>
import PagesMix from '@/mixins/admin_pages'
import CustomDropdown from '@/components/widgets/Dropdown'
import { copy, clone, Notification } from '@/utils/helpers'

export default {
  mixins: [PagesMix],
  components: {
    CustomDropdown
  },
  data() {
    return {
      form: {}
    }
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
    },
    addBlock() {
      if (!this.EditMode) {
        let form = this.pageData.attributes.content.List
        let block = { title: '', link: '', image: {} }
        form.push(block)
      }
    }
  }
  // mounted() {
  //   let original = copy(this.pageData)
  // }
}
</script>

<style lang="scss" scoped>
</style>