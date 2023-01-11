<template>
  <div class="main-header is-hidden-mobile">
    <nav class="level">
      <div class="level-left">
        <div class="round is-bg-red">
          <a class="mdi mdi-phone has-text-white" href></a>
        </div>
        <div class="level-item">
          <div class="mobile-numbers">
            <a href class>{{ settings.mobile }}</a>
          </div>
        </div>
        <div class="level-item">
          <div class="mobile-numbers">
            <a href class>{{ settings.mobile2 }}</a>
          </div>
        </div>
      </div>
      <div class="level-right">
        <div class="round is-bg-red">
          <a class="mdi mdi-map-marker has-text-white" href></a>
        </div>
        <div class="address">
          <div class="level-item">{{ settings.address }}</div>
        </div>
        <div class="socials">
          <div class="level-item">
            <!-- <div class="round is-bg-grey">
              <a class="mdi mdi-twitter has-text-white" href></a>
            </div>-->
            <div class="round is-bg-grey">
              <a class="mdi mdi-linkedin has-text-white" target="_blank" :href="settings.linkedin"></a>
            </div>

            <div class="round is-bg-grey">
              <a
                class="mdi mdi-instagram has-text-white"
                target="_blank"
                :href="settings.instagram"
              ></a>
            </div>
            <div class="round is-bg-grey">
              <a class="mdi mdi-facebook has-text-white" target="_blank" :href="settings.facebook"></a>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  computed: {
    ...mapState(["settings"])
  },
  data() {
    return {
      Header: {
        contactNumbers: ["01322686765", "+447440713708"],
        address: "130 Old Street, London EC1V 9BD",
        socialMedia: {
          linkedin: "https://www.linkedin.com/company/prudent-domiciliary-care",
          twitter: "",
          facebook: "https://web.facebook.com/PrudentDomCare/",
          instagram:
            "https://instagram.com/prudentdomiciliarycare?igshid=7acpi27gh9fg",
          whatsapp: ""
        }
      }
    };
  }
};
</script>

<style lang="scss" scoped></style>
