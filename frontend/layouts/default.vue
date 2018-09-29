<template>
  <div class="container">
    <nav class="navbar navbar-expand-lg">
      <nuxt-link class="navbar-brand" to="/">トップ</nuxt-link>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"/>
      </button>
      <div id="navbarNav" class="collapse navbar-collapse">
        <ul class="navbar-nav">
          <li class="nav-item active">
            <nuxt-link class="nav-link" to="/admin/railways">鉄道</nuxt-link>
          </li>
          <li class="nav-item active">
            <nuxt-link class="nav-link" to="/admin/trains">列車位置情報(RT)</nuxt-link>
          </li>
          <li class="nav-item active">
            <nuxt-link class="nav-link" to="/admin/trainInformations">列車運行情報(RT)</nuxt-link>
          </li>
        </ul>
      </div>
    </nav>
    <nuxt/>
  </div>
</template>
