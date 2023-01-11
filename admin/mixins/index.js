// import { , Notify } from '@/utils/helpers'
import { copy, Notification } from '@/utils/helpers'

export default {
  // computed: {
  //   ...mapState('requestError')
  // }
  created() {
    this.$store.subscribe((mutation, state) => {
      if (mutation.type.includes('commitrequestError')) {
        // display error in global error output
        let payload = mutation.payload
        if (payload.status) {
          Notification(this, 'An Error Occured', 'is-danger', 1000)
        }
      }
    })
  }
}
