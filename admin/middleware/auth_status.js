export default async function({ store, redirect }) {
  let param = {
    url: 'auth/status'
  }
  let status = await store.dispatch('Get', param) // data now exists
  if (!status.done) {
    window.location = '/'
  }
}
