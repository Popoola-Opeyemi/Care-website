<template>
  <div>
    <b-field
      label="Subtitle"
      custom-class="is-small"
      :type="errors.has('subtitle') ? 'is-danger': ''"
      :message="errors.first('subtitle')"
    >
      <b-input
        expanded
        v-model="pageData.attributes.header.content.subtitle"
        name="subtitle"
        :disabled="EditMode"
        placeholder="Subtitle"
        v-validate="'required'"
      />
    </b-field>
    <b-field
      label="Subtitle"
      custom-class="is-small"
      :type="errors.has('subtitle') ? 'is-danger': ''"
      :message="errors.first('subtitle')"
    >
      <b-input
        expanded
        v-model="pageData.attributes.header.content.subtitle"
        name="subtitle"
        :disabled="EditMode"
        placeholder="Subtitle"
        v-validate="'required'"
      />
    </b-field>
  </div>
</template>

<script>
export default {}
</script>

<style lang="scss" scoped>
</style>