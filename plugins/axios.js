export default function(context) {
  // rewrite baseURL
  context.$axios.defaults.baseURL = context.env.BASE_URL

  context.$axios.onRequest(config => {
    // add consumerKey
    config.url = `${config.url}?acl:consumerKey=${context.env.CONSUMER_KEY}`
  })

  context.$axios.onError(err => {
    console.log(err.response)
  })
}
