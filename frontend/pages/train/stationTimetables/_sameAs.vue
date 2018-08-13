<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>StationTimetableID</th>
          <td>{{ data.ID }}</td>
        </tr>
        <tr>
          <th>StationTimetableSameAs</th>
          <td>{{ data.SameAs }}</td>
        </tr>
        <tr>
          <th>カレンダー</th>
          <td>{{ data.Calendar.Title }}</td>
        </tr>
        <tr>
          <th>進行方向</th>
          <td>{{ data.RailDirection.Title }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>出発時刻</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Objects" :key="index">
          <th>{{ index + 1 }}</th>
          <td>
            <nuxt-link :to="stationTimetableObjectPath(obj.ID)" class="nav-link">
              {{ obj.DepartureTime }}
            </nuxt-link>
          </td>
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
  },
  methods: {
    stationTimetableObjectPath(id) {
      return `/train/stationTimetableObjects/${id}`
    }
  }
}
</script>
