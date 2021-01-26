<template>
  <div class="contain">
    <CommentList :comments="comments"></CommentList>
    <CommentBox :userInfo="userInfo" :reply-info="replyInfo"  @submit-box="submitBox" :showCancle="showCancle"></CommentBox>

  </div>
</template>
<script>
import CommentList from "./CommentList.vue";
import CommentBox from "./CommentBox.vue";

// vuex中有mapState方法，相当于我们能够使用它的getset方法
//import {mapMutations} from 'vuex';

export default {

  props:['comments'],

  data() {
    return {
      showCancle: false,
      submitting: false,
      value: '',
      userInfo: {
        uid: "uid000001",
        userName: "张三",
        avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"

      },
      replyInfo: {
        uid: "",
        blogUid: "uid000003",
        replyUserUid: "uid000004",
        replyUid: 0,
        avatar: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
      },
     // comments: [
        // {
        //   uid: 'uid000',
        //   userName: "陌溪",
        //   avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        //   content: '我是一级评论',
        //   reply: [
        //     {
        //       uid: 'uid001',
        //       userName: "陌溪",
        //       avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        //       content: '我是二级评论',
        //     },
        //     {
        //       uid: 'uid002',
        //       userName: "陌溪",
        //       avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        //       content: '我是三级评论',
        //     }
        //
        //   ]
        // },
        // {
        //   uid: 'uid004',
        //   userName: "陌溪",
        //   avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
        //   content: '我是一级评论',
        //   reply: [
        //   ]
        // },
      //],
    };
  },
  watch: {},
  computed: {

  },
  components: {
    CommentList,
    CommentBox
  },
  mounted() {
    //this.setCommentList(this.comments);
  },
  methods: {
    //拿到vuex中的写的两个方法
   // ...mapMutations(['SET_COMMENT_LIST', 'INCREMENT']),
    submitBox(e) {
      // var comments = this.$store.state.app.commentList;
      // e.uid = this.$store.state.app.id
      // comments.push(e);
      // this.$store.commit("SET_COMMENT_LIST", comments);
      // this.$store.commit("INCREMENT");
      this.$emit("submit-comment", e)

    }
  },

};
</script>
<style>
.contain{
  width: 600px;
  margin:0 auto;
}
</style>
