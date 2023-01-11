import { copy, Notification } from '@/utils/helpers'
import CustomContainer from '@/components/widgets/CustomContainer'
import CustomBox from '@/components/widgets/CustomBox'
import ImageUpload from '@/components/widgets/ImageUpload.vue'
import { mapState } from 'vuex'

export default {
  components: {
    CustomContainer,
    CustomBox,
    ImageUpload
  },
  async asyncData({ store, route }) {
    let name = route.path.split('/')[2]
    let pgback = 'page'
    store.commit('commitpageTitle', `pages: ${name}`)
    let pgList = store.state.pagesList
    if (_.isEmpty(pgList)) {
      let param = {
        url: 'pages?page=1&perPage=20'
      }
      let req = await store.dispatch('Get', param)
      const pg = req.store.data.items
      store.commit('commitPages', pg)
    }
    let pg = store.getters['getPageId'](name)
    let param = {
      url: `pages/${pg.uuid}`
    }
    let page = await store.dispatch('Get', param)
    console.log('uuid', pg.uuid)
    return { pageData: copy(page.store), uuid: pg.uuid, backlink: pgback }
  },
  data() {
    return {
      isImage: false
    }
  },
  computed: {
    ...mapState(['EditMode', 'pagesList'])
  },
  methods: {
    isImageValid(validity) {
      this.isImage = validity
    },
    toggleEdit() {
      this.$store.commit('commitEdit', !this.EditMode)
    }
  }
}
