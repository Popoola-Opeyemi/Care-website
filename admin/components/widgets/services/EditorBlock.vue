<template>
  <div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-Amount">Edit block</p>
      </header>
      <section class="modal-card-body">
        <b-field
          label="Content"
          custom-class="is-small"
          :type="errors.has('content') ? 'is-danger' : ''"
          :message="errors.first('content')"
        >
          <custom-editor
            :class="errors.has('title') ? 'custom-help' : ''"
            name="content"
            :disabled="EditMode"
            v-validate="'required'"
            v-model="content"
          ></custom-editor>
        </b-field>
      </section>
    </div>
  </div>
</template>

<script>
import CustomEditor from '@/components/widgets/CustomEditor'

export default {
  components: {
    CustomEditor
  },
  props: {
    EditMode: {
      type: Boolean,
      default: () => false
    },
    form: {
      type: Object,
      default: () => {}
    }
  }
}
</script>

<style lang="scss" scoped></style>
