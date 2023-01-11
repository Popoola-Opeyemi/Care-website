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
  async asyncData({ store, route, params }) {
    let back = route
    let apiRoute = route.path.split('/')[1]
    let pgname = params.id
    store.commit('commitpageTitle', `View ${apiRoute}`)
    let param = {
      url: `${apiRoute}/${pgname}`
    }
    let req = await store.dispatch('Get', param)
    return { form: req.store, apiRoute: apiRoute, pgname: pgname }
  },
  data() {
    return {}
  },
  computed: {
    ...mapState(['EditMode', 'pagesList'])
  },
  methods: {
    async save() {
      let result = await this.$validator.validateAll()
      if (result) {
        let param = {
          url: `${this.apiRoute}/${this.form.uuid}`,
          data: this.form
        }
        let req = await this.$store.dispatch('Postdata', param)
        if (req.done) {
          Notification(this, 'Successfully Edited Post.', 'is-success')
          setTimeout(() => this.$router.push(`/${this.apiRoute}`), 1000)
        }
      }
    },
    toggleEdit() {
      this.$store.commit('commitEdit', !this.EditMode)
    },
    cancel() {
      this.$router.push(`/${this.apiRoute}`)
    },

    toggleEdit() {
      this.$store.commit('commitEdit', !this.EditMode)
    }
  }
}
