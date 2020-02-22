<template>
    <base-layout :loading.sync="loading" :fail.sync="fail">
        <template v-slot:header></template>
        <template v-slot:default>
            <h3>{{format}}</h3>
            <el-row>
                <el-col class="responsive" :span="6" v-for="(item, index) in teams" :key="index" :offset="0">
                    <card :item="item"><img :src="item.src" @click="lightbox(index)" :alt="item.alt"></card>
                </el-col>
            </el-row>
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
    import Card from "./Card";

    export default {
        name: 'Format',
        components: {
            photoswipe,
            BaseLayout,
            Card
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
                            d.src = t.rentalImgUrl; // expand some fields below for Photoswipe
                            d.alt = t.title; // alt:=title; title:=author+description
                            d.title = t.author + ': ' + t.description;
                            [d.w, d.h] = [1024, 576];
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
    /* Gallery */
    div {
        box-sizing: border-box;
    }

    .responsive {
        padding: 0 6px;
        margin: 6px 0;
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

    .el-row img {
        width: 100%;
        display: block;
    }
</style>