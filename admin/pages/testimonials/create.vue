<template>
  <div>
    <custom-container :hide="false" :hide1="true">
      <custom-box :title="'Create Testimonial'">
        <b-field
          label="Author"
          custom-class="is-small"
          :type="errors.has('author') ? 'is-danger' : ''"
          :message="errors.first('author')"
        >
          <b-input
            expanded
            v-model="form.author"
            name="author"
            placeholder="Author"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Content"
          custom-class="is-small"
          :type="errors.has('content') ? 'is-danger' : ''"
          :message="errors.first('content')"
        >
          <b-input
            v-model="form.content"
            expanded
            type="textarea"
            maxlength="230"
            name="content"
            placeholder="Content"
            v-validate="'required'"
          />
        </b-field>
      </custom-box>
      <div class="columns">
        <div class="column is-5">
          <button class="button has-text-centered is-danger is-fullwidth">cancel</button>
        </div>
        <div class="column">
          <button
            @click="submit"
            class="button has-text-centered is-bg-blue has-text-white is-fullwidth"
          >Submit</button>
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
  }
}
</script>

<style lang="scss" scoped></style>
