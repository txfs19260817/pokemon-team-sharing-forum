<template>
    <el-container class="home-container">
        <!-- header -->
        <el-header style="height: 180px;">
            <header>
                <img id="logo" src="../../public/logo.png" alt="Pokemon Team Sharing Forum"/>
                <a id="submit" href="#" @click="updateFormVisible">Submit</a>
                <teamform ref="formRef" :dialogformvisible.sync="dialogformvisible" :url="url"></teamform>
            </header>
        </el-header>
        <el-container>
            <!-- main -->
            <el-main>
                <div class="responsive" v-for="(item, index) in teams">
                    <div class="gallery">
                        <img :src="item.src" @click="lightbox(index)" :alt="item.alt" width="512" height="288">
                        <div class="desc">
                            <span class="title">{{item.title}}</span>
                            <span class="author">上传用户：{{item.author}}</span>
                        </div>
                    </div>
                </div>
                <div class="clearfix"></div>
                <photoswipe ref="photoswipe" :items="teams"></photoswipe>
            </el-main>
            <!-- footer -->
            <el-footer>
                <el-pagination
                        background
                        layout="prev, pager, next"
                        @current-change="handleCurrentChange"
                        :current-page="curPage"
                        :page-size="pageSize"
                        :total="total">
                </el-pagination>
            </el-footer>
        </el-container>
    </el-container>
</template>

<script>
    import teamform from './Submit'
    import photoswipe from "./Photoswipe";

    export default {
        name: 'home',
        components: {
            teamform,
            photoswipe
        },
        data() {
            return {
                url: "http://127.0.0.1:8888/",
                // dialog
                dialogformvisible: false,
                // data - num
                total: 1,
                pageSize: 8,
                curPage: 1,
                // data
                teams: []
            }
        },
        methods: {
            // submit dialog
            updateFormVisible() {
                this.dialogformvisible = true;
                setTimeout(this.$refs.formRef.resetForm, 100);
            },
            // request data
            getTeams(page) {
                this.$http.get('teams', {
                    params: {
                        state: 1,
                        page: page
                    }
                }).then(res => {
                    if (res.data.code === 200) {
                        console.log(res);
                        this.total = res.data.data.total;
                        this.teams.splice(0, this.teams.length); // empty the array
                        for (let t of res.data.data.lists) {
                            let d = t;
                            // expand some fields for Photoswipe
                            d.src = t.rentalImgUrl;
                            d.thumbnail = t.rentalImgUrl;
                            d.alt = t.title;
                            d.w = 1024;
                            d.h = 576;
                            this.teams.push(d);
                        }
                    } else {
                        this.$message.error("请求数据错误");
                    }
                }).catch(err => {
                    this.$message.error("请求数据错误");
                    console.log(err)
                })
            },
            // pagination handlers
            handleCurrentChange(v) {
                this.getTeams(v);
                this.curPage = v
            },
            lightbox(index) {
                this.$refs.photoswipe.imagePreview(index);
            }
        },
        created() {
            this.getTeams(1);
        },
    }
</script>

<style scoped>
    @import url("https://fonts.lug.ustc.edu.cn/css?family=Montserrat");
    /* black: #212221 blue: #1181b2 lightblue:#ddedf4 purple:#44449b white:ECF5FF*/
    .home-container {
        height: 100%;
        font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    }

    #logo {
        height: 50%;
        position: absolute;
        left: 0.5%;
        top: 5%;
    }

    header {
        position: sticky;
        width: 100%;
        height: 15em;
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
        height: 24em;
        z-index: -1;
        transform: skewY(-2.0deg);
        background: linear-gradient(rgba(0, 0, 0, .6), rgba(0, 0, 0, .6)),
        url(../../public/header.jpg) no-repeat center,
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
        color: #c83c3c;
        text-decoration: none;
        font-size: 2em;
        display: inline-block;
        font-family: Montserrat, sans-serif;
        text-transform: uppercase;
        padding: 0.5em 2em;
        border: 2px solid #c83c3c;
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
        background: #c83c3c;
        -webkit-transition: 0.3s 0.2s cubic-bezier(0.1, 0, 0.1, 1), left 0.3s cubic-bezier(0.1, 0, 0.1, 1);
        transition: 0.3s 0.2s cubic-bezier(0.1, 0, 0.1, 1), left 0.3s cubic-bezier(0.1, 0, 0.1, 1);
        z-index: -1;
    }

    #submit::after {
        content: "";
        display: inline-block;
        background-image: url(../../public/arrow.png);
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

    .el-header, .el-main, .el-footer {
        background-color: #212221;
        color: #ddedf4;
    }

    .el-header {
        padding: 0;
    }

    .el-main {
        padding: 20px 40px;
    }

    .el-footer {
        text-align: center;
        line-height: 60px;
    }

    /*teams preview https://www.w3schools.com/css/css_image_gallery.asp */
    div {
        box-sizing: border-box;
    }

    div.gallery {
        margin: 0 0 10px 0;
        border: 1px solid #ddedf4;
    }

    div.gallery:hover {
        border: 1px solid #44449b;
    }

    div.gallery img {
        width: 100%;
        height: auto;
    }

    div.desc {
        padding: 15px;
        /*text-align: center;*/
        display: flex;
        justify-content: space-between;
    }
    span.author {

        align-self: flex-end;
    }

    .responsive {
        padding: 0 6px;
        float: left;
        width: 24.99999%;
    }

    @media only screen and (max-width: 700px) {
        .responsive {
            width: 49.99999%;
            margin: 6px 0;
        }
    }

    @media only screen and (max-width: 500px) {
        .responsive {
            width: 100%;
        }
    }

    .clearfix:after {
        content: "";
        display: table;
        clear: both;
    }

</style>