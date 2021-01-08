<template>
  <div>
    <img v-bind:src="info.image_url" />
    <h4>{{info.name}}</h4>
    <h4>Author: {{info.author}}</h4>
    <h4>detail: {{info.detail}}</h4>


  </div>
</template>

<script>

import {mapGetters} from 'vuex'

export default{
  name: 'Detail',
  computed : mapGetters({
    // 获取传递的json信息
    info:'GET_INFO'
  })
}


</script>