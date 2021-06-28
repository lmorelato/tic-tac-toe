<template>
  <div id="app" class="mb-5">
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
          <GameBoard
            ref="boardComponent"
            :playerSymbol="playerSymbol"
            :gameOver="this.game.winner.symbol != ''"
            :winningPath="this.game.winningPath"
            :nextSymbol="this.game.nextSymbol"
            @addBoardLog="onAddBoardLog"
            @boardClicked="onBoardClicked"
          />
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
  components: { GameBoard, PlayersInfo, LaunchGameModal, Logs },
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
        moves: [],
        winner: {
          name: null,
          symbol: null,
        },
        winningPath: [],
        nextSymbol: "",
      },
      restApi: new GamesResource(),
      modalShow: false,
      modalAction: "",
      alertShow: false,
      alertMessage: "",
      alertVariant: "",
      playerSymbol: "",
      polling: null,
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
    addLog(text) {
      this.$refs.logsComponent.addLog(text);
    },
    onGameStarted: function (gameInfo) {
      this.reset();
      this.playerSymbol = "X";

      this.restApi.newGame(gameInfo.name).then((response) => {
        this.game = response.data;
        this.initPooling();

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
      this.reset();
      this.playerSymbol = "O";

      this.restApi.getGame(gameInfo.sessionId).then((response) => {
        if (!response.data) {
          this.showAlertMessage(
            `Game not found.`,
            "danger"
          );
          return;
        }
        if (response.data.player2.name) {
          this.showAlertMessage(
            `This game cannot be joined, it has 2 players already.`,
            "danger"
          );
          return;
        }

        this.restApi
          .joinGame(gameInfo.sessionId, gameInfo.name)
          .then((response) => {
            this.game = response.data;

            this.showAlertMessage(
              `${this.game.player2.name} has joined the game!`,
              "warning"
            );
            this.addLog(
              `${this.game.player1.name} has started a new game: ${this.game.sessionId}`
            );
            this.addLog(`${this.game.player2.name} has joined the game`, "O");
          });

        this.initPooling();
      });
    },
    onBoardClicked: function (index) {
      this.restApi.play(this.game.sessionId, index, this.playerSymbol);
    },
    onAddBoardLog: function (log) {
      this.addLog(`Position ${log.index + 1}, Symbol: ${log.symbol}`);
    },
    initPooling() {
      if (this.polling) return;

      this.polling = setInterval(() => {
        if (this.game.sessionId) {
          this.restApi.getGame(this.game.sessionId).then((response) => {
            var noPlayer2 = this.game.player2.name == "";
            var noWinner = this.game.winner.symbol == "";

            this.game = response.data;
            this.checkTurn(noPlayer2, noWinner);
            this.checkPlayer2(noPlayer2);
            this.fillBoard();
            this.checkWinner(noWinner);
          });
        }
      }, 500);
    },
    fillBoard() {
      // update the board with the other player moves
      for (let index = 0; index < this.game.moves.length; index++) {
        if (
          this.game.moves[index] != "" &&
          this.game.moves[index] != this.playerSymbol
        ) {
          let symbol = this.playerSymbol == "X" ? "O" : "X";
          this.$refs.boardComponent.play(index, symbol, false, true);
        }
      }
    },
    checkPlayer2(noPlayer2) {
      // check if plyear 2 joined the game
      let hasPlayer2 = noPlayer2 && this.game.player2.name != "";
      if (hasPlayer2) {
        this.showAlertMessage(
          `${this.game.player2.name} has joined the game!`,
          "warning"
        );
        this.addLog(`${this.game.player2.name} has joined the game`, "O");
      }
    },
    checkWinner(noWinner) {
      if (noWinner && this.game.winner.name != "") {
        if (this.game.winner.name == "NO_WINNER") {
          this.showAlertMessage(`Oh no! It is a tie!`, "info");
          this.addLog(`It is a tie.`);
        } else {
          if (this.game.winner.symbol == this.playerSymbol) {
            this.showAlertMessage(`You Win!`, "success");
          } else {
            this.showAlertMessage(`You Lose!`, "danger");
          }
          this.addLog(`${this.game.winner.name} won the game!`);
        }

        clearInterval(this.polling);
        this.polling = null;
      }
    },
    checkTurn(noPlayer2, noWinner) {
      if (!noPlayer2 && noWinner) {
        if (this.game.nextSymbol == this.playerSymbol) {
          this.showAlertMessage(`It is your turn, let's play!`, "info");
        } else {
          this.showAlertMessage(`Waiting for the other player...`, "warning");
        }
      }
    },
    reset() {
      this.$refs.boardComponent.reset();
      this.$refs.logsComponent.reset();
    },
  },
  created() {
    this.initPooling();
  },
  mounted() {
    this.showAlertMessage(
      "Tip: Start playing by hosting a game or join an existing one!",
      "dark"
    );
  },
  beforeDestroy() {
    clearInterval(this.polling);
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
