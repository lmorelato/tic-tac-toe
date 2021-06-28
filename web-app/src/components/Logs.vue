<template>
  <div class="d-flex justify-content-center">
    <div class="logs">
      <div class="px-3 py-3 title"> /* GAME LOGS *****************************************/</div>
      <div
        v-for="(item, index) in logs"
        :key="index"
        class="line px-3 pt- ml-1 symbol-x"
        :class="{
          'symbol-x': item.symbol == 'X',
          'symbol-o': item.symbol == 'O',
        }"
      >
        {{ logs.length - index }}- {{ item.text }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Logs",
  data: function () {
    return {
      logs: new Array(),
    };
  },
  methods: {
    addLog(log) {
      this.logs.unshift(log);
    },
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
  background-color: #fcfcfc;
}

.symbol-x {
  color: #4169e1;
}
.symbol-o {
  color: #ff4500;
}

.title {
  font-size: 105%;
}
</style>
