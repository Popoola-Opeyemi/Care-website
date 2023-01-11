<template>
  <div>
    <div class="custom-box">
      <div class="level">
        <div class="level-item level-right">
          <nuxt-link :to="croute">
            <b-button
              @click="create"
              v-if="display()"
              class="create has-text-white is-bg-blue"
              icon-left="book"
            >Create</b-button>
            <!-- <button 
            class="button is-bg-blue has-text-white" 
            v-if="display()" 
            @click="create">Create</button>-->
          </nuxt-link>
          <b-button
            @click="toggleEdit"
            v-if="EditMode && hide"
            class="has-text-white is-bg-blue"
            icon-left="pencil"
          >Edit</b-button>
          <!-- <button
            class="button is-bg-blue has-text-white"
            v-if="EditMode && hide"
            @click="toggleEdit"
          >Edit</button>-->

          <button
            class="button cancel is-bg-red has-text-white"
            v-if="!EditMode && hide"
            @click="cancel"
          >Cancel</button>
          <b-button
            v-if="!EditMode && hide"
            @click="save"
            class="save has-text-white is-bg-blue"
            icon-left="content-save"
          >Save</b-button>
          <!-- <button
            class="button save has-text-white is-bg-blue"
            v-if="!EditMode && hide"
            @click="save"
          >Save</button>-->
        </div>
      </div>
      <slot></slot>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  computed: {
    ...mapState(['EditMode'])
  },
  props: {
    edit: {
      type: Boolean,
      default: () => true
    },
    croute: {
      type: String,
      default: () => 'page'
    },
    hide: {
      type: Boolean,
      default: () => true
    },
    hide1: {
      type: Boolean,
      default: () => false
    },
    hideCreate: {
      type: Boolean,
      default: () => false
    }
  },
  methods: {
    display() {
      if (this.hideCreate) {
        return false
      }
      if (!this.hide1) {
        return true
      } else {
        return this.hide
      }
    },
    toggleEdit() {
      this.$emit('edit')
    },
    save() {
      this.$emit('save')
    },
    cancel() {
      this.$emit('cancel')
    },
    create() {
      this.$emit('create')
    }
  }
}
</script>

<style lang="scss" scoped></style>
