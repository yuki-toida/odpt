<template>
  <table class="table table-sm table-hover">
    <thead>
      <tr>
        <th>#</th>
        <th>TrainSameAs</th>
        <th>進行方向</th>
        <th>出発駅名</th>
        <th>終点駅名</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(obj, index) in data" :key="index">
        <th>{{ index + 1 }}</th>
        <td>
          <nuxt-link :to="trainPath(obj.SameAs)" class="nav-link">
            {{ obj.SameAs }}
          </nuxt-link>
        </td>
        <td>{{ obj.RailDirection.RailDirectionTitleJa }} ({{ obj.RailDirection.RailDirectionTitleEn }})</td>
        <td>{{ obj.FromStation.StationTitleJa }} ({{ obj.FromStation.StationTitleEn }})</td>
        <td>{{ obj.ToStation.StationTitleJa }} ({{ obj.ToStation.StationTitleEn }})</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  async asyncData(context) {
    const { data } = await context.app.$axios.$get("/train/trains")
    return { data: data }
  },
  methods: {
    trainPath(sameAs) {
      return `/train/trains/${sameAs}`
    }
  }
}
</script>
