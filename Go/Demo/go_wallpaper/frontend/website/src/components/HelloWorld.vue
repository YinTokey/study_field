<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      For a guide and recipes on how to configure / customize this project,<br>
      check out the
      <a href="https://cli.vuejs.org" target="_blank" rel="noopener">vue-cli documentation</a>.
    </p>

    <button v-on:click="greet">fetch 500 px pupular</button>

  </div>


</template>

<script>

import ax from 'axios'
import qs from 'qs';

ax.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'


export default {
  name: 'HelloWorld',
  props: {
    msg: String,
    info: String,
  },

  methods: {
    greet: function (event) {
      // `this` 在方法里指当前 Vue 实例
      console.log(event)
     // var url1 = "https://api.500px.com/v1/photos?feature=popular"
      var url2 = "http://127.0.0.1:8080/api/v1/papular"
      var params = {}
      ax.get(url2, qs.stringify(params))
          .then(response => {
            //this.info = response.data.bpi
            console.log(response)
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
