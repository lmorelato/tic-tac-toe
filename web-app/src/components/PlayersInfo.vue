<template>
  <div class="mt-4 mb-5">
    <div class="d-flex justify-content-center align-items-center">
      <span class="h1 player1 text-truncate">{{ player1 || "???" }}</span>
      <span class="h2 px-5">vs</span>
      <span class="h1 player2 text-truncate">{{ player2 || "???" }}</span>
    </div>
    <div class="d-flex justify-content-center">
      <hr style="width: 630px; margin:0;" />
    </div>
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
  width: 250px;
}
.player2 {
  color: #ff4500;
  text-transform: uppercase;
  text-align: left;
  width: 250px;
}
</style>
