import Vue from "vue";
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import axios from 'axios';
import App from "./App.vue";

Vue.use(BootstrapVue);
Vue.use(BootstrapVueIcons);
Vue.config.productionTip = false;
axios.defaults.baseURL = "http://localhost:8889/api";

new Vue({
  render: (h) => h(App),
}).$mount("#app");
