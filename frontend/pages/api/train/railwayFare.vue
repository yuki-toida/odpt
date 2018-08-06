<template>
  <table class="table table-sm table-hover">
    <thead>
      <tr>
        <th scope="col">#</th>
        <!-- <th scope="col">@context</th>
        <th scope="col">@id</th>
        <th scope="col">@type</th>
        <th scope="col">dc:date</th> -->
        <th scope="col">odpt:operator</th>
        <th scope="col">odpt:fromStation</th>
        <th scope="col">odpt:toStation</th>
        <th scope="col">odpt:icCardFare</th>
        <th scope="col">odpt:ticketFare</th>
        <th scope="col">odpt:childIcCardFare</th>
        <th scope="col">odpt:childTicketFare</th>
        <th scope="col">owl:sameAs</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(obj, index) in data" :key="index">
        <th scope="row">{{ index }}</th>
        <!-- <td>{{ obj['@context'] }}</td>
        <td>{{ obj['@id'] }}</td>
        <td>{{ obj['@type'] }}</td>
        <td>{{ obj['dc:date'] }}</td> -->
        <td>{{ obj['odpt:operator'] }}</td>
        <td>{{ obj['odpt:fromStation'] }}</td>
        <td>{{ obj['odpt:toStation'] }}</td>
        <td>{{ obj['odpt:icCardFare'] }}</td>
        <td>{{ obj['odpt:ticketFare'] }}</td>
        <td>{{ obj['odpt:childIcCardFare'] }}</td>
        <td>{{ obj['odpt:childTicketFare'] }}</td>
        <td>{{ obj['owl:sameAs'] }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  async asyncData(context) {
    const { data } = await context.app.$axios.$get("/train/railwayFare")
    return { data: data }
  }
}
</script>
