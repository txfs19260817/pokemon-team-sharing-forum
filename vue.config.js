module.exports = {
    chainWebpack: config => {
        config.when(process.env.NODE_ENV === 'production', config => {
            config
                .entry('app')
                .clear()
                .add('./src/main-prod.js');
            config.devServer.disableHostCheck(true);

            config.plugin('html').tap(args => {
                args[0].isProd = true;
                return args
            });

            config.set('externals', {
                vue: 'Vue',
                // 'vue-router': 'VueRouter',
                axios: 'axios',
                // koffing: 'Koffing',
                photoswipe: 'PhotoSwipe',
                html2canvas: 'html2canvas',
                echarts: 'echarts'
            });
        });

        config.when(process.env.NODE_ENV === 'development', config => {
            config
                .entry('app')
                .clear()
                .add('./src/main.js');
            config.devServer.disableHostCheck(true);

            config.plugin('html').tap(args => {
                args[0].isProd = false;
                return args
            });
        });
    },
    publicPath: './',
    css: {
        extract: false
    },

    pluginOptions: {
        i18n: {
            locale: 'zh-hans',
            fallbackLocale: 'en',
            localeDir: 'locales',
            enableInSFC: false
        }
    }
};
