<template>
  <div class="join-us">
    <div class="pg-title is-blue-bold">{{page.attributes.content.title}}</div>
    <div class="join-content" v-html="page.attributes.content.content"></div>

    <!-- <div class="join-us-email is-blue">
      All Applications should be sent to
      <a
        class="email is-blue-bold is-lowercase"
        href="mailto:info@pbgcare.co.uk"
      >info@pbgcare.co.uk</a>
    </div>-->
  </div>
</template>

<script>
import PageMixin from "@/mixins/index";
export default {
  mixins: [PageMixin],

  head() {
    return {
      title: "Join Us",
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "title",
          name: "title",
          content:
            "Prudent Domiciliary Join Us. Care Company with coverage in Coverage areas include Bexley and Royal Borough of Greenwich"
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

<style lang="scss" scoped>
</style>