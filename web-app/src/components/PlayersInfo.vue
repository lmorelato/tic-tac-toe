<template>
  <div class="d-flex justify-content-center align-items-center my-5">
    <span class="h1 player1 text-truncate">{{ player1 || "Player 1 ?" }}</span>
    <span class="h2 px-5">vs</span>
    <span class="h1 player2 text-truncate">{{ player2 || "Player 2 ?" }}</span>
  </div>
</template>

<script>
export default {
  name: "PlayersInfo",
  props: {
    player1: { type: String, default: null },
    player2: { type: String, default: null },
  },
};
</script>

<style scoped>
.player1 {
  color: #4169e1;
  text-transform: uppercase;
  text-align: right;
  width: 270px;
}
.player2 {
  color: #ff4500;
  text-transform: uppercase;
  text-align: left;
  width: 270px;
}
</style>
