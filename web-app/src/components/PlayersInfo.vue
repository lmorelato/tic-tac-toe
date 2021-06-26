<template>
  <div class="d-flex justify-content-center align-items-center mb-5">
    <span class="h2 player1">{{ player1 }}</span>
    <span class="h3 px-5">vs</span>
    <span class="h2 player2">{{ player2 }}</span>
  </div>
</template>

<script>
export default {
  name: "PlayersInfo",
  props: {
    player1: {type: String, default: "Player 1"},
    player2: {type: String, default: "Player 2"},
  },
};
</script>

<style scoped>
.player1 {
  color: #4169e1;
}
.player2 {
  color: #ff4500;
}
</style>
