<template>
    <base-layout :loading.sync="loading" :fail.sync="fail">
        <template v-slot:header></template>
        <template v-slot:default>
            <h3>{{format}}</h3>
            <div class="responsive" v-for="(item, index) in teams">
                <div class="gallery">
                    <!-- team preview -->
                    <img :src="item.src" @click="lightbox(index)" :alt="item.alt" width="512" height="288">
                    <!-- showdown text dialog -->
                    <el-dialog
                            title="Showdown"
                            :modal-append-to-body="false"
                            :visible.sync="dialogshowdownvisible"
                            width="30%">
                        <el-input
                                id="showdowntext"
                                type="textarea"
                                v-model="showdownText"
                                :rows="6">
                        </el-input>
                        <span slot="footer" class="dialog-footer">
                                <el-button @click="copyButton">Copy to clipboard</el-button>
                                <el-button type="primary" @click="dialogshowdownvisible = false">确 定</el-button>
                            </span>
                    </el-dialog>
                    <!-- caption below the image -->
                    <div class="desc">
                        <div class="column">
                            <!-- title with link to detail page -->
                            <router-link :to="'/team/'+item.id">
                                <span class="title">[{{item.format}}]{{item.alt}}</span>
                            </router-link>
                            <!-- author -->
                            <span class="author">作者：{{item.author}}</span>
                        </div>
                        <div class="column">
                            <!-- time -->
                            <span class="time">{{item.created_at}}</span>
                            <!-- showdown button to open the dialog -->
                            <a href="#">
                                <img id="showdown" src="../../public/showdown.png"
                                     @click.stop="showdownButton(index)" alt="showdown"/>
                            </a>
                        </div>
                    </div>
                </div>
            </div>
            <div class="clearfix"></div>
            <photoswipe ref="photoswipe" :items="teams"></photoswipe>
        </template>
        <template v-slot:footer>
            <!-- left: dummy element -->
            <div style="height: 62px;width: 62px;"></div>
            <!-- middle: paginator  -->
            <el-pagination
                    background
                    layout="prev, pager, next"
                    @current-change="handleCurrentChange"
                    :current-page="curPage"
                    :page-size="pageSize"
                    :total="total">
            </el-pagination>
            <!-- right: language selector  -->
            <el-select id="lang" size="mini" style="width: 100px;" v-model="curLang" placeholder="Language..."
                       @change="switchLang">
                <el-option
                        v-for="item in languages"
                        :key="item"
                        :value="item">
                </el-option>
            </el-select>
        </template>
    </base-layout>
</template>

<script>
    import photoswipe from "./Photoswipe";
    import BaseLayout from "./BaseLayout";

    export default {
        name: 'Format',
        components: {
            photoswipe,
            BaseLayout
        },
        data() {
            return {
                format:'',
                url: process.env.VUE_APP_URL,
                // base
                loading: false,
                fail: false,
                // lang
                languages: ['zh-hans', 'en', 'ja'],
                curLang: 'zh-hans',
                // showdown text dialog visible
                dialogshowdowntextvisible: false,
                // page
                total: 1,
                pageSize: 8,
                curPage: 1,
                // data
                teams: [],
                showdownText: ''
            }
        },
        methods: {
            // request data
            getTeamsByFormat(format, page) {
                this.loading = true;
                this.$http.get('formats/' + format, {
                    params: {
                        state: 1,
                        page: page
                    }
                }).then(res => {
                    this.loading = false;
                    if (res.data.code === 200) {
                        this.total = res.data.data.total;
                        this.teams.splice(0, this.teams.length); // empty the array
                        for (let t of res.data.data.lists) {
                            let d = t;
                            // expand some fields for Photoswipe
                            // alt:=title; title:=author+description
                            d.src = t.rentalImgUrl;
                            d.alt = t.title;
                            d.title = t.author + ': ' + t.description;
                            d.w = 1024;
                            d.h = 576;
                            this.teams.push(d);
                        }
                    } else {
                        this.fail = true;
                        this.$message.error("请求数据错误");
                    }
                }).catch(error => {
                    // Error
                    if (error.response) {
                        /*
                         * The request was made and the server responded with a
                         * status code that falls out of the range of 2xx
                         */
                        console.log(error.response.data);
                        console.log(error.response.status);
                        console.log(error.response.headers);
                    } else if (error.request) {
                        /*
                         * The request was made but no response was received, `error.request`
                         * is an instance of XMLHttpRequest in the browser and an instance
                         * of http.ClientRequest in Node.js
                         */
                        console.log(error.request);
                    } else {
                        // Something happened in setting up the request and triggered an Error
                        console.log('Error', error.message);
                    }
                    console.log(error.config);

                    this.loading = false;
                    this.fail = true;
                    this.$message.error("请求数据错误");
                })
            },
            // pagination handlers
            handleCurrentChange(v) {
                this.getTeamsByFormat(this.format, v);
                this.curPage = v
            },
            // click img to view details
            lightbox(index) {
                this.$refs.photoswipe.imagePreview(index);
            },
            // show showdown text
            showdownButton(index) {
                if (this.teams[index].showdown.length > 0) {
                    this.showdownText = this.teams[index].showdown;
                    this.dialogshowdowntextvisible = true;
                } else {
                    this.$message.warning("该队伍没有提供showdown队伍文本");
                    this.dialogshowdowntextvisible = false;
                }
            },
            copyButton() {
                let copyText = document.getElementById("showdowntext");
                copyText.select();
                copyText.setSelectionRange(0, 99999);
                document.execCommand("copy");
            },
            switchLang(lang) {
                this.$i18n.locale = lang;
                // save lang in localStorage
                localStorage.setItem('lang', lang);
            }
        },
        created() {
            this.format = this.$route.params.format;
            this.getTeamsByFormat(this.format, 1);
            this.curLang = localStorage.getItem('lang') || 'zh-hans'
        },
    }
</script>

<style scoped>
    /*teams preview https://www.w3schools.com/css/css_image_gallery.asp */
    div {
        box-sizing: border-box;
    }

    div.gallery {
        margin: 0 0 10px 0;
        border: 1px solid #e8dcdc;
    }

    div.gallery:hover {
        border: 1px solid #c83c3c;
    }

    div.gallery img {
        width: 100%;
        height: auto;
    }

    div.desc {
        height: 80px;
        font-size: 14px;
        display: flex;
        flex-wrap: wrap;
        align-content: space-between;
    }

    div.column {
        flex-basis: 100%;
        display: flex;
        justify-content: space-between;
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