<template>
  <div class="faq">
    <div class="pg-title is-blue-bold is-size-3">{{page.attributes.header.title}}</div>
    <div class="container">
      <!-- <p class="is-blue-bold is-size-5">General</p> -->
      <div class="faq-list">
        <b-collapse
          class="card faq-card"
          :open="isOpen"
          v-for="faq in faqsList.data.items"
          :key="faq.id"
        >
          <div slot="trigger" class="card-header" role="button" aria-controls="contentIdForA11y3">
            <p class="card-header-title is-blue-bold is-size-6">{{faq.name}}</p>
            <a class="card-header-icon">
              <b-icon :icon="isOpen ? 'menu-down' : 'menu-up'"></b-icon>
            </a>
          </div>
          <div class="card-content">
            <div class="content is-blue is-glight is-size-5" v-html="faq.content"></div>
          </div>
        </b-collapse>
      </div>
    </div>
  </div>
</template>

<script>
import PageMixin from "@/mixins/index";
export default {
  mixins: [PageMixin],
  head() {
    return {
      title: "Faqs",
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "title",
          name: "title",
          content:
            "Prudent Domiciliary Care Frequently asked Questions. Care Company with coverage in Coverage areas include Bexley and Royal Borough of Greenwich"
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
  // async asyncData({ store }) {
  //   store.commit("setbanner", "faq.jpg");
  // },
  data() {
    return {
      isOpen: false
    };
  }
};
</script>

<style lang="scss" scoped></style>
