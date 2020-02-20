<template>
  <div>
    <div v-for="tweet in feed" :key="tweet.title" class="tweet">
      <div class="avatarContainer">
        <img :src="tweet.avatar" class="avatar" width="70" height="70" />
      </div>

      <div class="tweetBody">
        <div class="tweetHead">
          <span class="author">
            {{ tweet.title }}
          </span>
          {{ tweet.date }}
        </div>
        <div class="tweetContent">
          {{ tweet.description }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      feed: [],
      comments: []
    };
  },
  mounted() {
    axios.get("http://localhost:8081/api/threads").then(response => {
      // JSON responses are automatically parsed.
      this.feed = response.data;
    });
  }
};
</script>

<style>
.avatarContainer {
  margin-right: 20px;
}
.avatar {
  margin-top: 3px;
  border-radius: 50%;
}
.tweetBody {
  margin-top: 5px;
}
.tweetHead {
  font-size: 13px;
}
.author {
  font-size: 14px;
  font-weight: bold;
}
.tweetContent {
  font-size: 14px;
  line-height: 20px;
  margin-top: 8px;
}
</style>
