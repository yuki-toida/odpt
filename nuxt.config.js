module.exports = {
  head: {
    title: 'odpt',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'Nuxt.js project' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  loading: { color: '#3B8070' },
  build: {
    extend (config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/,
          options : {
            fix : true
          },
        })
      }
    },
    cssSourceMap: false,
  },
  css: [
    '@/assets/scss/app.scss'
  ],
  modules: [
    '@nuxtjs/axios',
    'bootstrap-vue/nuxt',
    ['bootstrap-vue/nuxt', { css: false }],
  ],
  axios: {
    debug: true,
    https: false,
    baseURL: process.env.BASE_URL || 'http://localhost:8080',
  },
  plugins: [
    '~/plugins/axios'
  ],
}
