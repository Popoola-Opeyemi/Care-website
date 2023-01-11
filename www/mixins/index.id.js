import { mapState } from "vuex";
import _ from "lodash";

export default {
  async asyncData({ store, params }) {
    let title = params.title;
    let Result = [];

    let asyncReq = [
      { name: "services", url: `services/${title}` },
      { name: "settings", url: "settings/1" },

      { name: "serviceList", url: "services/all" }
    ];

    for (let a = 0; a < asyncReq.length; a++) {
      Result[asyncReq[a].name] = await store.dispatch("Get", {
        url: asyncReq[a].url
      });
      Result[asyncReq[a].name] = Result[asyncReq[a].name].store;
    }
    store.commit("commitServicesList", Result.serviceList.data);
    store.commit("commitSettings", Result.settings.attributes);
    store.commit("commitservicesBanner", Result.services);

    return { pg: Result.services };
  }
};
