<template>
    <div id="app">
        <el-button type="info" @click="dialogformvisible = true">打开嵌套表单的 Dialog</el-button>
        <teamform :dialogformvisible.sync="dialogformvisible"></teamform>
        <router-view></router-view>
    </div>
</template>

<script>
    import teamform from './components/Submit'
    export default {
        name: 'app',
        components:{
            teamform
        },
        data() {
            return {
                dialogformvisible: false,
            }
        },
        methods: {
            updateFormVisible(b){
                this.dialogformvisible = b;
            }
        }
    }
</script>

<style>

</style>
