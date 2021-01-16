<template>
  <div>
    <img v-bind:src="info.image_url" />
    <h4>{{info.name}}</h4>
    <h4>Author: {{info.author}}</h4>
    <h4>{{info.detail}}</h4>
    <Comment :comments="comments"></Comment>

  </div>
</template>

<script>
import Comment from "./Comment.vue";
import {mapGetters} from 'vuex'
import ax from 'axios'

export default{
  name: 'Detail',
  computed : mapGetters({
    // 获取传递的json信息
    info:'GET_INFO'
  }),
  components: {
    Comment,
  },
  mounted() {
    var url = "http://localhost:8080/api/v1/detail"
    var param = {"ID":this.info.picture_id}

    ax.get(url,{params:param})
        .then(response => {
          //this.info = response.data.bpi
          var info = response.data
          console.log(info)

        })
        .catch(error => {
          console.log(error)
          //this.errored = true
        })
  }
}


</script>