<template>
  <div>
    <nuxt-link class="nav-link" to="/admin">管理画面</nuxt-link>
    <br>
    <select v-model="fromStation">
      <option v-for="station in stations" :key="station.ID" :value="station">
        {{ station.Operator.OperatorTitleJa }} : {{ station.Railway.RailwayTitleJa }} : {{ station.StationTitleJa }}
      </option>
    </select>
    <br>
    <select v-model="toStation">
      <option v-for="station in stations" :key="station.ID" :value="station">
        {{ station.Operator.OperatorTitleJa }} : {{ station.Railway.RailwayTitleJa }} : {{ station.StationTitleJa }}
      </option>
    </select>
    <br>
    <button @click="click">決定</button>
  </div>
</template>

<script>
export default {
  async asyncData(context) {
    const { data, error } = await context.app.$axios.$get("/app/connectings")
    if (error) {
      context.error({ message: error })
      return
    }
    const firstStation = data[0]
    return {
      fromStation: firstStation,
      toStation: firstStation,
      stations: data
    }
  },
  methods: {
    async click() {
      const params = {
        from: this.fromStation.StationTitleJa,
        to: this.toStation.StationTitleJa
      }
      const { data, error } = await this.$axios.$post("/app/routes", params)
      if (error) {
        alert(error)
        return
      }
      console.log(data)
    }
  }
}
</script>
