<template>
  <div>
    <custom-container
      @save="save()"
      @edit="toggleEdit"
      @cancel="cancel"
      :hide="true"
      :hide1="false"
      :hideCreate="true"
    >
      <custom-box :title="'Edit Testimonial'">
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
            :disabled="EditMode"
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
            :disabled="EditMode"
            maxlength="230"
            name="content"
            placeholder="Content"
            v-validate="'required'"
          />
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
      form: {
        isImage: true
      }
    }
  }
}
</script>

<style lang="scss" scoped></style>
