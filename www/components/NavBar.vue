<template>
  <div>
    <b-navbar>
      <template slot="brand">
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img src="/img/Dom-Care-logo.png" alt="Logo for Prudent Business Group" />
        </b-navbar-item>
      </template>
      <template slot="end">
        <b-navbar-item
          tag="nuxt-link"
          class="navigation-item is-red"
          v-for="nav in Navigation"
          :key="nav.id"
          :to="nav.link"
        >{{ nav.name }}</b-navbar-item>
      </template>
      <template slot="end">
        <b-navbar-item tag="div">
          <nuxt-link to="join">
            <b-button rounded class="is-bg-red has-text-white">Join Us</b-button>
          </nuxt-link>
        </b-navbar-item>
      </template>
    </b-navbar>
  </div>
</template>

<script>
export default {
  data() {
    return {
      Navigation: [
        { name: "Abouts Us", link: "/about" },
        { name: "Services", link: "/services" },
        { name: "Process", link: "/process" },
        { name: "Faqs", link: "/faqs" },
        { name: "Contact Us", link: "/contact" }
      ]
    };
  }
};
</script>

<style lang="scss" scoped></style>
