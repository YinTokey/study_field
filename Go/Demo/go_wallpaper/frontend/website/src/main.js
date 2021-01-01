import Vue from 'vue'
import App from './App.vue'
import router from "./router";

Vue.config.productionTip = false

// 路由跳转
Vue.prototype.$goRoute = function (index) {
  console.log(index)
  this.$router.push(index)
}

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})

