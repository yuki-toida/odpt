<template>
  <div class="container">
    <nav class="navbar navbar-expand-lg">
      <nuxt-link class="navbar-brand" to="/">refodpt</nuxt-link>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"/>
      </button>
      <div id="navbarNav" class="collapse navbar-collapse">
        <ul class="navbar-nav">
          <li class="nav-item active">
            <nuxt-link class="nav-link" to="/">api <span class="sr-only">(current)</span></nuxt-link>
          </li>
        </ul>
      </div>
    </nav>
    <nuxt/>
  </div>
</template>
