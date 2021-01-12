import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/components/Home';
import Detail from '@/components/Detail';
import Login from '@/components/Login';

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
        }
    ]
})