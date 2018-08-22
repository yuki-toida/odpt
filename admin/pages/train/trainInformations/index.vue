<template>
  <table class="table table-sm table-hover">
    <thead>
      <tr>
        <th>#</th>
        <th>TrainInformationSameAs</th>
        <th>発生路線</th>
        <th>発生場所起点</th>
        <th>発生場所終点</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(obj, index) in data" :key="index">
        <th>{{ index + 1 }}</th>
        <td>
          <nuxt-link :to="trainInformationPath(obj.SameAs)" class="nav-link">
            {{ obj.SameAs }}
          </nuxt-link>
        </td>
        <td>{{ obj.Railway.RailwayTitleJa }} ({{ obj.Railway.RailwayTitleEn }})</td>
        <td>{{ obj.StationFrom.StationTitleJa }} ({{ obj.StationFrom.StationTitleEn }})</td>
        <td>{{ obj.StationTo.StationTitleJa }} ({{ obj.StationTo.StationTitleEn }})</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  async asyncData(context) {
    const { data } = await context.app.$axios.$get("/train/trainInformations")
    return { data: data }
  },
  methods: {
    trainInformationPath(sameAs) {
      return `/train/trainInformations/${sameAs}`
    }
  }
}
</script>
