<template>
  <div class="d-flex justify-content-center">
    <div class="logs">
      <div class="px-3 py-3 title">
        /* GAME LOG ******************************************/
      </div>
      <div
        v-for="(item, index) in logs"
        :key="index"
        class="line px-3 pt- ml-1"
      >
        {{ logs.length - index }}- {{ item }}
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
    reset() {
      this.logs = new Array();
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

.title {
  font-size: 105%;
}
</style>
