import Vue from 'vue'
import {
    Tag,
    Tabs,
    TabPane,
    Drawer,
    Pagination,
    Container,
    Header,
    Main,
    Footer,
    Alert,
    Upload,
    Dialog,
    Option,
    Select,
    Cascader,
    Button,
    Form,
    FormItem,
    Input,
    Message
} from 'element-ui'

Vue.use(Tag);
Vue.use(Tabs);
Vue.use(TabPane);
Vue.use(Drawer);
Vue.use(Pagination);
Vue.use(Container);
Vue.use(Header);
Vue.use(Main);
Vue.use(Footer);
Vue.use(Alert);
Vue.use(Upload);
Vue.use(Dialog);
Vue.use(Option);
Vue.use(Select);
Vue.use(Cascader);
Vue.use(Button);
Vue.use(Form);
Vue.use(FormItem);
Vue.use(Input);
Vue.prototype.$message = Message;
// Vue.prototype.$loading = Loading.service;