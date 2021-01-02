import Vue from 'vue'
import App from "@/App"
import router from "./router";
import store from './store'

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,  // 全局使用vuex
  render: h => h(App)
})
