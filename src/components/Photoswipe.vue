<template>
    <div id="photoswipe">
        <div class="pswp" ref="pswb" tabindex="-1" role="dialog" aria-hidden="true">
            <div class="pswp__bg"></div>
            <div class="pswp__scroll-wrap">
                <div class="pswp__container">
                    <div class="pswp__item"></div>
                    <div class="pswp__item"></div>
                    <div class="pswp__item"></div>
                </div>
                <div class="pswp__ui pswp__ui--hidden">
                    <div class="pswp__top-bar">
                        <div class="pswp__counter"></div>
                        <button class="pswp__button pswp__button--close" title="Close (Esc)"></button>
                        <div class="pswp__preloader">
                            <div class="pswp__preloader__icn">
                                <div class="pswp__preloader__cut">
                                    <div class="pswp__preloader__donut"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="pswp__share-modal pswp__share-modal--hidden pswp__single-tap">
                        <div class="pswp__share-tooltip"></div>
                    </div>
                    <button class="pswp__button pswp__button--arrow--left" title="Previous (arrow left)"></button>
                    <button class="pswp__button pswp__button--arrow--right" title="Next (arrow right)"></button>
                    <div class="pswp__caption">
                        <div class="pswp__caption__center"></div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>


<script>
    import PhotoSwipe from 'photoswipe/dist/photoswipe'
    import UI from 'photoswipe/dist/photoswipe-ui-default'
    import 'photoswipe/dist/photoswipe.css'
    import 'photoswipe/dist/default-skin/default-skin.css'

    export default {
        name: "photoswipe",
        props: {
            items: {
                type: Array,
                required: true
            }
        },
        data() {
            return {
                teams: [],
            }
        },
        watch: {
            items() {
                this.teams = this.items;
            }
        },
        methods: {
            imagePreview(currentIndex) {
                const options = {
                    index: currentIndex
                };

                const gallery = new PhotoSwipe(this.$refs.pswb, UI, this.items, options);
                gallery.init()
            }
        }
    }
</script>

<style scoped>

</style>