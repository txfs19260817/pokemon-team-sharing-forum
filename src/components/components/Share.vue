<template>
    <span>
        <el-button v-if="destination.includes('twitter')"
                   type="primary" icon="el-icon-social-twitter" circle
                   class="button" data-sharer="twitter"
                   :data-title="shareTitle"
                   data-hashtags="PokeShare"
                   :data-url="shareUrl"
                   data-via="poke_share">
        </el-button>
        <el-button v-if="destination.includes('weibo')"
                   type="danger" icon="el-icon-social-weibo" circle
                   class="button" data-sharer="weibo"
                   :data-title="shareTitle"
                   :data-image="team.rentalImgUrl"
                   :data-url="shareUrl"
                   data-ralateuid="7403524801">
        </el-button>
    </span>
</template>

<script>
    export default {
        name: "Share",
        props: {
            destination: {
                type: Array,
                default: () => ['twitter', 'weibo']
            },
            team: {
                type: Object,
                required: true
            }
        },
        computed: {
            shareUrl() {
                return window.location.href
            },
            shareTitle() {
                const title = this.team.title;
                const format = this.team.format;
                const author = this.team.author;
                const pokemon1 = this.team.pokemon1;
                const pokemon2 = this.team.pokemon2;
                const pokemon3 = this.team.pokemon3;
                const pokemon4 = this.team.pokemon4;
                const pokemon5 = this.team.pokemon5;
                const pokemon6 = this.team.pokemon6;
                return `Checkout a team found on PokeShare!
    [${format}]${title} by ${author}
    Pokemon: ${pokemon1}, ${pokemon2}, ${pokemon3}, ${pokemon4}, ${pokemon5}, ${pokemon6}`
            }
        },
        mounted() {
            window.Sharer.init()
        }
    }
</script>

<style scoped>

</style>