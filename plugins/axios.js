export default function({ $axios }) {
  $axios.onRequest(config => {
    console.log(config.url)
  })
  $axios.onError(error => {
    console.log(error)
  })
}
