<template>
    <div id="app">
        <el-container class="home-container">
            <header>
                <img id="logo" src="../public/logo.png" alt="Pokemon Team Sharing Forum"/>
                <a id="submit" href="#" @click="updateFormVisible">Submit</a>
                <teamform ref="formRef" :dialogformvisible.sync="dialogformvisible"></teamform>
            </header>
            <el-main>Main</el-main>
        </el-container>
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

<style scoped>
    @import url("https://fonts.googleapis.com/css?family=Montserrat");
    .home-container {
        height: 100%;
        font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    }

    #logo {
        height: 60%;
    }

    header {
        position: fixed;
        width: 100%;
        height: 12em;
        display: -webkit-box;  /* OLD - iOS 6-, Safari 3.1-6, BB7 */
        display: -ms-flexbox;  /* TWEENER - IE 10 */
        display: -webkit-flex; /* NEW - Safari 6.1+. iOS 7.1+, BB10 */
        display: flex;         /* NEW, Spec - Firefox, Chrome, Opera */
        -webkit-box-pack: justify;
        -webkit-justify-content: space-between;
        -ms-flex-pack: justify;
        justify-content: space-between;
    }
    /* Create angled background with 'before' pseudo-element */
    header::before {
        content: "";
        display: block;
        align-items: end;
        position: absolute;
        left: 0;
        bottom: 6em;
        width: 100%;
        height: 20em;
        z-index: -1;
        transform: skewY(-2.0deg);
        background: linear-gradient(rgba(0, 0, 0, .6), rgba(0, 0, 0, .6)),
        url(../public/header.png) no-repeat center,
        linear-gradient(#4e4376, #2b5876);
        background-size: cover;
        border-bottom: .2em solid #fff;
    }

    #submit {
        position: absolute;
        right: 5%;
        top: 10%;
        -webkit-transform: translate(-@left, -@top);
        transform: translate(-@left, -@top);
        /*color: #cecd24;*/
        color: #F2F6FC;
        text-decoration: none;
        font-size: 2em;
        display: inline-block;
        font-family: Montserrat, sans-serif;
        text-transform: uppercase;
        padding: 0.5em 2em;
        border: 2px solid #F2F6FC;
        -webkit-transition: 0.02s 0.2s cubic-bezier(0.1, 0, 0.1, 1);
        transition: 0.02s 0.2s cubic-bezier(0.1, 0, 0.1, 1);
    }
    #submit::before {
        content: "";
        display: inline-block;
        position: absolute;
        top: 0;
        left: 0;
        right: 100%;
        bottom: 0;
        background: #F2F6FC;
        -webkit-transition: 0.3s 0.2s cubic-bezier(0.1, 0, 0.1, 1), left 0.3s cubic-bezier(0.1, 0, 0.1, 1);
        transition: 0.3s 0.2s cubic-bezier(0.1, 0, 0.1, 1), left 0.3s cubic-bezier(0.1, 0, 0.1, 1);
        z-index: -1;
    }
    #submit::after {
        content: "";
        display: inline-block;
        background-image: url(../public/arrow.png);
        position: absolute;
        top: 0;
        left: calc(100% - 3em);
        right: 3em;
        bottom: 0;
        background-size: 1.5em;
        background-repeat: no-repeat;
        background-position: center;
        -webkit-transition: right 0.3s cubic-bezier(0.1, 0, 0.1, 1);
        transition: right 0.3s cubic-bezier(0.1, 0, 0.1, 1);
    }
    #submit:hover {
        padding: 0.5em 3.5em 0.5em 0.5em;
    }
    #submit:hover::before {
        left: calc(100% - 3em);
        right: 0;
        -webkit-transition: 0.3s cubic-bezier(0.1, 0, 0.1, 1), left 0.3s 0.2s cubic-bezier(0.1, 0, 0.1, 1);
        transition: 0.3s cubic-bezier(0.1, 0, 0.1, 1), left 0.3s 0.2s cubic-bezier(0.1, 0, 0.1, 1);
    }
    #submit:hover::after {
        right: 0;
        -webkit-transition: right 0.3s 0.2s cubic-bezier(0.1, 0, 0.1, 1);
        transition: right 0.3s 0.2s cubic-bezier(0.1, 0, 0.1, 1);
    }

    .el-main {
        background-color: #00000f;
        color: #333;
    }

</style>
