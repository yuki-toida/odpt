export default function(context) {
  context.$axios.onRequest(config => {
    console.log(config.url)
  })

  context.$axios.onError(err => {
    console.log(err.response)
  })
}
