export default {
  mode: "universal",
  /*
   ** Headers of the page
   */
  head: {
    title: process.env.npm_package_name || "",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: process.env.npm_package_description || ""
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },
  /*
   ** Customize the progress-bar color
   */
  loading: {
    color: "#46518c",
    height: "5px"
  },
  /*
   ** Global CSS
   */
  css: [
    "@/assets/scss/style.scss",
    "@node_modules/@mdi/font/css/materialdesignicons.css"
  ],
  /*
   ** Plugins to load before mounting the App
   */

  plugins: ["~plugins/veevalidate"],

  buildModules: [],
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://buefy.github.io/#/documentation
    "nuxt-buefy",
    // Doc: https://axios.nuxtjs.org/usage
    "@nuxtjs/axios"
  ],
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {
    proxy: true,
    credentials: true
  },

  proxy: {
    "/api": "http://127.0.0.1:5000",
    "/public": "http://127.0.0.1:5000",
    "/static": "http://127.0.0.1:5000"
    // "/api": "https://pbgcare.co.uk/",
    // "/public": "https://pbgcare.co.uk/",
    // "/static": "https://pbgcare.co.uk/"
  },
  /*
   ** Build configuration
   */
  build: {
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {}
  }
};
