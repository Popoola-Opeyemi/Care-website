import CustomContainer from '@/components/widgets/CustomContainer'
import DeleteModal from '@/components/widgets/DeleteModal'

import { copy, Notification } from '@/utils/helpers'

export default {
  components: {
    CustomContainer,
    DeleteModal
  },
  data() {
    return {
      perPage: 15,
      current: 1
    }
  },
  async asyncData({ route, store }) {
    let pgname = route.path.split('/')[1]
    let croute = pgname + '/create'
    store.commit('commitpageTitle', pgname)
    let param = {
      url: `${pgname}?page=1&perPage=15`
    }
    let page = await store.dispatch('Get', param)
    let pgstore = copy(page.store.data)
    return {
      pgname: pgname,
      pageList: pgstore,
      croute: croute,
      total: pgstore.count
    }
  },
  methods: {
    deleteAttrib(id) {
      let that = this
      this.$buefy.modal.open({
        parent: this,
        component: DeleteModal,
        hasModalCard: true,
        customClass: '',
        trapFocus: true,
        events: {
          deleteItem(event) {
            that.deleteItem(id)
          }
        }
      })
    },
    async deleteItem(id) {
      let param = {
        url: `${this.pgname}/${id}`
      }
      let page = await this.$store.dispatch('Delete', param)
      if (page.done) {
        this.spliceItem(id)
      }
      Notification(this, 'Item Deleted.', 'is-success')
    },
    spliceItem(id) {
      let item = this.pageList.items
      let itm = _.find(item, { id })
      item.splice(item.indexOf(itm), 1)
    },
    minify(content) {
      if (content.length > 20) {
        let minified = content.substring(0, 70)
        return minified + '....'
      }
    },
    async setPage(val) {
      let param = { page: val, offset: this.perPage }
      let params = {
        url: `${this.pgname}?page=${param.page}&perPage=${param.offset}`
      }
      let page = await this.$store.dispatch('Get', params)
      this.pageList = page.store.data
    }
  }
}
