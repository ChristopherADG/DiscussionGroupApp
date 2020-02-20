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
        <div class="sectionTitle"><p>Home</p></div>
        <div class="columns sectionTweets">
          <div
            v-for="tweet in tweets"
            v-bind:key="tweet.thread_id"
            class="tweet"
          >
            <div class="column is-2 avatar">
              <i class="far fa-user-circle fa-4x"></i>
            </div>
            <div class="column infoTweet">
              <div class="userName">
                <p class="name">User</p>
                <p class="user">@user</p>
              </div>
              <div class="messageTweet">
                <p class="titleTweet">{{ tweet.title }}</p>
                <p>
                  {{ tweet.description }}
                </p>
              </div>
              <div class="sectionIcons">
                <i class="far fa-comment fa-1x"></i>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="column is-3 trends">
        <article class="panel is-primary">
          <p class="panel-heading">
            Trends for you
          </p>
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

    <!-- <div id="twitter">
      <header>
        <div class="inner">
          <div class="headerSection">
            <router-link to="/home" class="link">Home</router-link>
          </div>
          <div class="headerSection"></div>
          <div class="headerSection">
            <button class="tweetLink">Tweet</button>
          </div>
        </div>
      </header>

      <router-view></router-view>
    </div>
    <Home /> -->
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
      newTweet: ""
    };
  },
  mounted() {
    axios.get("http://localhost:8081/api/threads").then(response => {
      // JSON responses are automatically parsed.
      this.tweets = response.data;
    });
  }
};
</script>

<style>
#app {
  background-color: rgb(21, 32, 43);
  height: 100%;
}

.columns {
  height: 100%;
}
.column {
  text-align: center;
  border: 1px solid rgb(56, 68, 77);
  height: 100%;
}
.trends {
  height: auto;
}
.menuPrincipal {
  text-align: left;
  height: auto;
}
.menuPrincipalContainer {
  padding-left: 3em;
  padding-top: 2em;
}
.menuPrincipalContainer i {
  margin-bottom: 25px;
}
.buttonsMenu .containerButton {
  display: flex;
  width: fit-content;
}
.buttonsMenu .containerButton:hover {
  cursor: pointer;
}
.buttonsMenu .containerButton i {
  margin-bottom: 0px;
  padding-right: 15px;
  color: rgba(29, 161, 242, 1);
}
.buttonsMenu .containerButton p {
  font-size: 24px;
  color: rgba(29, 161, 242, 1);
  font-weight: bold;
}

.menuTweets {
  text-align: left;
}
.menuTweets .sectionTitle {
  padding-top: 2em;
  padding-bottom: 1em;
  padding-left: 1em;
  border-bottom: 1px solid rgb(56, 68, 77);
}
.menuTweets .sectionTitle p {
  color: white;
  font-weight: bold;
  font-size: 24px;
}

.menuTweets .sectionTweets {
  padding-top: 2em;
  flex-direction: column;
}

.sectionTweets .tweet {
  display: flex;
  height: fit-content;
  border-bottom: 1px solid rgb(56, 68, 77);
}

.menuTweets .sectionTweets .avatar {
  color: rgb(136, 153, 166);
}

.menuTweets .sectionTweets .infoTweet {
  text-align: left;
}

.messageTweet {
  margin-top: 10px;
  /* max-height: 150px;
  overflow: scroll; */
}

.messageTweet p {
  color: white;
}

.messageTweet .titleTweet {
  font-size: 18px;
  font-weight: bold;
  text-decoration: underline;
}

.menuTweets .sectionTweets .infoTweet .sectionIcons {
  padding-top: 10px;
  color: rgb(136, 153, 166);
  text-align: right;
  padding-right: 25px;
}

.menuTweets .sectionTweets .infoTweet .sectionIcons i:hover {
  cursor: pointer;
}

.menuTweets .sectionTweets .infoTweet .userName {
  display: flex;
  color: white;
}
.menuTweets .sectionTweets .infoTweet .userName p {
  padding-right: 10px;
}

.menuTweets .sectionTweets .infoTweet .userName .name {
  font-weight: bolder;
  font-size: 18px;
}
.menuTweets .sectionTweets .infoTweet .userName .user {
  color: rgb(136, 153, 166);
  font-size: 18px;
}

.menuTweets .column {
  border: 0px;
}

.panel {
  margin-top: 2em;
  margin-left: 1em;
  margin-right: 1em;
  background-color: rgb(25, 39, 52);
}

#app .panel .panel-heading {
  background-color: rgb(25, 39, 52);
  border-bottom: 1px solid rgb(56, 68, 77);
  text-align: left;
  font-weight: bold;
  padding-bottom: 10px;
  padding-top: 10px;
}

#app .panel .panel-block {
  border-bottom: 1px solid rgb(56, 68, 77);
  color: white;
  padding: 1em;
}

#app .panel .panel-block .subTitle {
  text-align: left;
  color: rgb(136, 153, 166);
  font-size: 12px;
}

#app .panel .panel-block:hover {
  background-color: #20303e;
}
</style>
