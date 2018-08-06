<template>
  <div>
    <dl v-for="(value, key) in error.response" :key="key">
      <dt>{{ key }}</dt>
      <dd>{{ value }}</dd>
    </dl>
  </div>
</template>
<script>
export default {
  props: ["error"],
  created: function() {
    console.log(this.error.response)
  }
}
</script>
