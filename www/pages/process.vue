<template>
  <div class="process">
    <div class="pg-title is-blue-bold">Our Care Process</div>
    <div class="pg-process-content">
      <div class="columns" v-for="process in page.attributes.content.List" :key="process.id">
        <div class="column is-2">
          <img class="process-image" :src="process.image.data" alt />
        </div>
        <div class="column">
          <div class="process-title is-blue is-gmedium has-text-weight-semibold">{{process.title}}</div>
          <div class="proceess-content" v-html="process.content"></div>
        </div>
      </div>
      <nuxt-link to="/contact">
        <b-button type="is-bg-red has-text-white is-size-5" outlined rounded>GET IN TOUCH WITH US</b-button>
      </nuxt-link>
    </div>
  </div>
</template>

<script>
import PageMixin from "@/mixins/index";
export default {
  mixins: [PageMixin],

  head() {
    return {
      title: "Process",
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "title",
          name: "title",
          content:
            "Prudent Domiciliary Care Our Process. Care Company with coverage in Coverage areas include Bexley and Royal Borough of Greenwich"
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
          content:
            "Bexely, Care, BexleyCare, Bexely Dementia Care, Home Care United Kingdom, Elder Care "
        }
      ]
    };
  }
};
</script>

<style lang="scss" scoped></style>
