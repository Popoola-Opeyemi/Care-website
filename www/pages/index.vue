<template>
  <div class="homepage">
    <nav class="level">
      <div class="level-left"></div>
      <!-- Left side -->
      <div class="level-right">
        <div class="level-item">
          <img src="/img/banner/Dementia-Care-logo.jpg" alt />
        </div>
      </div>
    </nav>
    <div class="columns is-8 is-variable">
      <div
        class="column"
        v-for="pg in page.attributes.content.List"
        :key="pg.id"
      >
        <div
          class="home-img-box"
          :style="{ 'background-image': `url(${pg.image.data})` }"
        >
          <div class="box-content">
            <p class="box-title has-text-centered has-text-white">
              {{ pg.title }}
            </p>
            <nuxt-link :to="`${pg.link}`">
              <b-button type="is-white is-size-5" outlined rounded
                >Learn More</b-button
              >
            </nuxt-link>
          </div>
        </div>
      </div>
    </div>
    <section class="testimonials">
      <div class="testimonial-title is-red">Clients Testimonials</div>
      <div class="testimonial-container">
        <div class="columns">
          <div
            class="column"
            v-for="testimony in testimonialList.data.items"
            :key="testimony.id"
          >
            <div class="box testimonial-box">
              <div class="testimonial-text is-capitalized">
                {{ testimony.content }}
              </div>
              <div class="testimonial-author is-red">
                {{ testimony.author }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import PageMixin from "@/mixins/index";

export default {
  layout: "home",
  mixins: [PageMixin],

  head() {
    return {
      title: this.pgname,
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "title",
          name: "title",
          content:
            "Prudent Domiciliary Care - London, UK Based Intl. Care Company"
        },
        {
          hid: "description",
          name: "description",
          content:
            "Bexley Foremost Care Company, Uk. We Providing the Quality home care you need Coverage areas include Bexley and Royal Borough of Greenwich. Contact us today!"
        },
        {
          hid: "keywords",
          name: "keywords",
          content:
            "Bexley, Care, BexleyCare, Bexley Dementia Care, Home Care United Kingdom, Elder Care "
        }
      ]
    };
  },
  data() {
    return {};
  }
};
</script>
