<template>
  <table class="table table-sm table-hover">
    <thead>
      <tr>
        <th scope="col">#</th>
        <!-- <th scope="col">@context</th>
        <th scope="col">@id</th>
        <th scope="col">@type</th> -->
        <th scope="col">dc:date</th>
        <th scope="col">odpt:operator</th>
        <th scope="col">odpt:railway</th>
        <th scope="col">odpt:station</th>
        <th scope="col">odpt:passengerSurveyObject</th>
        <th scope="col">owl:sameAs</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(obj, index) in data" :key="index">
        <th scope="row">{{ index }}</th>
        <!-- <td>{{ obj['@context'] }}</td>
        <td>{{ obj['@id'] }}</td>
        <td>{{ obj['@type'] }}</td> -->
        <td>{{ obj['dc:date'] }}</td>
        <td>{{ obj['odpt:operator'] }}</td>
        <td>{{ obj['odpt:railway'] }}</td>
        <td>{{ obj['odpt:station'] }}</td>
        <td>{{ obj['odpt:passengerSurveyObject'] }}</td>
        <td>{{ obj['owl:sameAs'] }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  async asyncData(context) {
    const data = await context.app.$axios.$get("/odpt:PassengerSurvey")
    return { data: data }
  }
}
</script>
