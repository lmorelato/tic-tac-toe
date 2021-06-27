<template>
  <div id="app">
    <div class="mb-4">
      <b-navbar type="dark" variant="dark">
        <b-navbar-nav>
          <b-navbar-brand class="mx-4">Tic-Tac-Toe</b-navbar-brand>
          <b-nav-item href="#" @click="newGame()">New Game</b-nav-item>
          <b-nav-item href="#" @click="joinGame()">Join Game</b-nav-item>
          <!-- <b-nav-item href="#">About</b-nav-item> -->
        </b-navbar-nav>
      </b-navbar>
    </div>

    <div class="container">
      <b-alert :show="alertShow" :variant="alertVariant"
        ><span v-html="alertMessage"></span>
      </b-alert>

      <PlayersInfo :player1="player1" :player2="player2" />

      <div class="row">
        <div class="col mt-2">
          <GameBoard />
        </div>
        <div class="col">
          <Logs />
        </div>
      </div>

      <LaunchGameModal
        ref="launchGameModal"
        @gameStarted="onGameStarted"
        @gameJoined="onGameJoined"
      />
    </div>
  </div>
</template>

<script>
import GameBoard from "./components/GameBoard.vue";
import PlayersInfo from "./components/PlayersInfo.vue";
import LaunchGameModal from "./components/LaunchGameModal.vue";
import Logs from "./components/Logs.vue";

export default {
  name: "App",
  components: {
    GameBoard,
    PlayersInfo,
    LaunchGameModal,
    Logs,
  },
  data: function () {
    return {
      player1: null,
      player2: null,
      sessionId: null,
      modalShow: false,
      modalAction: "",
      alertShow: false,
      alertMessage: "",
      alertVariant: "",
    };
  },
  methods: {
    showAlertMessage: function (message, variant) {
      this.alertShow = true;
      this.alertMessage = message;
      this.alertVariant = variant;
    },
    hideAlert: function () {
      this.alertShow = false;
    },
    newGame: function () {
      this.$refs.launchGameModal.show("NEW_GAME");
    },
    joinGame: function () {
      this.$refs.launchGameModal.show("JOIN_GAME");
    },
    about: function () {},
    onGameStarted: function (gameInfo) {
      this.player1 = gameInfo.name;
      this.showAlertMessage(
        "Invite someone to start playing sharing this game code: <span class='h4 mx-1'>3kj4</span>",
        "info"
      );
    },
    onGameJoined: function (gameInfo) {
      this.player2 = gameInfo.name;
      this.sessionId = gameInfo.sessionId;
      this.hideAlert();
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
