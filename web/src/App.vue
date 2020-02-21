<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png" /> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <div class="columns">
      <div class="column is-3 menuPrincipal">
        <div class="menuPrincipalContainer">
          <i class="fas fa-users fa-2x white" style="color: white;"></i>
          <div class="buttonsMenu">
            <div class="containerButton">
              <i class="fas fa-home fa-2x"></i>
              <p>Home</p>
            </div>
          </div>
        </div>
      </div>
      <div class="column menuTweets">
        <div class="sectionTitle">
          <p>Home</p>
        </div>
        <div class="columns sectionTweets">
          <div class="tweet newTweet">
            <div class="column is-2 avatar">
              <i class="far fa-user-circle fa-4x"></i>
            </div>
            <div class="column infoTweet">
              <div class="messageTweet">
                <input
                  class="input titleTweet"
                  type="text"
                  placeholder="Title"
                  v-model="newTweet.title"
                />
                <input
                  class="input"
                  type="text"
                  placeholder="Description"
                  v-model="newTweet.description"
                />
              </div>
              <div class="sectionSubmit sectionIcons">
                <button
                  :disabled="sendTweet"
                  v-on:click="saveTweet()"
                  class="button is-link is-rounded"
                >Tweet</button>
              </div>
            </div>
          </div>
          <div v-for="tweet in tweets" v-bind:key="tweet.thread_id" class="tweet">
            <div style="display: flex; width: 100%;">
              <div class="column is-2 avatar">
                <i class="far fa-user-circle fa-4x"></i>
              </div>
              <div class="column infoTweet">
                <div class="userName">
                  <div style="display:flex">
                    <p class="name">User</p>
                    <p class="user">@user</p>
                  </div>
                  <i class="fas fa-pen" v-on:click="openModal(tweet)"></i>
                </div>
                <div class="messageTweet">
                  <p class="titleTweet">{{ tweet.title }}</p>
                  <p>{{ tweet.description }}</p>
                </div>
                <div class="sectionIcons">
                  <i class="far fa-comment fa-1x" v-on:click="openComments(tweet)"></i>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="column is-3 trends">
        <article class="panel is-primary">
          <p class="panel-heading">Trends for you</p>
          <a class="panel-block">
            <div>
              <p class="subTitle">Trending</p>
              <p>#ParoNacional</p>
            </div>
          </a>
          <a class="panel-block">
            <div>
              <p class="subTitle">Trending</p>
              <p>#AMLONoEstaSolo</p>
            </div>
          </a>
          <a class="panel-block">
            <div>
              <p class="subTitle">Trending</p>
              <p>#FeminicidasDetenidos</p>
            </div>
          </a>
          <a class="panel-block">
            <div>
              <p class="subTitle">Trending</p>
              <p>#LopezAceptamosTuRenuncia</p>
            </div>
          </a>
        </article>
      </div>
    </div>
    <div class="modal" :class="{ 'is-active': showModal }">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <i class="fas fa-times fa-2x" v-on:click="closeModal()"></i>
        </header>
        <section class="modal-card-body">
          <div class="infoTweet">
            <div class="messageTweet">
              <input
                class="input titleTweet"
                type="text"
                placeholder="Title"
                v-model="editedTweet.title"
              />
              <input
                class="input"
                type="text"
                placeholder="Description"
                v-model="editedTweet.description"
              />
            </div>
          </div>
        </section>
        <footer class="modal-card-foot">
          <button class="button is-success" v-on:click="updateTweet()">Save changes</button>
          <button class="button" v-on:click="closeModal()">Cancel</button>
        </footer>
      </div>
    </div>
    <div class="modal modalComments" :class="{ 'is-active': showComments }">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <i class="fas fa-times fa-2x" v-on:click="closeComments()"></i>
        </header>
        <section class="modal-card-body">
          <div class="columns tweet newTweet">
            <div class="column is-2 avatar">
              <i class="far fa-user-circle fa-4x"></i>
            </div>
            <div class="column infoTweet">
              <div class="messageTweet">
                <p class="titleTweet">{{ selectedTweet.title }}</p>
                <p>{{ selectedTweet.description }}</p>
              </div>
            </div>
          </div>
          <div class="columns" v-for="comment in comments" v-bind:key="comment.replie_id">
            <div class="column is-1"></div>
            <div class="column comentarioTweet">
              <p :id="comment.replie_id">{{ comment.content }}</p>
              <input
                :id="comment.replie_id + 'editMode'"
                class="input titleTweet"
                style="display:none"
                type="text"
                v-model="comment.content"
              />
            </div>
            <div class="column is-1">
              <i class="fas fa-pen" v-if="editMode == false" v-on:click="openEditMode(comment)"></i>
            </div>
          </div>
        </section>
        <footer class="modal-card-foot">
          <input
            v-if="editMode == false"
            class="input titleTweet"
            type="text"
            placeholder="Reply"
            v-model="newReply"
          />
          <button v-if="editMode == false" class="button is-success" v-on:click="saveReply()">Reply</button>

          <div v-if="editMode == true">
            <button class="button is-success" v-on:click="updateReply()">Save changes</button>
            <button class="button" v-on:click="exitEditMode()">Cancel</button>
          </div>
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
      showComments: false,
      editMode: false,
      newReply: "",
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
    }
  },
  methods: {
    exitEditMode() {
      this.editMode = false;
    },
    openEditMode(comment) {
      console.log(comment);
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
    openComments(tweet) {
      this.showComments = true;
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
      // this.editedTweet.id = tweet.thread_id;
      // this.editedTweet.title = tweet.title;
      // this.editedTweet.description = tweet.description;
    },
    closeComments() {
      this.showComments = false;
      this.newReply = "";
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
          alert("Thread created sucessfully!");
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
      console.log(this.newReply);
      console.log(this.selectedTweet);
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
    updateReply() {
      const form = new FormData();
      form.append("content", "Update Reply with Vue");
      axios.post("http://localhost:8081/api/replies/6", form).then(response => {
        axios
          .get(`http://localhost:8081/api/replies/thread/8`)
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
    }
  }
};
</script>

<style>
@import "./styles/style.css";
</style>
