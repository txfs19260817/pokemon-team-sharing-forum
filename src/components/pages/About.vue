<template>
    <base-layout>
        <template v-slot:default>
            欢迎使用本网站，您可以在此分享您使用过的队伍，或者查看其他玩家分享的队伍。
            <h3>网站特色</h3>
            <h4>分享队伍</h4>
            您可以在本网站上与全世界的玩家们分享您喜爱的队伍。
            您不仅可以选择分享游戏软件中导出的租借队伍ID界面（下面简称租借ID），
            还可以上传<a href="https://play.pokemonshowdown.com/">Pokemon Showdown</a>格式的队伍文本（下面简称Showdown文本）。
            <strong>我们强烈建议您在发布队伍的时候同时上传二者。</strong>
            如果您选择分享Showdown文本，本网站还会从文本导出一张租借ID样式的队伍预览图，这张图将代替租借ID显示在我们的网站上。
            您的队伍在提交完毕后不会立即显示在本网站上，请您耐心等待审核。
            <h4>获取队伍</h4>
            网页上会展示用户上传的队伍预览图（如果上传用户没有提交游戏软件中导出的租借队伍ID，则以根据Showdown文本导出的队伍预览图作为代替）。
            点击预览图下方的<b>S</b>按钮，可以查看并复制队伍的Showdown文本。
            <h3>与我们联系</h3>
            我们开通了推特<a href="https://twitter.com/poke_share">@poke_share</a>
            和微博<a href="https://www.weibo.com/u/7403524801">@Poke_Share</a>。
            如果您在使用本网站中遇到任何问题，或愿意作出贡献，
            热烈欢迎到<a href="https://github.com/txfs19260817/pokemon-team-sharing-forum">Github</a>上提交issue或pull request，
            或者发送邮件至
            <a href="mailto:z.jiang@pokeshare.monster?subject=Pokemon%20team%20sharing%20forum%20">z.jiang@pokeshare.monster</a>进行反馈。
            如果您认为网站上所发表的队伍侵犯了您的创作权，请发送邮件并附上相关证明。
            <h3>License</h3>
            本网站使用的资源（图标等）版权归 © Nintendo/Creatures Inc./GAME FREAK Inc. 所有
        </template>
    </base-layout>
</template>

<script>
    import BaseLayout from "../layouts/BaseLayout";

    export default {
        name: "About",
        components: {
            BaseLayout
        }
    }
</script>

<style scoped>

</style>