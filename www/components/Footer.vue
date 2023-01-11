<template>
  <div class="footer-content">
    <div class="columns">
      <div class="column">
        <div class="footer-content__footer-title">EXPLORE PBG DOMICILIARY CARE</div>
        <div class="footer-content__footer-link width30" v-for="a in Navigation" :key="a.id">
          <span class="nav-title">></span>
          <nuxt-link :to="a.link" class="has-text-white">
            {{
            a.name
            }}
          </nuxt-link>
        </div>
      </div>
      <div class="column">
        <div class="footer-content__footer-title">Our Services</div>
        <div v-for="service in servicesList" :key="service.id" class="footer-content__footer-link">
          <span class="nav-title">></span>
          <nuxt-link :to="'/services/'+service.url" class="has-text-white">
            {{
            service.name
            }}
          </nuxt-link>
          <div class="footer-underline"></div>
        </div>
      </div>
      <div class="column">
        <div class="footer-content__footer-title">Office Address</div>
        <div class="footer-content__footer-link">
          <div class="office-title">PRUDENT DOMICILIARY CARE LIMITED</div>
          <div class="office-reg">
            <div class="office-title">Registered Office</div>
            <div class="office-address">{{settings.address}}</div>
          </div>
          <a
            class="office-email has-text-white is-lowercase"
            :href="`mailto:${settings.email}`"
          >{{settings.email}}</a>
        </div>
        <div class="box contact-box">
          <div class="contact-title is-blue-bold">We are here to help</div>
          <div class="contact-subtitle is-red is-gbold">Request a callback</div>
          <form @submit.prevent="submit">
            <b-field
              label="Fullname"
              custom-class="is-small"
              :type="errors.has('Fullname') ? 'is-danger': ''"
              :message="errors.first('Fullname')"
            >
              <b-input
                v-model="form.name"
                expanded
                name="Fullname"
                placeholder="Fullname"
                v-validate="'required'"
              />
            </b-field>
            <b-field
              label="Mobile Number"
              custom-class="is-small"
              :type="errors.has('Mobile Number') ? 'is-danger': ''"
              :message="errors.first('Mobile Number')"
            >
              <b-input
                v-model="form.mobile_number"
                expanded
                name="Mobile Number"
                placeholder="Mobile Number"
                v-validate="'required'"
              />
            </b-field>
            <b-field
              label="Email"
              custom-class="is-small"
              :type="errors.has('Email') ? 'is-danger': ''"
              :message="errors.first('Email')"
            >
              <b-input
                v-model="form.email"
                expanded
                name="Email"
                placeholder="Email"
                v-validate="'required'"
              />
            </b-field>
            <b-field
              label="Message"
              custom-class="is-small"
              :type="errors.has('Message') ? 'is-danger': ''"
              :message="errors.first('Message')"
            >
              <b-input
                v-model="form.message"
                expanded
                type="textarea"
                name="Message"
                placeholder="Message"
                v-validate="'required'"
              />
            </b-field>
            <button
              :class="isLoading ? 'is-loading':''"
              class="button is-bg-blue has-text-white"
            >Submit</button>
            <!-- <b-button class="button is-bg-red has-text-white">Submit</b-button> -->
          </form>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column has-text-centered">
        <img class src="/img/banner/white-logo.png" alt />
        <p>Copyright &#9400; 2019. Prudent Domiciliary Care Ltd. All Rights Reserved.</p>
      </div>
    </div>
  </div>
</template>
<script>
import { Notification } from "@/utils/helpers";
import { mapState } from "vuex";

export default {
  data() {
    return {
      form: {},
      isLoading: false,
      Navigation: [
        { name: "Abouts Us", link: "about" },
        { name: "Services", link: "services" },
        { name: "Process", link: "process" },
        { name: "FaQs", link: "faqs" },
        { name: "Contact Us", link: "contact" },
        { name: "Join Us", link: "join-us" }
      ]
    };
  },
  computed: {
    ...mapState(["servicesList", "settings"])
  },
  methods: {
    async submit() {
      let result = await this.$validator.validateAll();
      if (result) {
        try {
          let param = {
            url: "contacts",
            data: this.form
          };
          let req = await this.$store.dispatch("Postdata", param);
          this.isLoading = true;

          if (req) {
            this.isLoading = false;
            Notification(
              this,
              "Thank you for filling the form",
              "is-success",
              "is-top",
              1000
            );
            this.form = {};
          } else {
            this.isLoading = false;
            Notification(
              this,
              "An error occured. Please try again ",
              "is-danger",
              "is-top",
              1000
            );
          }
        } catch (error) {
          this.isLoading = false;

          Notification(
            this,
            "An error occured. Please try again ",
            "is-danger",
            "is-top",
            1000
          );
        }
      }
    }
  }
};
</script>

<style lang="scss" scoped></style>
