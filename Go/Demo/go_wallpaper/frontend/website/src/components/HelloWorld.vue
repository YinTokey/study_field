<template>
  <div class="hello">
    <h1>{{ msg }}</h1>

    <ul class="prev">
      <li v-for="v in dataSource" :key="v.value">
        <img v-bind:src="v.image_url" alt="">
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

ax.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'

export default {

  name: 'HelloWorld',
  props: {
    msg: String,
    dataSource: Array,
    pageInfo :{
      page: 0,
      pageSize: 0,
    },
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
