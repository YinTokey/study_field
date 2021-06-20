<template>
  <div id="Home">
    <h1>{{ msg }}</h1>

    <ul class="prev">
      <li v-for="v in dataSource" :key="v.value">
        <img v-bind:href="href" v-bind:src="v.image_url" alt="" v-on:click="detail(v)">
        <h4>{{v.name}}</h4>
        <!--
        <p>{{v.id}}</p>
        -->
      </li>
    </ul>

    <button v-on:click="greet">fetch unsplash pupular</button>


  </div>

</template>

<script>

import ax from 'axios'

// 跨域配置
ax.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'


export default {

  name: 'Home',
  components: {},
  props: {
    msg: String,
    dataSource: Array,
    pageInfo :{
      page: 0,
      pageSize: 0,
    },
    href: {
      type:String,
      required: true
    }
  },

  methods: {
    greet: function (event) {
      // `this` 在方法里指当前 Vue 实例

      console.log(event)

      if (this.pageInfo === undefined) {
        this.pageInfo = {page:1,pageSize:10}
      }
      if (this.dataSource === undefined) {
        this.dataSource = new Array()
      }

      var url2 = "http://localhost:8080/api/v1/papular"
      var param = {"page":this.pageInfo.page,"pageSize":this.pageInfo.pageSize}

      console.log(param)

      ax.get(url2,{params:param})
          .then(response => {
            //this.info = response.data.bpi
            var info = response.data

            this.dataSource = this.dataSource.concat(info)
            console.log(info)
            this.pageInfo.page += 1
          })
          .catch(error => {
            console.log(error)
            //this.errored = true
          })
    },


    detail: function (v) {
      this.$store.commit('SET_INFO',v)
      this.$router.push({path: '/detail'}) // 路由跳转到详情页

    }

  }

}


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
