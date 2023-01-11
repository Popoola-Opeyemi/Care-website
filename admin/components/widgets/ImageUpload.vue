<template>
  <b-field
    :label="label"
    custom-class="is-small"
    :type="errors.has('image') ? 'is-danger': ''"
    :message="errors.first('image')"
  >
    <b-field class="__imageupload__">
      <b-input
        disabled
        v-model="imgName"
        expanded
        name="image"
        placeholder="Image Name"
        v-validate="'required'"
      />

      <p class="fileInput">
        <input ref="fileInput" type="file" @change="getFile" />
      </p>
      <p class="control" v-if="!readonly">
        <button
          type="button"
          class="button upload mdi mdi-upload has-background-grey-light"
          @click="onUpload"
        >&nbsp;Upload</button>
      </p>

      <p class="control">
        <button
          type="button"
          :disabled="validImage"
          class="button preview mdi mdi-eye has-background-grey-light"
          @click="previewImage"
        >&nbsp;Preview</button>
      </p>
    </b-field>
  </b-field>
</template>

<script>
import ImagePreview from '@/components/widgets/ImagePreview.vue'
import { mapState } from 'vuex'
import { clone } from '@/utils/helpers'

export default {
  props: {
    value: {
      type: Object,
      default: () => {
        return {
          data: '',
          name: '',
          size: 0
        }
      }
    },

    label: {
      type: String,
      default: ''
    },

    readonly: {
      type: Boolean,
      default: false
    }
  },

  data() {
    return {
      isImage: false,
      imgName: this.value.name,
      validImage: false,
      image: {
        data: '',
        name: '',
        size: 0
      }
    }
  },

  mounted() {
    if (!_.isEmpty(this.value)) {
      this.isImageFile(this.value.name)
    } else {
      this.isImageFile(this.image.name)
    }
  },

  watch: {
    value: {
      handler: function(nv) {
        this.imgName = nv.name
        this.isImage = false
      },
      immediate: true
    }
  },

  methods: {
    onUpload() {
      this.$refs.fileInput.click()
    },

    isImageFile(val) {
      if (_.isEmpty(val)) {
        this.validImage = true
        val = []
      } else {
        this.validImage = false
      }
      let imgArray = ['jpg', 'png', 'jpeg', 'svg']
      let valid = false
      for (var i = 0; i < imgArray.length; i++) {
        if (val.indexOf(imgArray[i]) != -1) {
          valid = true
          break
        }
      }
      this.isImage = valid
      this.$emit('image', this.isImage)
      return valid
    },

    getFile(evt) {
      let el = evt.target
      if (el.files && el.files[0]) {
        if (!this.isImageFile(el.files[0].name)) {
          let name = el.files[0].name
          this.$buefy.snackbar.open({
            message: `The selected file (${name}) is not an image file (only jpg, jpeg, png files allowed)`,
            position: 'is-top',
            type: 'is-danger'
          })
          return
        }

        let reader = new FileReader()
        reader.onload = e => {
          if (!this.previewAfterUpload) {
            this.image.data = e.target.result
            this.image.name = el.files[0].name
            this.image.size = e.total

            let data = clone(this.image)

            this.$emit('input', data)
          }
        }

        reader.readAsDataURL(el.files[0])
      }
    },

    previewImage() {
      this.$buefy.modal.open({
        parent: this,
        component: ImagePreview,
        props: {
          image: this.value.data
        }
      })
    }
  }
}
</script>

<style>
</style>