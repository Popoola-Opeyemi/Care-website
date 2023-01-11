<template>
  <div>
    <section class="custom-dropdown">
      <b-collapse
        class="card"
        :class="displayError ? 'custom-help' : ''"
        :key="index"
        :open="isOpen == index"
        @open="isOpen = index; displayError = false"
        @close="displayError = false"
      >
        <div slot="trigger" slot-scope="props" class="card-header" role="button">
          <p class="card-header-title">Item: {{ index + 1 }}</p>
          <a class="card-header-icon">
            <b-icon :icon="props.open ? 'menu-down' : 'menu-up'"></b-icon>
          </a>
        </div>
        <div class="card-content">
          <div class v-if="!EditMode">
            <div class="level">
              <div class="level-right level-item">
                <button
                  class="button is-danger is-small has-text-right is-outlined print-invoice-button"
                  @click="delSection(form)"
                >Delete</button>
              </div>
            </div>
          </div>
          <div class="columns">
            <div class="column is-9">
              <b-field
                label="Title"
                custom-class="is-small"
                :type="errors.has('title') ? 'is-danger': ''"
                :message="errors.first('title')"
              >
                <b-input
                  expanded
                  v-model="form.title"
                  name="title"
                  :disabled="EditMode"
                  placeholder="Title"
                  v-validate="'required'"
                />
              </b-field>
            </div>
            <div class="column">
              <b-field
                label="Page Link"
                custom-class="is-small"
                :type="errors.has('Link') ? 'is-danger': ''"
                :message="errors.first('Link')"
              >
                <b-select
                  v-model="form.link"
                  expanded
                  name="Page Link"
                  :disabled="EditMode"
                  placeholder="Link To"
                  v-validate="'required'"
                >
                  <option
                    :value="option.route"
                    v-for="(option, index) in pagesList"
                    :key="index"
                  >{{ option.name }}</option>
                </b-select>
              </b-field>
            </div>
          </div>
          <image-upload
            v-model="form.image"
            label="Background Image"
            @image="isImageValid($event)"
            :readonly="EditMode"
          />
        </div>
      </b-collapse>
      <span
        class="displayDanger help is-danger"
        v-if="displayError"
      >There is pending data thats needs to save.</span>
    </section>
  </div>
</template>

<script>
import ImageUpload from '@/components/widgets/ImageUpload.vue'

export default {
  components: { ImageUpload },
  props: {
    total: {
      type: Number,
      default: () => 0
    },
    index: {
      type: Number,
      default: () => 0
    },
    itemIdx: {
      type: Number,
      default: () => 0
    },
    EditMode: {
      type: Boolean,
      default: () => false
    },
    form: {
      type: Object,
      default: () => {}
    },
    pagesList: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      isOpen: false,
      displayError: false,
      isImage: false
    }
  },
  methods: {
    delSection(obj) {
      this.$emit('deleteSection', obj)
    },
    // delSection(obj) {
    //   let catalog = this.form
    //   let delobj = _.find(catalog, obj)
    //   catalog.splice(catalog.indexOf(delobj), 1)
    // },
    isImageValid(validity) {
      this.isImage = validity
    },
    async isValid() {
      let result = await this.$validator.validateAll()
      if (!result) {
        return
      }
      this.$emit('input', copy(this.form))
      return true
    },
    emitForm() {
      this.$emit(this.form)
    },
    async validate() {
      this.displayError = false

      if (!this.isImage) {
        this.displayError = true
        return false
      }
      let result = await this.$validator.validateAll()
      if (!result) {
        this.displayError = true
        return false
      }
      return true
    }
  }
}
</script>

<style lang="scss" scoped>
</style>