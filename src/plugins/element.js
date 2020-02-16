import Vue from 'vue'
import {Alert, Upload, Dialog, Option, Select, Cascader, Button, Form, FormItem, Input, Message} from 'element-ui'

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