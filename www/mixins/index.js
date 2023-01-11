// import { Notification, Notify } from '@/utils/helpers'
import { copy } from "@/utils/helpers";
import { mapState } from "vuex";
import _ from "lodash";

export default {
  async asyncData({ store, route }) {
    let Result = [];
    let asyncReq = [
      { name: "testimonials", url: "testimonials?page=1&perPage=3" },
      { name: "faq", url: "faq?page=1&perPage=50" },
      { name: "settings", url: "settings/1" },
      { name: "serviceList", url: "services/all" }
    ];
    let name = route.path.split("/")[1];

    if (route.name == "index") {
      name = "home";
    }

    let pgList = store.state.pagesList;

    if (_.isEmpty(pgList)) {
      let param = {
        url: "pages?page=1&perPage=20"
      };
      let req = await store.dispatch("Get", param);
      const pg = req.store.data.items;
      store.commit("commitPages", pg);
    }

    let pg = store.getters["getPageId"](name);
    let param = {
      url: `pages/${pg.uuid}`
    };

    let req = await store.dispatch("Get", param);
    for (let a = 0; a < asyncReq.length; a++) {
      Result[asyncReq[a].name] = await store.dispatch("Get", {
        url: asyncReq[a].url
      });
      Result[asyncReq[a].name] = Result[asyncReq[a].name].store;
    }

    if (req.done) {
      store.commit("commitSettings", Result.settings.attributes);
      store.commit("commitServicesList", Result.serviceList.data);

      if (name == "home") {
        store.commit("commitHomePage", req.store.attributes.header);
      }
      store.commit("commitPageData", req.store.attributes.header);
      return {
        page: req.store,
        pgname: name,
        testimonialList: Result.testimonials,
        faqsList: Result.faq
      };
    }
  },
  computed: {
    ...mapState(["PageData"])
  }
};
