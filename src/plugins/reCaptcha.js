import Vue from 'vue'
import { VueReCaptcha } from 'vue-recaptcha-v3'

// For more options see below
Vue.use(VueReCaptcha, {
    siteKey: process.env.VUE_APP_RECAPTCHA,
    loaderOptions: {
        useRecaptchaNet: true
    }
});