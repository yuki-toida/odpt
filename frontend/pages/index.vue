<template>
  <div>
    <p class="font-weight-bold">
      <a :href="doclink('_共通クラス')" target="_blank">SharedAPI</a>
    </p>
    <ul class="list-group">
      <li class="list-group-item">
        <nuxt-link to="/raw/calendar">odpt:Calendar</nuxt-link>
      </li>
      <li class="list-group-item">
        <nuxt-link to="/raw/operator">odpt:Operator</nuxt-link>
      </li>
    </ul>
    <p class="font-weight-bold mt-3">
      <a :href="doclink('_odpt_train_api')" target="_blank">TrainAPI</a>
    </p>
    <ul class="list-group">
      <li class="list-group-item">
        <nuxt-link to="/raw/passengerSurvey">odpt:PassengerSurvey</nuxt-link>
      </li>
      <li class="list-group-item">
        <nuxt-link to="/raw/railDirection">odpt:RailDirection</nuxt-link>
      </li>
      <li class="list-group-item">
        <nuxt-link to="/raw/railway">odpt:Railway</nuxt-link>
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
