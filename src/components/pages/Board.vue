<template>
    <base-layout :loading="false" :fail="false">
        <!--     Twitter: https://publish.twitter.com/# -->
        <h3>Twitter {{$t('navigation.board')}}</h3>
        <div class="twitter-area">
            <a class="twitter-timeline" data-width="800" data-height="1000" data-dnt="true" data-theme="dark"
               href="https://twitter.com/poke_share?ref_src=twsrc%5Etfw">Tweets by PokéShare!</a>
        </div>
    </base-layout>
</template>

<script>
    import BaseLayout from "../layouts/BaseLayout";

    export default {
        name: "Board",
        components: {
            BaseLayout
        },
        beforeCreate() {
            const s = document.createElement('script');
            s.type = 'text/javascript';
            s.async = true;
            s.src = 'https://platform.twitter.com/widgets.js';
            document.body.appendChild(s);
        }
    }
</script>

<style scoped>
    .twitter-area {
        text-align: center;
    }
</style>