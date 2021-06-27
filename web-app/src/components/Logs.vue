<template>
  <div class="d-flex justify-content-center">
    <div class="logs">
      <div class="px-3 pt-1"> /* TIC-TAC-TOE GAME LOG ********************************/ </div>
      <div class="line px-3 pt-1 symbol-x">game log</div>
      <div class="line px-3 pt-1 symbol-o">game log</div>
      <div class="line px-3 pt-1">game log</div>
      <div class="line px-3 pt-1">game log</div>
      <div class="line px-3 pt-1">game log</div>
      <div class="line px-3 pt-1">game log</div>
      <div class="line px-3 pt-1">game log</div>
      <div class="line px-3 pt-1">game log</div>
      <div class="line px-3 pt-1">game log</div>
      <div class="line px-3 pt-1">game log</div>

    </div>
  </div>
</template>

<script>
export default {
  name: "Logs",
  props: {
    player1: { type: String, default: null },
    player2: { type: String, default: null },
  },
};
</script>

<style scoped>
.logs {
  border: 1px solid;
  width: 600px;
  height: 360px;
  max-height: 400px;
  overflow-y: auto;
  text-align: left;
  font-family: "Courier New", Courier, monospace;
  font-weight: 600;
  background-color: #F8F8F8;
}

.line::before {
  content: "";
}

.line::after {
  content: "...";
}

.symbol-x {
  color: #4169e1;
}
.symbol-o {
  color: #ff4500;
}
</style>
