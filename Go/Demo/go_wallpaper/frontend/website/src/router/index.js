import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/components/Home';
import Detail from '@/components/Detail';

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/home',
            component: Home
        },
        {
            path: '/detail',
            component: Detail
        }
    ]
});