<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png" /> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item" href="https://bulma.io">
          <img src="https://bulma.io/images/bulma-logo.png" width="112" height="28" />
        </a>
      </div>
    </nav>
    <br />
    <div class="container">
      <div class="columns">
        <div class="column is-3 menuPrincipal"></div>
        <div class="column menuTweets">
          <div class="sectionTweets">
            <div id="newTweet" class="card">
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
              <div v-for="tweet in tweets" v-bind:key="tweet.thread_id">
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
              </div>
            </div>
          </div>
        </div>
        <div class="column is-3 info">
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
      </div>
      <div class="modal" :class="{ 'is-active': showModal }">
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
            <button class="button is-success" v-on:click="updateTweet()">Save changes</button>
            <button class="button" v-on:click="closeModal()">Cancel</button>
          </footer>
        </div>
      </div>
    </div>

    <div class="modal" :class="{ 'is-active': showModalComment }">
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
  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import Home from "./components/home.vue";
import axios from "axios";

export default {
  name: "App",
  components: {
    HelloWorld,
    Home
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
    sendTweet: function() {
      if (this.newTweet.title !== "" && this.newTweet.description !== "") {
        return false;
      } else {
        return true;
      }
    },
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
    saveReply() {
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

<style>
@import "./styles/style.css";
</style>
