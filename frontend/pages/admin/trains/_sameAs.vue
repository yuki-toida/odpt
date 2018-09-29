<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>TrainID</th>
          <td>{{ train.ID }}</td>
        </tr>
        <tr>
          <th>TrainSameAs</th>
          <td>{{ train.SameAs }}</td>
        </tr>
        <tr>
          <th>車両数</th>
          <td>{{ train.CarComposition }}</td>
        </tr>
        <tr>
          <th>遅延時間(秒)</th>
          <td>{{ train.Delay }}</td>
        </tr>
        <tr>
          <th>列車順序</th>
          <td>{{ train.Index }}</td>
        </tr>
        <tr>
          <th>列車番号</th>
          <td>{{ train.TrainNumber }}</td>
        </tr>
        <tr>
          <th>列車種別</th>
          <td>{{ train.TrainType.TrainTypeTitleJa }} ({{ train.TrainType.TrainTypeTitleEn }})</td>
        </tr>
        <tr>
          <th>出発駅名</th>
          <td>{{ train.FromStation.StationTitleJa }} ({{ train.FromStation.StationTitleEn }})</td>
        </tr>
        <tr>
          <th>終点駅名</th>
          <td>{{ train.ToStation.StationTitleJa }} ({{ train.ToStation.StationTitleEn }})</td>
        </tr>
        <tr>
          <th>進行方向</th>
          <td>{{ train.RailDirection.RailDirectionTitleJa }} ({{ train.RailDirection.RailDirectionTitleEn }})</td>
        </tr>
        <tr>
          <th>路線名</th>
          <td>{{ train.Railway.RailwayTitleJa }} ({{ train.Railway.RailwayTitleEn }})</td>
        </tr>
        <tr>
          <th>運行会社</th>
          <td>{{ train.Operator.OperatorTitleJa }} ({{ train.Operator.OperatorTitleEn }})</td>
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
        <tr v-for="(obj, index) in train.DestinationStations" :key="index">
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
        <tr v-for="(obj, index) in train.OriginStations" :key="index">
          <th>{{ index + 1 }}</th>
          <td>{{ obj.StationSameAs }}</td>
        </tr>
      </tbody>
    </table>

    <div v-if="timetable">
      <table class="table table-sm table-hover">
        <tbody>
          <tr>
            <th>TrainTimetableID</th>
            <td>{{ timetable.ID }}</td>
          </tr>
          <tr>
            <th>TrainTimetableSameAs</th>
            <td>{{ timetable.SameAs }}</td>
          </tr>
          <tr>
            <th>進行方向</th>
            <td>{{ timetable.RailDirection.Title }}</td>
          </tr>
          <tr>
            <th>列車番号</th>
            <td>{{ timetable.TrainNumber }}</td>
          </tr>
          <tr>
            <th>列車種別</th>
            <td>{{ timetable.TrainType.TrainTypeTitleJa }} ({{ timetable.TrainType.TrainTypeTitleEn }})</td>
          </tr>
          <tr>
            <th>カレンダー</th>
            <td>{{ timetable.Calendar.CalendarTitleJa }} ({{ timetable.Calendar.CalendarTitleEn }})</td>
          </tr>
          <tr>
            <th>進行方向</th>
            <td>{{ timetable.RailDirection.RailDirectionTitleJa }} ({{ timetable.RailDirection.RailDirectionTitleEn }})</td>
          </tr>
          <tr>
            <th>運行会社</th>
            <td>{{ timetable.Operator.OperatorTitleJa }} ({{ timetable.Operator.OperatorTitleEn }})</td>
          </tr>
        </tbody>
      </table>
      <table class="table table-sm table-hover">
        <thead>
          <tr>
            <th>#</th>
            <td>出発時刻</td>
            <td>出発駅名</td>
            <td>到着時刻</td>
            <td>到着駅名</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(obj, index) in timetable.Objects" :key="index">
            <th>{{ index + 1 }}</th>
            <td>{{ obj.DepartureTime }}</td>
            <td>{{ obj.DepartureStation }}</td>
            <td>{{ obj.ArrivalTime }}</td>
            <td>{{ obj.ArrivalStation }}</td>
          </tr>
        </tbody>
      </table>
      <table class="table table-sm table-hover">
        <thead>
          <tr>
            <th>#</th>
            <th>終着駅名</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(obj, index) in timetable.DestinationStations" :key="index">
            <th>{{ index + 1 }}</th>
            <td>{{ obj.Station.StationTitleJa }} ({{ obj.Station.StationTitleEn }})</td>
          </tr>
        </tbody>
      </table>
      <table class="table table-sm table-hover">
        <thead>
          <tr>
            <th>#</th>
            <th>始発駅名</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(obj, index) in timetable.OriginStations" :key="index">
            <th>{{ index + 1 }}</th>
            <td>{{ obj.Station.StationTitleJa }} ({{ obj.Station.StationTitleEn }})</td>
          </tr>
        </tbody>
      </table>
      <table class="table table-sm table-hover">
        <thead>
          <tr>
            <th>#</th>
            <th>次の時刻表</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(obj, index) in timetable.Nexts" :key="index">
            <th>{{ index + 1 }}</th>
            <td>{{ obj.TrainTimetableSameAs }}</td>
          </tr>
        </tbody>
      </table>
      <table class="table table-sm table-hover">
        <thead>
          <tr>
            <th>#</th>
            <th>前の時刻表</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(obj, index) in timetable.Previous" :key="index">
            <th>{{ index + 1 }}</th>
            <td>{{ obj.TrainTimetableSameAs }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  async asyncData(context) {
    const url = `/admin/trains/${context.params.sameAs}`
    const train = await context.app.$axios.$get(url)
    if (train.error) {
      context.error({ message: train.error })
      return
    }
    const timetableUrl = `/admin/trainTimetables/${train.data.SameAs}`
    const timetable = await context.app.$axios.$get(timetableUrl)
    if (timetable.error) {
      timetable.data = null
    }
    return { train: train.data, timetable: timetable.data }
  }
}
</script>
