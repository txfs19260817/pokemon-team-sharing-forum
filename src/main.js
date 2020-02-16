import Vue from 'vue'
import App from './App.vue'
import './plugins/element.js'
// Global style
import './assets/global.css'

import axios from 'axios'
axios.defaults.baseURL = 'http://127.0.0.1:8888/api/v1/';// 配置请求根路径
Vue.prototype.$http = axios;

Vue.config.productionTip = false;

new Vue({
    render: h => h(App)
}).$mount('#app');
