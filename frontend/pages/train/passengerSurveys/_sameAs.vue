<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>PassengerSurveyID</th>
          <td>{{ data.ID }}</td>
        </tr>
        <tr>
          <th>PassengerSurveySameAs</th>
          <td>{{ data.SameAs }}</td>
        </tr>
        <tr>
          <th>運行会社</th>
          <td>{{ data.Operator.OperatorTitleJa }} ({{ data.Operator.OperatorTitleEn }})</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>調査年度</th>
          <th>一日平均乗降人員数</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Objects" :key="index">
          <th>{{ index + 1 }}</th>
          <td>{{ obj.SurveyYear }}</td>
          <td>{{ obj.PassengerJourneys }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>路線名</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Railways" :key="index">
          <th>{{ index + 1 }}</th>
          <td>{{ obj.Railway.RailwayTitleJa }} ({{ obj.Railway.RailwayTitleEn }})</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>駅名</th>
          <th>StationSameAs</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Stations" :key="index">
          <th>{{ index + 1 }}</th>
          <td>{{ obj.Station.StationTitleJa }} ({{ obj.Station.StationTitleEn }})</td>
          <td>{{ obj.Station.SameAs }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  async asyncData(context) {
    const url = `/train/passengerSurveys/${context.params.sameAs}`
    const { data, error } = await context.app.$axios.$get(url)
    if (error) {
      context.error({ message: error })
      return
    }
    return { data: data }
  }
}
</script>
