<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
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
      <caption>乗降数</caption>
      <thead>
        <tr>
          <th>#</th>
          <th>調査年度</th>
          <th>一日平均乗降人員数</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Objects" :key="index">
          <th>{{ index }}</th>
          <td>{{ obj.SurveyYear }}</td>
          <td>{{ obj.PassengerJourneys }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <caption>路線名</caption>
      <thead>
        <tr>
          <th>#</th>
          <th>路線名</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Railways" :key="index">
          <th>{{ index }}</th>
          <td>{{ obj.Railway.Title }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <caption>駅名</caption>
      <thead>
        <tr>
          <th>#</th>
          <th>駅名</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Stations" :key="index">
          <th>{{ index }}</th>
          <td>{{ obj.Station.Title }}</td>
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
    return { data: data, error: error }
  },
  mounted: function() {
    if (this.error) {
      this.$toast.error(this.error)
    }
  }
  // methods: {
  //   path(sameAs) {
  //     return `/train/stations/${sameAs}`
  //   }
  // }
}
</script>
