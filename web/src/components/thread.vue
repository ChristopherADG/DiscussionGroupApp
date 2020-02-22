<template>
  <div class="card" style="margin-top: 2px">
    <div class="card-content tweet" style="display: flex; width: 100%;">
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
        <div class="messageTweet">
          <p class="titleTweet">{{ tweet.title }}</p>
          <p>{{ tweet.description }}</p>
        </div>
      </div>
      <div class="media-right">
        <button class="button is-white" v-on:click="openModal(tweet)">
          <i style="cursor: pointer;" class="fas fa-pen"></i>
        </button>
        <br />
        <button class="button is-white" v-on:click="openComments(tweet)">
          <i style="cursor: pointer;" class="far fa-comment fa-1x"></i>
        </button>
      </div>
    </div>
    <div :id="tweet.thread_id" style="display:none">
      <div
        v-for="comment in comments"
        v-bind:key="comment.replie_id"
        class="columns is-gapless is-vcentered repliesContainer"
      >
        <div class="column is-1"></div>
        <div class="column replyText">- {{comment.content}}</div>
        <div class="column is-2">
          <button class="button is-light" v-on:click="openModelComment(comment)">
            <i class="fas fa-pen icon" style="cursor: pointer;"></i>
          </button>
        </div>
      </div>
      <div>
        <div class="columns is-gapless is-vcentered repliesContainer">
          <div class="column is-1"></div>
          <div class="column">
            <div class>
              <input class="input" type="text" placeholder="Reply" v-model="newReply" />
            </div>
          </div>
          <div class="column is-1">
            <div class="sectionSubmit sectionIcons">
              <button
                v-on:click="saveReply()"
                :disabled="sendReply"
                class="button is-link is-outlined send"
              >
                <i class="fas fa-comment-dots"></i>
              </button>
            </div>
          </div>
          <div class="column is-1"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    tweet: {
      required: true
    },
    comments: {
      required: true
    },
    reply: {
      required: true
    }
  },
  data() {
    return {
      newReply: this.newReply
    };
  },
  computed: {
    // a computed getter
    sendReply: function() {
      if (this.newReply !== "") {
        return false;
      } else {
        return true;
      }
    }
  },
  methods: {
    openModal(tweet) {
      this.$emit("openModal", tweet);
    },
    openComments(tweet) {
      this.$emit("openComments", tweet);
    },
    openModelComment(comment) {
      this.$emit("openModelComment", comment);
    },
    saveReply() {
      this.$emit("saveReply", this.newReply);
      this.newReply = "";
    }
  }
};
</script>