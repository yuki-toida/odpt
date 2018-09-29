<template>
  <div>
    <dl>
      <dt>マスタDB最終更新時刻</dt>
      <dd>{{ time.MasterAt }}</dd>
      <dt>リアルタイムDB最終更新時刻</dt>
      <dd>{{ time.TranAt }}</dd>
    </dl>
    <p class="font-weight-bold">
      <a :href="docpath" target="_blank">公式ドキュメント</a>
    </p>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th scope="col">#</th>
          <th scope="col">種別</th>
          <th scope="col">rdf:type</th>
          <th scope="col">名称</th>
          <th scope="col">詳細</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in info" :key="index">
          <th scope="row">{{ index + 1 }}</th>
          <td>{{ obj.Category }}</td>
          <td>{{ obj.Type }}</td>
          <td>{{ obj.Name }}</td>
          <td>{{ obj.Desc }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      docpath: "https://developer-tokyochallenge.odpt.org/documents",
      info: [
        {
          Category: "共通",
          Type: "odpt:Calendar",
          Name: "曜日・日付区分",
          Desc: "曜日・日付区分が記述される"
        },
        {
          Category: "共通",
          Type: "odpt:Operator",
          Name: "事業者",
          Desc: "公共交通機関の事業者が記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:PassengerSurvey",
          Name: "駅別乗降人員",
          Desc: "駅の乗降数集計結果が記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:RailDirection",
          Name: "進行方向",
          Desc: "上り、下りなど、列車の進行方向が記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:Railway",
          Name: "路線情報",
          Desc: "鉄道における路線、運行系統が記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:RailwayFare",
          Name: "運賃情報",
          Desc: "運賃が記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:Station",
          Name: "駅情報",
          Desc: "駅の名称や緯度系など、駅に関する情報が記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:StationTimetable",
          Name: "駅時刻表",
          Desc: "駅を発着する列車の時刻表情報が記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:Train",
          Name: "列車情報",
          Desc: "列車の位置、行き先など、列車のリアルタイムな情報が記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:TrainTimetable",
          Name: "列車時刻表",
          Desc: "列車がどの駅にいつ到着するか、出発するか記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:TrainInformation",
          Name: "運行情報",
          Desc: "鉄道路線のリアルタイムな運行状況が記述される"
        },
        {
          Category: "鉄道",
          Type: "odpt:TrainType",
          Name: "列車種別",
          Desc: "普通、快速など、列車の種別が記述される"
        },
        {
          Category: "バス",
          Type: "odpt:Bus",
          Name: "バス情報",
          Desc:
            "バス車両の位置、行先など、バス車両のリアルタイムな情報が記述される"
        },
        {
          Category: "バス",
          Type: "odpt:BusTimetable",
          Name: "バス時刻表",
          Desc:
            "バスがあるバス停、バス標柱にいつ到着するか、いつ出発するかが記述される"
        },
        {
          Category: "バス",
          Type: "odpt:BusroutePattern",
          Name: "バス運行路線情報",
          Desc: "バス運行における運行路線情報が記述される"
        },
        {
          Category: "バス",
          Type: "odpt:BusroutePatternFare",
          Name: "バス運賃情報",
          Desc: "バス運賃が記述される"
        },
        {
          Category: "バス",
          Type: "odpt:BusstopPole",
          Name: "バス時刻表",
          Desc:
            "バスがあるバス停、バス標柱にいつ到着するか、いつ出発するかが記述される"
        },
        {
          Category: "バス",
          Type: "odpt:BusstopPoleTimetable",
          Name: "バス停標柱時刻表",
          Desc:
            "バスがあるバス停標柱にいつ到着するか、いつ出発するかが記述される"
        },
        {
          Category: "航空機",
          Type: "odpt:FlightInfomationArrival",
          Name: "フライト出発情報",
          Desc: "空港に当日到着する航空機のリアルタイムな情報が記述される"
        },
        {
          Category: "航空機",
          Type: "odpt:FlightInformationDeparture",
          Name: "フライト出発情報",
          Desc: "空港に当日到着する航空機のリアルタイムな情報が記述される"
        },
        {
          Category: "航空機",
          Type: "odpt:FlightSchedule",
          Name: "航空機月間時刻表",
          Desc: "航空機の予定される発着時刻情報が記述される"
        }
      ]
    }
  },
  async asyncData(context) {
    const { data, error } = await context.app.$axios.$get("/admin/time")
    if (error) {
      context.error({ message: error })
      return
    }
    return { time: data }
  }
}
</script>
