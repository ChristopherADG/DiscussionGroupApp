<template>
  <div>
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Edit Replie</p>
        <button class="delete" aria-label="close" v-on:click="closeModelComment()"></button>
      </header>
      <section class="modal-card-body">
        <div class="infoTweet">
          <div class>
            <input
              class="input"
              type="text"
              placeholder="Description"
              v-model="selectedReply.content"
            />
          </div>
        </div>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" v-on:click="updateReply()">Save Reply</button>
        <button class="button" v-on:click="closeModelComment()">Cancel</button>
      </footer>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    selectedReply: {
      required: true
    }
  },
  methods: {
    closeModelComment() {
      this.$emit("closeComponent");
    },
    updateReply() {
      this.$emit("updateReply", this.selectedReply);
    }
  }
};
</script>
