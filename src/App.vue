<template>
  <div id="app">
    <el-button type="info" @click="updateFormVisible">提交队伍</el-button>
    <teamform ref="formRef" :dialogformvisible.sync="dialogformvisible"></teamform>
  </div>
</template>

<script>
    import teamform from './components/Submit'

    export default {
        name: 'app',
        components: {
            teamform
        },
        data() {
            return {
                dialogformvisible: false,
            }
        },
        methods: {
            updateFormVisible() {
                this.dialogformvisible = true;
                setTimeout(this.$refs.formRef.resetForm, 100);
            }
        }
    }
</script>

<style>

</style>
