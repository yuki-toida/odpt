<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>駅名</th>
          <td>{{ data.Title }}</td>
        </tr>
        <tr>
          <th>緯度</th>
          <td>{{ data.Lat }}</td>
        </tr>
        <tr>
          <th>経度</th>
          <td>{{ data.Long }}</td>
        </tr>
        <tr>
          <th>駅コード</th>
          <td>{{ data.StationCode }}</td>
        </tr>
        <tr>
          <th>運行会社(Ja)</th>
          <td>{{ data.Operator.OperatorTitleJa }}</td>
        </tr>
        <tr>
          <th>運行会社(En)</th>
          <td>{{ data.Operator.OperatorTitleEn }}</td>
        </tr>
        <tr>
          <th>駅名(Ja)</th>
          <td>{{ data.StationTitleJa }}</td>
        </tr>
        <tr>
          <th>駅名(En)</th>
          <td>{{ data.StationTitleEn }}</td>
        </tr>
        <tr>
          <th>地物情報</th>
          <td>{{ data.Region }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <caption>乗換可能路線</caption>
      <thead>
        <tr>
          <th>#</th>
          <th>路線名</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.ConnectingRailways" :key="index">
          <th>{{ index }}</th>
          <td>
            <nuxt-link :to="railwayPath(obj.RailwaySameAs)" class="nav-link">
              {{ obj.Railway.Title }}
            </nuxt-link>
          </td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <caption>出入り口</caption>
      <thead>
        <tr>
          <th>#</th>
          <th>出入り口</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Exits" :key="index">
          <th>{{ index }}</th>
          <td>{{ obj.Exit }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <caption>乗降人員数</caption>
      <thead>
        <tr>
          <th>#</th>
          <th>乗降人員数</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.PassengerSurveys" :key="index">
          <th>{{ index }}</th>
          <td>
            <nuxt-link :to="passengerSurveyPath(obj.PassengerSurveySameAs)" class="nav-link">
              {{ obj.PassengerSurveySameAs }}
            </nuxt-link>
          </td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <caption>駅時刻表</caption>
      <thead>
        <tr>
          <th>#</th>
          <th>駅時刻表</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Timetables" :key="index">
          <th>{{ index }}</th>
          <td>
            <nuxt-link :to="stationTimetablePath(obj.StationTimetable.SameAs)" class="nav-link">
              {{ obj.StationTimetable.SameAs }}
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
    const url = `/train/stations/${context.params.sameAs}`
    const { data, error } = await context.app.$axios.$get(url)
    return { data: data, error: error }
  },
  mounted: function() {
    if (this.error) {
      this.$toast.error(this.error)
    }
  },
  methods: {
    railwayPath(sameAs) {
      return `/train/railways/${sameAs}`
    },
    passengerSurveyPath(sameAs) {
      return `/train/passengerSurveys/${sameAs}`
    },
    stationTimetablePath(sameAs) {
      return `/train/stationTimetables/${sameAs}`
    }
  }
}
</script>
