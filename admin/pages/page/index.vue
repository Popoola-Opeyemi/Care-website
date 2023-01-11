<template>
  <div class="pages">
    <div class="custom-box">
      <div class="box-items" v-for="page in pagesList.items" :key="page.id">
        <div class="level item">
          <nuxt-link :to="'/page/' + page.route">
            <div class="left">{{ page.name }}</div>
          </nuxt-link>
          <div class="right">
            <nuxt-link :to="'/page/' + page.route" class="mdi mdi-pencil"></nuxt-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  async asyncData({ store }) {
    store.commit('commitpageTitle', 'Pages')
    let param = {
      url: 'pages?page=1&perPage=20'
    }
    let req = await store.dispatch('Get', param)
    store.commit('commitPages', req.store.data.items)
    // console.log(req.store)
    return { pagesList: req.store.data }
  }
}
</script>

<style lang="scss" scoped></style>
