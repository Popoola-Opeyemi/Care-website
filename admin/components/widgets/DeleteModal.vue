<template>
  <div class="modal-card" style="width: auto">
    <section class="modal-card-body">
      <section class="modal-card-body is-titleless is-flex">
        <div class="media">
          <div class="media-left">
            <span class="icon has-text-danger is-large"
              ><i class="mdi mdi-alert-circle mdi-48px"></i
            ></span>
          </div>
          <div class="media-content">
            <p>
              Are you sure you want to delete this Item?<br />
              Please note that this action <b>cannot</b> be undone
            </p>
          </div>
        </div>
      </section>
    </section>
    <footer class="modal-card-foot">
      <div class="modal-level-right">
        <button class="button" type="button" @click="$parent.close()">
          Cancel
        </button>
        <button class="button is-danger" @click="deleteItem">
          Delete Item
        </button>
      </div>
    </footer>
  </div>

  <!-- <div>
    <div class="modal-card-body">
      
    </div>
  </div> -->
</template>

<script>
export default {
  methods: {
    deleteItem() {
      this.$emit('deleteItem')
      this.$parent.close()
    }
  }
}
</script>

<style lang="scss" scoped></style>
