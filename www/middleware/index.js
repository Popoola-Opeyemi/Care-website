export default async function({ store, redirect }) {
  let pgList = store.state.pagesList;
  let name = route.path.split("/")[2];
  let pgback = "page";
  store.commit("commitpageTitle", `pages: ${name}`);
  if (_.isEmpty(pgList)) {
    let param = {
      url: "pages?page=1&perPage=20"
    };
    let req = await store.dispatch("Get", param);
    const pg = req.store.data.items;
    store.commit("commitPages", pg);
  }
}
