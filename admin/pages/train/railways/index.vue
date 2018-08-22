<template>
  <table class="table table-sm table-hover">
    <thead>
      <tr>
        <th>#</th>
        <th>路線名</th>
        <th>路線コード</th>
        <th>駅数</th>
        <th>運行会社</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(obj, index) in data" :key="index">
        <th>{{ index + 1 }}</th>
        <td>
          <nuxt-link :to="railwayPath(obj.SameAs)" class="nav-link">
            {{ obj.RailwayTitleJa }} ({{ obj.RailwayTitleEn }})
          </nuxt-link>
        </td>
        <td>{{ obj.LineCode }}</td>
        <td>{{ obj.StationOrders.length }}</td>
        <td>{{ obj.Operator.OperatorTitleJa }} ({{ obj.Operator.OperatorTitleEn }})</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  async asyncData(context) {
    const { data } = await context.app.$axios.$get("/train/railways")
    return { data: data }
  },
  methods: {
    railwayPath(sameAs) {
      return `/train/railways/${sameAs}`
    }
  }
}
</script>
