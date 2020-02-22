<template>
  <div class="container">
    <div class="columns">
      <div class="column is-3 menuPrincipal"></div>
      <div class="column menuTweets">
        <div class="sectionTweets">
          <div id="newTweet" class="card">
            <newPost :newTweet="newTweet" v-on:saveTweet="saveTweet" />
            <div v-for="tweet in tweets" v-bind:key="tweet.thread_id">
              <thread
                :tweet="tweet"
                :comments="comments"
                :reply="newReply"
                v-on:openModal="openModal"
                v-on:openComments="openComments"
                v-on:openModelComment="openModelComment"
                v-on:saveReply="saveReply"
              />
            </div>
          </div>
        </div>
      </div>
      <div class="column is-3 info">
        <infoTeam
          :showNotificacion="showNotificacion"
          v-on:closeNotification="closeNotification()"
        />
      </div>
    </div>
    <div class="modal" :class="{ 'is-active': showModal }">
      <modalThread
        :editedTweet="editedTweet"
        v-on:closeComponent="closeModal()"
        v-on:updateThread="updateTweet"
      />
    </div>
    <div class="modal" :class="{ 'is-active': showModalComment }">
      <modalReply
        v-bind:selectedReply="selectedReply"
        v-on:closeComponent="closeModelComment()"
        v-on:updateReply="updateReply"
      />
    </div>
  </div>
</template>
<script>
import modalReply from "../components/modelReply.vue";
import modalThread from "../components/modelThread.vue";
import newPost from "../components/newPost.vue";
import infoTeam from "../components/infoTeam.vue";
import thread from "../components/thread.vue";
import axios from "axios";

export default {
  components: {
    modalReply,
    modalThread,
    newPost,
    infoTeam,
    thread
  },
  data() {
    return {
      tweets: [],
      comments: [],
      newTweet: "",
      showModal: false,
      showModalComment: false,
      showNotificacion: false,
      editMode: false,
      newReply: "",
      selectedReply: {
        id: "",
        content: ""
      },
      selectedTweet: {
        id: 0,
        title: "",
        description: ""
      },
      newTweet: {
        title: "",
        description: ""
      },
      editedTweet: {
        id: 0,
        title: "",
        description: ""
      }
    };
  },
  mounted() {
    axios.get("http://localhost:8081/api/threads").then(response => {
      // JSON responses are automatically parsed.

      this.tweets = response.data.reverse();
    });
  },
  computed: {
    // a computed getter
    sendReply: function() {
      if (this.newReply !== "") {
        return false;
      } else {
        return true;
      }
    },
    editReply: function() {
      if (this.selectedReply !== "") {
        return false;
      } else {
        return true;
      }
    }
  },
  methods: {
    exitEditMode() {
      this.editMode = false;
    },
    openEditMode(comment) {
      this.editMode = true;
    },
    openModal(tweet) {
      this.showModal = true;
      this.editedTweet.id = tweet.thread_id;
      this.editedTweet.title = tweet.title;
      this.editedTweet.description = tweet.description;
    },
    closeModal() {
      this.showModal = false;
      this.editedTweet.id = 0;
      this.editedTweet.title = "";
      this.editedTweet.description = "";
    },
    openNotification() {
      this.showNotificacion = true;
    },
    closeNotification() {
      this.showNotificacion = false;
    },
    openComments(tweet) {
      const element = document.getElementById(tweet.thread_id);

      if (element.style.display != "none") {
        element.style.display = "none";
      } else {
        this.closeComments();
        this.selectedTweet.id = tweet.thread_id;
        this.selectedTweet.title = tweet.title;
        this.selectedTweet.description = tweet.description;
        axios
          .get(`http://localhost:8081/api/replies/thread/${tweet.thread_id}`)
          .then(response => {
            // JSON responses are automatically parsed.
            if (response.data == null) {
              this.comments = [];
            } else {
              this.comments = response.data;
            }
          });
        element.style.display = "block";
      }
    },
    closeComments() {
      this.newReply = "";
      this.selectedReply.id = 0;
      this.selectedReply.content = "";
      for (let i = 0; i <= this.tweets.length - 1; i++) {
        const item = document.getElementById(this.tweets[i].thread_id);
        if (item !== null) {
          item.style.display = "none";
        }
      }
      // this.editedTweet.id = 0;
      // this.editedTweet.title = "";
      // this.editedTweet.description = "";
    },
    saveTweet() {
      const form = new FormData();
      form.append("title", this.newTweet.title);
      form.append("description", this.newTweet.description);
      axios.post("http://localhost:8081/api/threads", form).then(response => {
        axios.get("http://localhost:8081/api/threads").then(response => {
          this.tweets = response.data.reverse();
          this.newTweet.title = "";
          this.newTweet.description = "";
          this.openNotification();
        });
      });
    },

    updateTweet() {
      const form = new FormData();
      form.append("title", this.editedTweet.title);
      form.append("description", this.editedTweet.description);
      let id = this.editedTweet.id;
      axios
        .post("http://localhost:8081/api/threads/" + id, form)
        .then(response => {
          axios.get("http://localhost:8081/api/threads").then(response => {
            this.tweets = response.data.reverse();
            this.editedTweet.id = 0;
            this.editedTweet.title = "";
            this.editedTweet.description = "";
            this.closeModal();
          });
        });
    },
    saveReply(reply) {
      this.newReply = reply;
      const form = new FormData();
      form.append("threadParent_id", this.selectedTweet.id);
      form.append("content", this.newReply);
      axios.post("http://localhost:8081/api/replies", form).then(response => {
        axios
          .get(
            `http://localhost:8081/api/replies/thread/${this.selectedTweet.id}`
          )
          .then(response => {
            // JSON responses are automatically parsed.
            if (response.data == null) {
              this.comments = [];
            } else {
              this.comments = response.data;
              this.newReply = "";
            }
          });
      });
    },
    openModelComment(reply) {
      this.showModalComment = true;
      this.selectedReply.id = reply.replie_id;
      this.selectedReply.content = reply.content;
    },
    closeModelComment() {
      this.showModalComment = false;
    },
    updateReply() {
      const form = new FormData();
      form.append("content", this.selectedReply.content);
      axios
        .post(
          "http://localhost:8081/api/replies/" + this.selectedReply.id,
          form
        )
        .then(response => {
          axios
            .get(
              `http://localhost:8081/api/replies/thread/` +
                this.selectedTweet.id
            )
            .then(response => {
              // JSON responses are automatically parsed.
              if (response.data == null) {
                this.comments = [];
              } else {
                this.comments = response.data;
                this.newReply = "";
              }
            });
        });

      this.closeModelComment();
    }
  }
};
</script>