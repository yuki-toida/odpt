<template>
  <div>
    <h1>エラーが発生しました</h1>
    <nuxt-link to="/">ホーム</nuxt-link>
  </div>
</template>
<script>
export default {
  props: {
    error: {
      type: Object,
      default: null
    }
  },
  created: function() {
    console.log(this.error)
  }
}
</script>
