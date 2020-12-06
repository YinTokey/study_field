<template>
  <div class="hello">
    <h1>{{ msg }}</h1>

    <button v-on:click="greet">fetch 500 px pupular</button>

    <ul>
      <li v-for="v in dataSource" :key="v.value">
        <img v-bind:src="v.image_url" alt="">
        <h4>{{v.name}}</h4>
        <!--
        <p>{{v.id}}</p>
        -->
      </li>
    </ul>

  </div>


</template>

<script>

import ax from 'axios'

ax.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'


export default {
  name: 'HelloWorld',
  props: {
    msg: String,
    dataSource: [],
  },

  methods: {
    greet: function (event) {
      // `this` 在方法里指当前 Vue 实例
      console.log(event)
      // var url1 = "https://api.500px.com/v1/photos?feature=popular"
      //http://localhost:8080/api/v1/papular
      var url2 = "http://localhost:8080/api/v1/papular"
      //var params = {}
      ax.get(url2)
          .then(response => {
            //this.info = response.data.bpi
            var info = response.data
            this.dataSource = info
            console.log(info)
          })
          .catch(error => {
            console.log(error)
            //this.errored = true
          })




    }
  }

}



</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
