<template>
  <div>
    <custom-container
      @save="submit()"
      @edit="toggleEdit"
      @cancel="cancel"
      :hide="true"
      :hide1="false"
      :hideCreate="true"
    >
      <custom-box :title="'Edit'">
        <b-field
          label="Title"
          custom-class="is-small"
          :type="errors.has('title') ? 'is-danger' : ''"
          :message="errors.first('title')"
        >
          <b-input
            expanded
            v-model="form.name"
            :disabled="EditMode"
            name="title"
            placeholder="Title"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Slug Name"
          custom-class="is-small"
          :type="errors.has('Slug Name') ? 'is-danger' : ''"
          :message="errors.first('Slug Name')"
        >
          <b-input
            v-model="slugUrl"
            :disabled="EditMode"
            expanded
            name="Slug Name"
            placeholder="Slug Name"
            v-validate="'required'"
          />
        </b-field>
        <image-upload
          label="Header Image"
          :readonly="EditMode"
          v-model="form.image"
          @image="isImageValid($event)"
        />
        <b-field
          label="Content"
          custom-class="is-small"
          :type="errors.has('content') ? 'is-danger' : ''"
          :message="errors.first('content')"
        >
          <custom-editor
            name="content"
            v-validate="'required'"
            :disabled="EditMode"
            v-model="form.content"
          ></custom-editor>
        </b-field>
      </custom-box>
    </custom-container>
  </div>
</template>

<script>
import CustomDropdown from '@/components/widgets/Dropdown'
import CustomContainer from '@/components/widgets/CustomContainer'
import CustomEditor from '@/components/widgets/CustomEditor'
import { Slugify, unSlugify } from '@/utils/helpers'
import { copy, Notification } from '@/utils/helpers'

import PagesMix from '@/mixins/page.id'

export default {
  mixins: [PagesMix],
  components: {
    CustomDropdown,
    CustomContainer,
    CustomEditor
  },
  data() {
    return {
      slugUrl: null,
      isImage: false,

      form: {}
    }
  },
  methods: {
    isImageValid(validity) {
      this.isImage = validity
    },
    async submit() {
      if (!this.isImage) {
        Notification(this, 'Please Upload an Image.', 'is-danger')
        return
      }
      await this.save()
    }
  },
  watch: {
    slugUrl(newval, oldval) {
      this.form.url = Slugify(newval)
    },
    'form.name': function(val, oldval) {
      this.slugUrl = Slugify(val)
    },
    slugUrl(newval, oldval) {
      let slug = Slugify(newval)
      this.slugUrl = slug
      this.form.url = slug
    }
  },
  mounted() {
    this.slugUrl = unSlugify(this.form.url)
  }
}
</script>

<style lang="scss" scoped></style>
