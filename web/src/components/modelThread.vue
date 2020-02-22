<template>
  <div>
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Edit Post</p>
        <button class="delete" aria-label="close" v-on:click="closeModal()"></button>
      </header>
      <section class="modal-card-body">
        <div class="media">
          <figure class="media-left">
            <p class="image is-64x64 is-circle">
              <img
                alt="Image"
                src="http://qtwitter-demo.qcode.in/img/avatar-placeholder.svg"
                style="border-radius: 50%"
              />
            </p>
          </figure>
          <div class="media-content">
            <div class="field">
              <div class="control">
                <input
                  class="input titleInput"
                  type="text"
                  placeholder="Title"
                  v-model="editedTweet.title"
                />
                <textarea
                  class="textarea"
                  type="text"
                  placeholder="Text"
                  rows="2"
                  v-model="editedTweet.description"
                />
              </div>
            </div>
          </div>
        </div>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" v-on:click="updateThread()">Save changes</button>
        <button class="button" v-on:click="closeModal()">Cancel</button>
      </footer>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    editedTweet: {
      required: true
    }
  },
  methods: {
    closeModal() {
      this.$emit("closeComponent");
    },
    updateThread() {
      this.$emit("updateThread", this.selectedReply);
    }
  }
};
</script>