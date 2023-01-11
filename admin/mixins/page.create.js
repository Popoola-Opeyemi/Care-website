import { copy, Notification } from '@/utils/helpers'
import CustomContainer from '@/components/widgets/CustomContainer'
import CustomBox from '@/components/widgets/CustomBox'
import ImageUpload from '@/components/widgets/ImageUpload.vue'
import CustomEditor from '@/components/widgets/CustomEditor'

import { mapState } from 'vuex'

export default {
  components: {
    CustomContainer,
    CustomBox,
    ImageUpload,
    CustomEditor
  },
  async asyncData({ store, route }) {
    let name = route.path.split('/')[1]
    store.commit('commitpageTitle', `Create ${name}`)

    return { route: name }
  },
  data() {
    return {
      isImage: false,
      form: {}
    }
  },
  computed: {
    ...mapState(['EditMode', 'pagesList'])
  },
  methods: {
    async submit() {
      let result = await this.$validator.validateAll()
      if (result) {
        let param = {
          url: `${this.route}`,
          data: this.form
        }
        let req = await this.$store.dispatch('Postdata', param)
        if (req.done) {
          Notification(this, 'Successfully Created Post.', 'is-success')
          setTimeout(() => this.$router.push(`/${this.route}`), 1000)
        }
      }
    },
    isImageValid(validity) {
      this.isImage = validity
    },
    toggleEdit() {
      this.$store.commit('commitEdit', !this.EditMode)
    }
  }
}
