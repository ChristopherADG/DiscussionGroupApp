<template>
  <div class="card-content newTweetCard">
    <div class="media">
      <div class="media-left">
        <figure class="image is-32x32">
          <img
            alt="Image"
            src="http://qtwitter-demo.qcode.in/img/avatar-placeholder.svg"
            style="border-radius: 50%"
          />
        </figure>
      </div>
      <div class="media-content">
        <div class="field">
          <div class="control">
            <input
              class="input titleInput"
              type="text"
              placeholder="Title"
              v-model="newTweet.title"
            />
            <textarea
              class="textarea"
              type="text"
              placeholder="Text"
              rows="2"
              v-model="newTweet.description"
            />
          </div>
        </div>
        <div style="display: flow-root;">
          <div class="is-pulled-right">
            <button
              :disabled="sendTweet"
              v-on:click="saveTweet()"
              class="button is-link is-outlined"
            >Post</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    newTweet: {
      required: true
    }
  },
  computed: {
    sendTweet: function() {
      if (this.newTweet.title !== "" && this.newTweet.description !== "") {
        return false;
      } else {
        return true;
      }
    }
  },
  methods: {
    saveTweet() {
      this.$emit("saveTweet");
    }
  }
};
</script>