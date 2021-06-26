<template>
  <div id="app">
    <div class="mb-4">
      <b-navbar type="dark" variant="dark">
        <b-navbar-nav>
          <b-navbar-brand class="mx-4">Tic-Tac-Toe</b-navbar-brand>
          <b-nav-item href="#">New Game</b-nav-item>
          <b-nav-item href="#">Join Game</b-nav-item>
          <b-nav-item href="#">About</b-nav-item>
        </b-navbar-nav>
      </b-navbar>
    </div>

    <div class="container">
      <b-alert show variant="info"
        >Share your session Id with someone to start playing:
        <span class="h4 mx-1">ase3</span>
      </b-alert>
      <PlayersInfo :player1="player1" :player2="player2" />
      <GameBoard />
    </div>
  </div>
</template>

<script>
import GameBoard from "./components/GameBoard.vue";
import PlayersInfo from "./components/PlayersInfo.vue";

export default {
  name: "App",
  components: {
    GameBoard,
    PlayersInfo,
  },
  data: function () {
    return {
      player1: "leandro",
      player2: "robert",
    };
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
