<template>
  <div>
    <quill-editor
      :value="value"
      :disabled="disabled"
      @change="handleInput($event)"
      placeholder
      :options="editorOption"
    />
  </div>
</template>

<script>
export default {
  props: {
    disabled: {
      type: Boolean,
      default: false
    },

    value: {
      type: String,
      default: ''
    },

    name: {
      type: String,
      default: ''
    }
  },

  data() {
    return {
      editorOption: {
        modules: {
          toolbar: [
            ['bold', 'italic', 'underline', 'strike'],
            ['clean'],
            [{ color: ['#46518c', '#fff', '#000', '#d93636'] }],
            [{ align: [] }],
            [{ header: [1, 2, 3, 4, 5, 6, false] }],
            [{ list: 'ordered' }, { list: 'bullet' }]
          ]
        }
      },
      data: '',
      hasValue: true
    }
  },

  computed: {
    // hasValue(){
    //   return _.isEmpty(this.value);
    // }
  },

  methods: {
    handleInput(e) {
      this.$emit('input', e.html)
      this.hasValue = !_.isEmpty(e.html)
    }
  }
}
</script>

<style></style>
