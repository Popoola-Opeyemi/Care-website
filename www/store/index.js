import _ from "lodash";

export const state = () => ({
  servicesBanner: {},
  settings: {
    mobile: "01322686765",
    address: "Old Street, London EC1V 9BD",
    mobile2: "+447440713708",
    facebook: "https://web.facebook.com/PrudentDomCare",
    linkedin: "https://www.linkedin.com/company/prudent-domiciliary-care",
    instagram: "https://www.linkedin.com/company/prudent-domiciliary-care"
  },
  homePage: {
    image: {
      data: "",
      name: "",
      size: ""
    }
  },
  PageData: {
    image: {
      data: "",
      name: "",
      size: ""
    }
  },
  pagesList: [],
  servicesList: []
});

export const mutations = {
  commitPages(state, data) {
    state.pagesList = data;
  },
  commitServicesList(state, data) {
    state.servicesList = data;
  },
  commitPageData(state, data) {
    state.PageData = data;
  },
  commitHomePage(state, data) {
    state.homePage = data;
  },
  commitSettings(state, data) {
    state.settings = data;
  },
  commitservicesBanner(state, data) {
    state.servicesBanner = data;
  }
};

export const getters = {
  getPageId(state) {
    return function(route) {
      let pg = state.pagesList;
      if (_.isEmpty(pg)) {
        return [];
      }
      route = route.toLowerCase();

      let id = _.find(pg, s => s.route.toLowerCase() == route);
      return id;
    };
  },
  getServiceId(state) {
    return function(route) {
      let pg = state.servicesList;
      if (_.isEmpty(pg)) {
        return [];
      }
      route = route.toLowerCase();

      let id = _.find(pg, s => s.url.toLowerCase() == route);
      return id;
    };
  }
};

export const actions = {
  async Get({ commit }, param) {
    try {
      let req = await this.$axios.$get(`api/${param.url}`);
      return req;
    } catch (error) {
      throw error;
    }
  },

  async Postdata({ commit }, param) {
    try {
      let req = await this.$axios.$post(`api/${param.url}`, param.data);
      return req.store;
    } catch (error) {
      console.log("an error occured", error);
      throw error;
    }
  }
};
