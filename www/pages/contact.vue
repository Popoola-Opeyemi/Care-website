<template>
  <div class="contact">
    <div class="pg-title is-blue-bold is-size-3">{{page.attributes.header.title}}</div>
    <div class="contact-content">
      <div class="columns">
        <div class="column">
          <div class="contact-container">
            <div class="head-office is-blue">{{page.attributes.content.List[0].title}}</div>
            <p class="address">{{page.attributes.content.List[0].address}}</p>
            <a
              class="email is-blue-bold is-lowercase"
              :href="`mailto:${page.attributes.content.List[0].email}`"
            >{{page.attributes.content.List[0].email}}</a>
            <p>
              <a class="email is-blue-bold is-lowercase">{{page.attributes.content.List[0].mobile}}</a>
            </p>
          </div>
        </div>
        <div class="column is-7">
          <div class="head-office is-blue">Write Us</div>

          <div class="pg-form">
            <form @submit.prevent="submit">
              <b-field
                label="Name"
                custom-class="is-small"
                :type="errors.has('name') ? 'is-danger': ''"
                :message="errors.first('name')"
              >
                <b-input
                  v-model="form.name"
                  expanded
                  name="name"
                  placeholder="Title"
                  v-validate="'required'"
                />
              </b-field>
              <b-field
                label="Email"
                custom-class="is-small"
                :type="errors.has('email') ? 'is-danger': ''"
                :message="errors.first('email')"
              >
                <b-input
                  v-model="form.email"
                  expanded
                  name="email"
                  placeholder="Email"
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
              <div class="field">
                <b-checkbox
                  v-validate="'required'"
                  v-model="checkbox"
                  name="checkbox"
                  class="is-grey"
                >{{ checkboxmsg }}</b-checkbox>
              </div>
              <button
                :disabled="isLoading"
                :class="isLoading ? 'is-loading':''"
                class="button is-bg-red has-text-white is-size-6 is-rounded"
              >Send Message</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { Notification } from "@/utils/helpers";

import PageMixin from "@/mixins/index";
export default {
  mixins: [PageMixin],

  head() {
    return {
      title: "Contact",
      meta: [
        // hid is used as unique identifier. Do not use `vmid` for it as it will not work
        {
          hid: "title",
          name: "title",
          content:
            "Prudent Domiciliary Care Contact Us. Care Company with coverage in Coverage areas include Bexley and Royal Borough of Greenwich"
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
    return {
      checkbox: false,
      checkboxmsg:
        "By using this form you agree with the storage and handling of your data on this website",
      form: {},
      isLoading: false
    };
  },
  methods: {
    async submit() {
      let result = await this.$validator.validateAll();

      if (result) {
        if (!this.checkbox) {
          Notification(
            this,
            "Please select the checkbox ",
            "is-danger",
            "is-top",
            1000
          );
          return;
        }
        try {
          let param = {
            url: "contacts",
            data: this.form
          };

          this.isLoading = true;
          let req = await this.$store.dispatch("Postdata", param);
          if (req) {
            this.isLoading = false;
            Notification(
              this,
              "Thank you for contacting us we will reach out to you soon.",
              "is-success",
              "is-top",
              6000
            );
            this.form = {};
            this.checkbox = false;
          } else {
            this.isLoading = false;
            Notification(
              this,
              "An error occured. Please try again ",
              "is-danger",
              "is-top",
              3000
            );
          }
        } catch (error) {
          this.isLoading = false;

          Notification(
            this,
            "An error occured. Please try again ",
            "is-danger",
            "is-top",
            3000
          );
        }
      }
    }
  }
};
</script>

<style lang="scss" scoped></style>
