<template>
  <div>
    <div class="notification is-success" v-if="showNotificacion">
      <button class="delete" v-on:click="closeNotification()"></button>
      Post created successfully!
    </div>
    <div class="card">
      <div class="card-content">
        <div class="content">
          Demo of a Discussion Board on
          <b>Vue.js</b> and
          <b>Golang</b>.
          <br />
          <br />
          <a>@Christopher Delgado</a>.
          <br />
          <a>@Luis Sanchez</a>.
          <br />
          <a>@Jesus Ruiz</a>.
          <br />
          <br />
          <small>22 Feb 2020</small>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    showNotificacion: {
      required: true
    }
  }
};
</script>