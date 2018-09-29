<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>StationTimetableObjectID</th>
          <td>{{ data.ID }}</td>
        </tr>
        <tr>
          <th>出発時刻</th>
          <td>{{ data.DepartureTime }}</td>
        </tr>
        <tr>
          <th>列車番号</th>
          <td>{{ data.TrainNumber }}</td>
        </tr>
        <tr>
          <th>列車種別</th>
          <td>{{ data.TrainType.TrainTypeTitleJa }} ({{ data.TrainType.TrainTypeTitleEn }})</td>
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
          <th>編成名称</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.TrainNames" :key="index">
          <th>{{ index + 1 }}</th>
          <td>{{ obj.TrainNameJa }} ({{ obj.TrainNameEn }})</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>経由路線</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.ViaRailways" :key="index">
          <th>{{ index + 1 }}</th>
          <td>{{ obj.RailwaySameAs }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  async asyncData(context) {
    const url = `/admin/stationTimetableObjects/${context.params.id}`
    const { data, error } = await context.app.$axios.$get(url)
    if (error) {
      context.error({ message: error })
      return
    }
    return { data: data }
  }
}
</script>
