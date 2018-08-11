<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>カレンダー</th>
          <td>{{ data.Calendar }}</td>
        </tr>
        <tr>
          <th>進行方向</th>
          <td>{{ data.RailDirection }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <caption>オブジェクト</caption>
      <thead>
        <tr>
          <th>#</th>
          <th>オブジェクト</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Objects" :key="index">
          <th>{{ index }}</th>
          <td>{{ obj }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  async asyncData(context) {
    const url = `/train/stationTimetables/${context.params.sameAs}`
    const { data, error } = await context.app.$axios.$get(url)
    return { data: data, error: error }
  },
  mounted: function() {
    if (this.error) {
      this.$toast.error(this.error)
    }
  }
}
</script>
