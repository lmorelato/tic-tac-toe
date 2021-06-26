<template>
  <div id="app" class="container-fluid">
    <PlayersInfo :player1="player1" :player2="player2" />
    <GameBoard />
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
  margin-top: 60px;
}
</style>
