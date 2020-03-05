import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/Home';
import Team from "../components/pages/Team";
import Format from "../components/pages/Format";
import Pokemon from "../components/pages/Pokemon";
import About from "../components/pages/About";
import Board from "../components/pages/Board";
import Stat from "../components/pages/Stat";

Vue.use(VueRouter);
const routes = [
    { path: '/', redirect: '/home' },
    { path: '/home&*', redirect: '/home' },//open an image and refresh
    { path: '/home', component: Home },// lazy load () => import('../components/Home')
    { path: '/team/:id', component: Team },
    { path: '/formats/:format', component: Format },
    { path: '/pokemon/:pokemon', component: Pokemon },
    { path: '/board', component: Board },
    { path: '/stat', component: Stat },
    { path: '/about', component: About }
];

const router = new VueRouter({
    routes
});

export default router