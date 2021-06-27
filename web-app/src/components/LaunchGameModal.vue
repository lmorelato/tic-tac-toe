<template>
  <div>
    <b-modal
      id="launchModal"
      v-model="modalShow"
      @hidden="resetModal"
      @ok="handleOk"
    >
      <template #modal-header>
        <h5>{{ action == "NEW_GAME" ? "New Game" : "Join Game" }}</h5>
      </template>

      <div v-if="action == 'JOIN_GAME'" class="row my-2">
        <div class="col-3 mt-2">Game Id*</div>
        <div class="col">
          <b-form-input
            v-model="sessionId"
            placeholder="Enter the invitation code"
          ></b-form-input>
        </div>
      </div>
      <div class="row">
        <div class="col-3 mt-2">Player Name*</div>
        <div class="col">
          <b-form-input
            v-model="name"
            placeholder="Enter your name"
          ></b-form-input>
        </div>
      </div>

      <b-alert class="mt-4" :show="invalidData" variant="danger"
        >*Please, review the required field(s)
      </b-alert>
    </b-modal>
  </div>
</template>

<script>
export default {
  name: "LaunchGameModal",
  data: function () {
    return {
      name: "",
      sessionId: "",
      action: "",
      modalShow: false,
      invalidData: false,
    };
  },
  props: {},
  methods: {
    show(action) {
      this.action = action;
      this.modalShow = true;
    },
    checkFormValidity() {
      if (this.action == "NEW_GAME") {
        return this.name.length > 0;
      }

      return this.name.length > 0 && this.sessionId.length > 0;
    },
    resetModal() {
      this.name = "";
      this.sessionId = "";
      this.action = "";
      this.modalShow = false;
      this.invalidData = false;
    },
    hiddenModal() {
      this.resetModal();
    },
    handleOk(bvModalEvt) {
      bvModalEvt.preventDefault();
      this.handleSubmit();
    },
    handleSubmit() {
      if (!this.checkFormValidity()) {
        this.invalidData = true;
        return;
      }

      this.$nextTick(() => {
        if (this.action == "NEW_GAME")
          this.$emit("gameStarted", { name: this.name });
        else
          this.$emit("gameJoined", {
            name: this.name,
            sessionId: this.sessionId,
          });

        this.$bvModal.hide("launchModal");
      });
    },
  },
};
</script>

<style scoped>
</style>
