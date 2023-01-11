<template>
  <div>
    <custom-container
      @save="save()"
      @edit="toggleEdit"
      @cancel="cancel"
      :hide="false"
      :hide1="false"
      :hideCreate="true"
    >
      <custom-box :title="'View Contact'">
        <b-field
          label="Name"
          custom-class="is-small"
          :type="errors.has('name') ? 'is-danger' : ''"
          :message="errors.first('name')"
        >
          <b-input
            expanded
            v-model="form.name"
            name="name"
            placeholder="Name"
            :disabled="EditMode"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Email"
          custom-class="is-small"
          :type="errors.has('email') ? 'is-danger' : ''"
          :message="errors.first('email')"
        >
          <b-input
            expanded
            v-model="form.email"
            name="email"
            placeholder="Email"
            :disabled="EditMode"
            v-validate="'required'"
          />
        </b-field>
        <b-field
          label="Date"
          custom-class="is-small"
          :type="errors.has('date') ? 'is-danger' : ''"
          :message="errors.first('date')"
        >
          <b-input
            expanded
            v-model="form.created_at"
            name="date"
            placeholder="Date"
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
            v-model="form.message"
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
