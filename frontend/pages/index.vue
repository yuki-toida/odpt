<template>
  <div>
    <p class="font-weight-bold">
      <a :href="doclink('_共通クラス')" target="_blank">SharedAPI</a>
    </p>
    <ul class="list-group">
      <li class="list-group-item">
        <nuxt-link to="/api/shared/calendar">odpt:Calendar</nuxt-link>
      </li>
      <li class="list-group-item">
        <nuxt-link to="/api/shared/operator">odpt:Operator</nuxt-link>
      </li>
    </ul>
    <p class="font-weight-bold mt-3">
      <a :href="doclink('_odpt_train_api')" target="_blank">TrainAPI</a>
    </p>
    <ul class="list-group">
      <li class="list-group-item">
        <nuxt-link to="/api/train/passengerSurvey">odpt:PassengerSurvey</nuxt-link>
      </li>
      <li class="list-group-item">
        <nuxt-link to="/api/train/railDirection">odpt:RailDirection</nuxt-link>
      </li>
      <li class="list-group-item">
        <nuxt-link to="/api/train/railway">odpt:Railway</nuxt-link>
      </li>
      <li class="list-group-item">
        <nuxt-link to="/api/train/railwayFare">odpt:RailwayFare</nuxt-link>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  methods: {
    doclink: function(path) {
      return `https://developer-tokyochallenge.odpt.org/documents#${path}`
    }
  }
}
</script>
