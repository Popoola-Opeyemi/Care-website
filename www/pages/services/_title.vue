<template>
  <div>
    <div class="services-children">
      <div class="pg-title is-blue-bold">{{ pg.name }}</div>
      <div class="service-content" v-html="pg.content"></div>
      <nuxt-link to="/contact">
        <b-button
          class="service-button"
          type="is-bg-red has-text-white is-size-6"
          outlined
          rounded
        >GET IN TOUCH WITH US</b-button>
      </nuxt-link>
    </div>
  </div>
</template>

<script>
import PageMixin from "@/mixins/index.id";
export default {
  mixins: [PageMixin],
  layout: "services",
  head() {
    return {
      title: this.pg.name,
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "title",
          name: "title",
          content: `Prudent Domiciliary ${this.pg.name}. Care Company with coverage in Coverage areas include Bexley and Royal Borough of Greenwich`
        },
        {
          hid: "description",
          name: "description",
          content:
            "Bexely Foremost Care Company, Uk. We Providing the Quality home care you need Coverage areas include Bexely and Royal Borough of Greenwich. Contact us today!"
        },
        {
          hid: "keywords",
          name: "keywords",
          content: `${this.pg.name}, Bexely, Care, BexleyCare, Bexely Dementia Care, Home Care United Kingdom, Elder Care `
        }
      ]
    };
  }
};
</script>

<style lang="scss" scoped></style>
