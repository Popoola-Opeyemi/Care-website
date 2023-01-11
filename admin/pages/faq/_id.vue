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
          label="Content"
          custom-class="is-small"
          :type="errors.has('content') ? 'is-danger' : ''"
          :message="errors.first('content')"
        >
          <custom-editor
            name="content"
            :disabled="EditMode"
            v-validate="'required'"
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
