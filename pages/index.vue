<template>
  <div>
    <button type="button" class="btn btn-primary" @click="$store.commit('increment')">
      {{ $store.state.counter }}
    </button>
    {{ ip }}
  </div>
</template>

<script>
export default {
  components: {},
  asyncData({ app, error }) {
    return app.$axios
      .$get("http://icanhazip.com")
      .then(res => {
        return { ip: res }
      })
      .catch(e => {
        error(e)
      })
  }
}
</script>
