<template>
  <div>
    <table class="table table-sm table-hover">
      <tbody>
        <tr>
          <th>TrainInformationID</th>
          <td>{{ data.ID }}</td>
        </tr>
        <tr>
          <th>TrainInformationSameAs</th>
          <td>{{ data.SameAs }}</td>
        </tr>
        <tr>
          <th>運行会社</th>
          <td>{{ data.Operator.OperatorTitleJa }} ({{ data.Operator.OperatorTitleEn }})</td>
        </tr>
        <tr>
          <th>発生路線</th>
          <td>{{ data.Railway.RailwayTitleJa }} ({{ data.Railway.RailwayTitleEn }})</td>
        </tr>
        <tr>
          <th>発生場所起点</th>
          <td>{{ data.StationFrom.StationTitleJa }} ({{ data.StationFrom.StationTitleEn }})</td>
        </tr>
        <tr>
          <th>発生場所終点</th>
          <td>{{ data.StationTo.StationTitleJa }} ({{ data.StationTo.StationTitleEn }})</td>
        </tr>
        <tr>
          <th>復旧見込時刻</th>
          <td>{{ data.ResumeEstimate }}</td>
        </tr>
        <tr>
          <th>発生時刻</th>
          <td>{{ data.TimeOfOrigin }}</td>
        </tr>
        <tr>
          <th>発生場所</th>
          <td>{{ data.TrainInformationAreaJa }} ({{ data.TrainInformationAreaEn }})</td>
        </tr>
        <tr>
          <th>発生理由</th>
          <td>{{ data.TrainInformationCauseJa }} ({{ data.TrainInformationCauseEn }})</td>
        </tr>
        <tr>
          <th>鉄道種類</th>
          <td>{{ data.TrainInformationKindJa }} ({{ data.TrainInformationKindEn }})</td>
        </tr>
        <tr>
          <th>運転方向</th>
          <td>{{ data.TrainInformationLineJa }} ({{ data.TrainInformationLineEn }})</td>
        </tr>
        <tr>
          <th>発生区間</th>
          <td>{{ data.TrainInformationRangeJa }} ({{ data.TrainInformationRangeEn }})</td>
        </tr>
        <tr>
          <th>運行状況</th>
          <td>{{ data.TrainInformationStatusJa }} ({{ data.TrainInformationStatusEn }})</td>
        </tr>
        <tr>
          <th>運行情報</th>
          <td>{{ data.TrainInformationTextJa }} ({{ data.TrainInformationTextEn }})</td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th>#</th>
          <th>振替路線</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(obj, index) in data.Railways" :key="index">
          <th>{{ index + 1 }}</th>
          <td>{{ obj.Railway.RailwayTitleJa }} ({{ obj.Railway.RailwayTitleEn }})</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  async asyncData(context) {
    const url = `/train/trainInformations/${context.params.sameAs}`
    const { data, error } = await context.app.$axios.$get(url)
    if (error) {
      context.error({ message: error })
      return
    }
    return { data: data }
  }
}
</script>
