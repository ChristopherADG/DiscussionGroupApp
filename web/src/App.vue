<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png" /> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <img src="./assets/logoBird.png" style="margin-top: 2px;height: 48px;" />
      </div>
    </nav>
    <br />
    <home />
  </div>
</template>

<script>
import home from "./views/home.vue";

export default {
  name: "App",
  components: {
    home
  }
};
</script>

<style>
@import "./styles/style.css";
</style>
