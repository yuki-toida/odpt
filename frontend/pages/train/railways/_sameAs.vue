<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>RailwayID</th>
          <td>{{ data.ID }}</td>
        </tr>
        <tr>
          <th>RailwaySameAs</th>
          <td>{{ data.SameAs }}</td>
        </tr>
        <tr>
          <th>路線名(Ja)</th>
          <td>{{ data.RailwayTitleJa }}</td>
        </tr>
        <tr>
          <th>路線名(En)</th>
          <td>{{ data.RailwayTitleEn }}</td>
        </tr>
        <tr>
          <th>路線コード</th>
          <td>{{ data.LineCode }}</td>
        </tr>
        <tr>
          <th>駅数</th>
          <td>{{ data.StationOrders.length }}</td>
        </tr>
        <tr>
          <th>運行会社(Ja)</th>
          <td>{{ data.Operator.OperatorTitleJa }}</td>
        </tr>
        <tr>
          <th>運行会社(En)</th>
          <td>{{ data.Operator.OperatorTitleEn }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>駅順序</th>
          <th>駅名</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="obj in data.StationOrders" :key="obj.Index">
          <th>{{ obj.Index }}</th>
          <td>
            <nuxt-link :to="path(obj.StationSameAs)" class="nav-link">
              {{ obj.StationTitleJa }} ({{ obj.StationTitleEn }})
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
    const url = `/train/railways/${context.params.sameAs}`
    const { data, error } = await context.app.$axios.$get(url)
    if (error) {
      context.error({ message: error })
      return
    }
    return { data: data }
  },
  methods: {
    path(sameAs) {
      return `/train/stations/${sameAs}`
    }
  }
}
</script>
