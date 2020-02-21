import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/Home';
import Team from "../components/Team";
import Format from "../components/Format";

Vue.use(VueRouter);
const routes = [
    { path: '/', redirect: '/home' },
    { path: '/home&*', redirect: '/home' },//open an image and refresh
    { path: '/home', component: Home },// lazy load () => import('../components/Home')
    { path: '/team/:id', component: Team },
    { path: '/formats/:format', component: Format }
];

const router = new VueRouter({
    routes
});

export default router