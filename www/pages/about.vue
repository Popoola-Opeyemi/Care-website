<template>
  <div class="about">
    <div class="pg-title is-blue-bold">About Us</div>
    <div class="pg-content" v-html="page.attributes.content.content"></div>
  </div>
</template>

<script>
import PageMixin from "@/mixins/index";
export default {
  mixins: [PageMixin],
  head() {
    return {
      title: "About",
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "title",
          name: "title",
          content:
            "Prudent Domiciliary Care About Us. Care Company with coverage in Coverage areas include Bexley and Royal Borough of Greenwich"
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
  },

  data() {
    return {};
  }
};
</script>

<style lang="scss" scoped></style>
