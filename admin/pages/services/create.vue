<template>
  <div>
    <custom-container :hide="false" :hide1="true">
      <custom-box :title="'Create Service'">
        <b-field
          label="Title"
          custom-class="is-small"
          :type="errors.has('title') ? 'is-danger' : ''"
          :message="errors.first('title')"
        >
          <b-input
            expanded
            v-model="form.name"
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
            expanded
            name="Slug Name"
            placeholder="Slug Name"
            v-validate="'required'"
          />
        </b-field>
        <image-upload
          label="Header Image"
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
            v-model="form.content"
          ></custom-editor>
        </b-field>
      </custom-box>
      <div class="columns">
        <div class="column is-5">
          <button class="button has-text-centered is-danger is-fullwidth">
            cancel
          </button>
        </div>
        <div class="column">
          <button
            @click="submit"
            class="button has-text-centered is-bg-blue has-text-white is-fullwidth"
          >
            Submit
          </button>
        </div>
      </div>
    </custom-container>
  </div>
</template>

<script>
import CustomDropdown from '@/components/widgets/Dropdown'
import PagesMix from '@/mixins/page.create'
import { Slugify } from '@/utils/helpers'

export default {
  mixins: [PagesMix],
  components: {
    CustomDropdown
  },
  data() {
    return {
      slugUrl: ''
    }
  },
  watch: {
    'form.name': function(val, oldval) {
      this.slugUrl = Slugify(val)
    },
    slugUrl(newval, oldval) {
      let slug = Slugify(newval)
      this.slugUrl = slug
      this.form.url = slug
    }
  }
}
</script>

<style lang="scss" scoped></style>
