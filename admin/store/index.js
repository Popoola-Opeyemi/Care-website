import { da } from 'date-fns/locale'

export const state = () => ({
  loading: false,
  requestError: { status: false, error: null },
  pageTitle: '',
  EditMode: true,
  pagesList: [],
  NavigationList: [
    // { icon: 'mdi-view-dashboard', name: 'Dashboard', url: '/dashboard' },
    { icon: 'mdi-library-books', name: 'Pages', url: '/page' },
    { icon: 'mdi-briefcase', name: 'Services', url: '/services' },
    { icon: 'mdi-account-group', name: 'Testimonials', url: '/testimonials' },
    { icon: 'mdi-frequently-asked-questions', name: 'Faq', url: '/faq' },
    { icon: 'mdi-contacts', name: 'Contacts', url: '/contacts' },
    { icon: 'mdi-settings', name: 'Settings', url: '/settings' }
  ]
})

export const mutations = {
  commitLoading(state, data) {
    state.loading = data
  },
  commitrequestError(state, data) {
    state.requestError = data
  },
  commitpageTitle(state, data) {
    state.pageTitle = data
  },
  commitPages(state, data) {
    state.pagesList = data
  },
  commitEdit(state, data) {
    state.EditMode = data
  }
}

export const getters = {
  getPageId(state) {
    return function(route) {
      let pg = state.pagesList
      if (_.isEmpty(pg)) {
        return []
      }
      route = route.toLowerCase()

      let id = _.find(pg, s => s.route.toLowerCase() == route)
      return id
    }
  }
}
export const actions = {
  async Get({ commit }, param) {
    try {
      let req = await this.$axios.$get(`api/${param.url}`)
      return req
    } catch (error) {
      throw error
    }
  },

  async Delete({ commit }, param) {
    try {
      let req = await this.$axios.$delete(`api/${param.url}`)
      return req
    } catch (error) {
      throw error
    }
  },

  async Postdata({ commit }, param) {
    try {
      let req = await this.$axios.$post(`api/${param.url}`, param.data)

      return req
    } catch (error) {
      console.log('an error occured', error)
      throw error
    }
  }
}
