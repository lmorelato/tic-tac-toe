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
        :style="{
          cursor:
            gameOver || playerSymbol == '' || board[index] != ''
              ? 'default'
              : 'pointer',
        }"
        @click="play(index, playerSymbol, true, false)"
      >
        <b-icon
          v-if="item == 'X'"
          class="symbol-x"
          icon="x"
          :animation="checkWinningPath(index)"
        ></b-icon>
        <b-icon
          v-if="item == 'O'"
          class="symbol-o"
          icon="circle"
          :animation="checkWinningPath(index)"
        ></b-icon>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "GameBoard",
  data: function () {
    return {
      board: null,
    };
  },
  props: {
    playerSymbol: String,
    gameOver: Boolean,
    winningPath: Array,
    nextSymbol: String,
  },
  methods: {
    play: function (index, playerSymbol, emit, forceFill) {
      if (
        !this.gameOver &&
        this.board[index] == "" &&
        (forceFill || playerSymbol == this.nextSymbol)
      ) {
        this.$set(this.board, index, playerSymbol);
        this.$emit("logAdded", { index: index, symbol: playerSymbol });

        if (emit) {
          this.$emit("boardClicked", index);
        }
      }
    },
    checkWinningPath(index) {
      return this.winningPath[0] != this.winningPath[1] &&
        this.winningPath.includes(index)
        ? "fade"
        : "none";
    },
    reset() {
      const boardSize = 9;
      this.board = new Array();

      for (let index = 0; index < boardSize; index++) {
        this.board.push("");
      }
    },
  },
  beforeMount() {
    this.reset();
  },
};
</script>

<style scoped>
.board {
  width: 330px;
}

.board-cell {
  height: 110px;
  border-bottom: 5px solid #000;
  border-right: 5px solid #000;
}

.board-cell:nth-child(n + 7) {
  border-bottom: 0;
}

.board-cell:nth-child(3n) {
  border-right: 0;
}

.symbol-x {
  width: 110px;
  height: 110px;
  color: #4169e1;
}

.symbol-o {
  width: 60px;
  height: 60px;
  color: #ff4500;
}
</style>
