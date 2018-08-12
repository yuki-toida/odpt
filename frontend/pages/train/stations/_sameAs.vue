<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>StationID</th>
          <td>{{ data.ID }}</td>
        </tr>
        <tr>
          <th>StationSameAs</th>
          <td>{{ data.SameAs }}</td>
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
          <th>駅コード</th>
          <td>{{ data.StationCode }}</td>
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
          <th>路線名(Ja)</th>
          <td>{{ data.Railway.RailwayTitleJa }}</td>
        </tr>
        <tr>
          <th>路線名(En)</th>
          <td>{{ data.Railway.RailwayTitleEn }}</td>
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
          <th>#</th>
          <th>乗換可能路線名</th>
          <th>RailwaySameAs</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.ConnectingRailways" :key="index">
          <th>{{ index + 1 }}</th>
          <td>
            <nuxt-link v-if="obj.Railway" :to="railwayPath(obj.Railway.SameAs)" class="nav-link">
              {{ obj.Railway.RailwayTitleJa }}({{ obj.Railway.RailwayTitleEn }})
            </nuxt-link>
          </td>
          <td>{{ obj.RailwaySameAs }}</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>乗降人員調査</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.PassengerSurveys" :key="index">
          <th>{{ index + 1 }}</th>
          <td>
            <nuxt-link :to="passengerSurveyPath(obj.PassengerSurveySameAs)" class="nav-link">
              {{ obj.PassengerSurveySameAs }}
            </nuxt-link>
          </td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>駅時刻表</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Timetables" :key="index">
          <th>{{ index + 1 }}</th>
          <td>
            <nuxt-link v-if="obj.StationTimetable" :to="stationTimetablePath(obj.StationTimetable.SameAs)" class="nav-link">
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
