import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

// 首先声明一个需要全局维护的状态 state

const store = new Vuex.Store({
    state: { // 共同维护的一个状态，state里面可以是很多个全局状态
        info: {}
    },
    mutations: { // 处理数据的唯一途径，state的改变或赋值只能在这里
        "SET_INFO": function(state, info) {
            state.info = info
            console.log('保存', state.info)
        }
    },
    getters: { // 获取数据并渲染
        "GET_INFO": function(state) {
            console.log('获取', state.info)
            return state.info
        }
    },
    actions: { // 数据的异步操作
        "SET_INFO": function(state, info) {
            console.log('获取', state.info)
            store.commit("SET_INFO", info)
        }
    }
})
export default store // 导出store并在 main.js中引用注册。

