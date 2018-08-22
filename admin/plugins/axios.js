export default function(context) {
  context.$axios.onError(err => {
    console.log(err)
  })
}
