<template>
  <table class="table table-sm table-hover">
    <thead>
      <tr>
        <th>#</th>
        <th>路線名</th>
        <th>路線コード</th>
        <th>運行会社</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(obj, index) in data" :key="index">
        <th>{{ index }}</th>
        <td>
          <nuxt-link :to="path(obj.SameAs)" class="nav-link">{{ obj.Title }}</nuxt-link>
        </td>
        <td>{{ obj.LineCode }}</td>
        <td>{{ obj.Operator.OperatorTitleJa }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  async asyncData(context) {
    const { data } = await context.app.$axios.$get("/train/railways")
    return { data: data }
  },
  methods: {
    path(sameAs) {
      return `/train/railways/${sameAs}`
    }
  }
}
</script>
