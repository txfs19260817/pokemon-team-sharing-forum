import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
// Global style
import './assets/global.css'

import axios from 'axios'
import i18n from './i18n'
axios.defaults.baseURL = 'http://127.0.0.1:8888/api/v1/';
Vue.prototype.$http = axios;

Vue.config.productionTip = false;

new Vue({
    router,
    i18n,
    render: h => h(App)
}).$mount('#app');
