export default {
  mode: 'spa',
  /*
   ** Headers of the page
   */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || ''
      }
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }]
  },
  /*
   ** Customize the progress-bar color
   */
  loading: { color: '#fff' },
  /*
   ** Global CSS
   */
  css: [
    '@node_modules/@mdi/font/css/materialdesignicons.css',
    '@/assets/scss/style.scss'
  ],
  /*
   ** Plugins to load before mounting the App
   */
  plugins: [
    '~plugins/veevalidate',
    '~plugins/lodash',
    '~plugins/vue-quill',
    '~/plugins/axios',
    '~/plugins/moment'
  ],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [],
  /*
   ** Nuxt.js modules
   */
  modules: ['nuxt-buefy', '@nuxtjs/axios', '@nuxtjs/dotenv'],
  axios: {
    proxy: true,
    credentials: true
  },
  proxy: {
    '/api': 'http://127.0.0.1:5000',
    '/public': 'http://127.0.0.1:5000',
    '/static': 'http://127.0.0.1:5000'
    // '/api': 'http://admin.pbgcare.co.uk/',
    // '/public': 'http://admin.pbgcare.co.uk/',
    // '/static': 'http://admin.pbgcare.co.uk/'
  },

  build: {
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {}
  }
}
