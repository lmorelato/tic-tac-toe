<template>
  <div class="d-flex justify-content-center">
    <div class="row board">
      <div
        v-for="(item, index) in board"
        :key="index"
        class="
          col-4
          board-cell
          d-flex
          align-items-center
          justify-content-center
        "
        @click="play(index)"
      >
        <b-icon v-if="item == 'X'" class="symbol-x" icon="x"></b-icon>
        <b-icon v-if="item == 'O'" class="symbol-o" icon="circle"></b-icon>
        <!-- animation="throb" -->
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue";

export default {
  name: "GameBoard",
  data: function () {
    return {
      board: new Array(),
    };
  },
  props: {
    msg: String,
  },
  methods: {
    play: function (index) {
      Vue.set(this.board, index, index % 2 == 0 ? "O" : "X");
    },
  },
  beforeMount() {
    const boardSize = 9;
    for (let index = 0; index < boardSize; index++) {
      this.board.push("");
    }
  },
};
</script>

<style scoped>
.board {
  width: 300px;
}

.board-cell {
  height: 100px;
  border-bottom: 4px solid #000;
  border-right: 4px solid #000;
}

.board-cell:nth-child(n + 7) {
  border-bottom: 0;
}

.board-cell:nth-child(3n) {
  border-right: 0;
}

.symbol-x {
  width: 90px;
  height: 90px;
  color: #4169E1;
}

.symbol-o {
  width: 52px;
  height: 52px;
  color: #FF4500;
}
</style>
