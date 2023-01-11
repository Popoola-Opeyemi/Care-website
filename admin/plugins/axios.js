export default function({ $axios, redirect, store }) {
  $axios.onRequest(config => {
    store.commit('commitrequestError', { status: false, error: null })
    store.commit('commitLoading', true)
  })

  $axios.onResponse(response => {
    store.commit('commitLoading', false)
    store.commit('commitEdit', true)
  })

  $axios.onResponseError(err => {
    store.commit('commitrequestError', { status: true, error: err })
    store.commit('commitLoading', false)
  })
}
