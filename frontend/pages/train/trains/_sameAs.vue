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
          <th>進行方向</th>
          <td>{{ train.RailDirection.Title }}</td>
        </tr>
        <tr>
          <th>列車番号</th>
          <td>{{ train.TrainNumber }}</td>
        </tr>
        <tr>
          <th>RailwaySameAs</th>
          <td>{{ train.RailwaySameAs }}</td>
        </tr>
        <tr>
          <th>TrainTypeSameAs</th>
          <td>{{ train.TrainTypeSameAs }}</td>
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
        </tbody>
      </table>
      <table class="table table-sm table-hover">
        <thead>
          <tr>
            <th>#</th>
            <th/>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(obj, index) in train.Objects" :key="index">
            <th>{{ index + 1 }}</th>
            <td>{{ obj }}</td>
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
      <table class="table table-sm table-hover">
        <thead>
          <tr>
            <th>#</th>
            <th>次の時刻表</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(obj, index) in train.Nexts" :key="index">
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
          <tr v-for="(obj, index) in train.Previous" :key="index">
            <th>{{ index + 1 }}</th>
            <td>{{ obj.TrainTimetableSameAs }}</td>
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
          <tr v-for="(obj, index) in train.ViaRailways" :key="index">
            <th>{{ index + 1 }}</th>
            <td>{{ obj.RailwaySameAs }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  async asyncData(context) {
    const url = `/train/trains/${context.params.sameAs}`
    const train = await context.app.$axios.$get(url)
    if (train.error) {
      context.error({ message: train.error })
      return
    }
    const timetableUrl = `/train/trainTimetables/${train.data.SameAs}`
    const timetable = await context.app.$axios.$get(timetableUrl)
    if (timetable.error) {
      timetable.data = null
    }
    return { train: train.data, timetable: timetable.data }
  }
}
</script>
