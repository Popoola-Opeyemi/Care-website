<template>
  <div class="dashboard-sidebar">
    <aside>
      <div class="top-header">Prudent Care</div>
      <div class="navigation">
        <ul class="navigation-list">
          <li
            v-for="list in NavigationList"
            :key="list.id"
            class="navigation-list-li"
          >
            <nuxt-link :to="list.url" class="navigation-list-link">
              <i :class="list.icon" class="mdicon mdi">
                <span class="navigation-list-name">
                  {{ list.name }}
                </span>
              </i>
            </nuxt-link>
          </li>
        </ul>
      </div>
      <!-- <ul class="navigation-list">
        <li class="nav" v-for="nav in NavigationList" :key="nav.id">
          <nuxt-link :to="nav.url">
            <i :class="['mdi', nav.icon]" />
            {{ nav.name }}
          </nuxt-link>
        </li>
      </ul> -->
    </aside>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  computed: {
    ...mapState(['NavigationList'])
  }
}
</script>

<style lang="scss" scoped></style>
