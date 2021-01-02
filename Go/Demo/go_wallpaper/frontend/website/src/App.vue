<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <Home msg="home page" @messageData="getData"/>
  </div>
</template>

<script>
import Home from './components/Home.vue'
//import Detail from './components/Detail.vue'

export default {
  name: 'app',
  components: {
    Home
  },

  methods : {
    // 接收来自 Home发送的消息
    getData(val) {
      console.log(val)
      this.$router.push(val)
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
