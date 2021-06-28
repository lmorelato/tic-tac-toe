<template>
  <div id="app">
    <div class="mb-4">
      <b-navbar type="dark" variant="dark">
        <b-navbar-nav>
          <b-navbar-brand class="mx-4">Tic-Tac-Toe</b-navbar-brand>
          <b-nav-item href="#" @click="newGame()">New Game</b-nav-item>
          <b-nav-item href="#" @click="joinGame()">Join Game</b-nav-item>
        </b-navbar-nav>
      </b-navbar>
    </div>

    <div class="container">
      <b-alert :show="alertShow" :variant="alertVariant"
        ><span v-html="alertMessage"></span>
      </b-alert>

      <PlayersInfo :player1="game.player1.name" :player2="game.player2.name" />

      <div class="row">
        <div class="col mt-2">
          <GameBoard />
        </div>
        <div class="col">
          <Logs ref="logsComponent" />
        </div>
      </div>

      <LaunchGameModal
        ref="launchGameComponent"
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
import GamesResource from "./services/GamesResource.js";

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
      game: {
        sessionId: null,
        player1: {
          name: null,
          symbol: "X",
        },
        player2: {
          name: null,
          symbol: "O",
        },
        winner: "",
      },
      restApi: new GamesResource(),
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
      this.$refs.launchGameComponent.show("NEW_GAME");
    },
    joinGame: function () {
      this.$refs.launchGameComponent.show("JOIN_GAME");
    },
    addLog(text, symbol) {
      this.$refs.logsComponent.addLog({
        symbol: symbol == null ? "X" : "O",
        text: `${this.game.player1.name} has started a new game: ${this.game.sessionId}`,
      });
    },
    onGameStarted: function (gameInfo) {
      this.restApi.newGame(gameInfo.name).then((response) => {
        this.game = response.data;

        this.showAlertMessage(
          `Invite someone to start playing sharing this game code: <span class='h4 mx-1'>${this.game.sessionId}</span>`,
          "info"
        );

        this.addLog(
          `${this.game.player1.name} has started a new game: ${this.game.sessionId}`
        );
      });
    },
    onGameJoined: function (gameInfo) {
      this.restApi
        .joinGame(gameInfo.sessionId, gameInfo.name)
        .then((response) => {
          this.game = response.data;
          console.log(this.game);
          this.showAlertMessage(
            `${this.game.player2.name} has joined the game!`,
            "warning"
          );

          this.addLog(`${this.game.player2.name} has joined the game`, "O");
          this.addLog(
            `${this.game.player1.name} has started a new game: ${this.game.sessionId}`
          );
        });
    },
  },
  mounted() {
    this.showAlertMessage(
      "Tip: Start playing by hosting a game or join an existing one!",
      "dark"
    );
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
