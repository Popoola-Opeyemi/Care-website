import CustomContainer from '@/components/widgets/CustomContainer'
import { copy, Notification } from '@/utils/helpers'

export default {
  components: {
    CustomContainer
  },
  async asyncData({ route, store }) {
    let pgname = route.path.split('/')[1]
    let croute = pgname + '/create'
    store.commit('commitpageTitle', pgname)
    return { croute: croute }
  }
}
