<template>
  <div>
    <p class="font-weight-bold">
      <a :href="docpath" target="_blank">公式ドキュメント</a>
    </p>
    <table class="table table-sm table-hover">
      <thead>
        <tr>
          <th scope="col">#</th>
          <th scope="col">Category</th>
          <th scope="col">Type</th>
          <th scope="col">Name</th>
          <th scope="col">Desc</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="obj in data" :key="obj.ID">
          <th scope="row">{{ obj.ID }}</th>
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
      docpath: "https://developer-tokyochallenge.odpt.org/documents"
    }
  },
  async asyncData(context) {
    const { data } = await context.app.$axios.$get("/shared/categories")
    return { data: data }
  }
}
</script>
