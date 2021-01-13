import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/components/Home';
import Detail from '@/components/Detail';
import Login from '@/components/Login';
import Comment from '@/components/Comment';
import CommentBox from '@/components/CommentBox';
//import CommentList from "@/components/CommentList";


Vue.use(Router);

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home
        },
        {
            path: '/home',
            name: 'home',
            component: Home
        },
        {
            path: '/detail',
            name: 'detail',
            component: Detail
        },
        {
            path: '/login',
            name: 'login',
            component: Login
        },
        {
            path: '/comment',
            name: 'comment',
            component: Comment
        },
        {
            path: '/commentbox',
            name: 'commentbox',
            component: CommentBox
        },
        // {
        //     path: '/commentlist',
        //     name: 'commentlist',
        //     component: CommentList
        // },
    ]
})