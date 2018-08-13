<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>TrainID</th>
          <td>{{ data.ID }}</td>
        </tr>
        <tr>
          <th>TrainSameAs</th>
          <td>{{ data.SameAs }}</td>
        </tr>
        <tr>
          <th>車両数</th>
          <td>{{ data.CarComposition }}</td>
        </tr>
        <tr>
          <th>遅延時間(秒)</th>
          <td>{{ data.Delay }}</td>
        </tr>
        <tr>
          <th>列車順序</th>
          <td>{{ data.Index }}</td>
        </tr>
        <tr>
          <th>進行方向</th>
          <td>{{ data.RailDirection.Title }}</td>
        </tr>
        <tr>
          <th>列車番号</th>
          <td>{{ data.TrainNumber }}</td>
        </tr>
        <tr>
          <th>RailwaySameAs</th>
          <td>{{ data.RailwaySameAs }}</td>
        </tr>
        <tr>
          <th>TrainTypeSameAs</th>
          <td>{{ data.TrainTypeSameAs }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>終着駅</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.DestinationStations" :key="index">
          <th>{{ index + 1 }}</th>
          <td>{{ obj.StationSameAs }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>始発駅</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.OriginStations" :key="index">
          <th>{{ index + 1 }}</th>
          <td>{{ obj.StationSameAs }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  async asyncData(context) {
    const url = `/train/trains/${context.params.sameAs}`
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
