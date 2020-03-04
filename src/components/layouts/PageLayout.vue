<template>
    <base-layout :loading.sync="loading" :fail.sync="fail">
        <template v-slot:header></template>
        <template v-slot:default>
            <h3>{{category}}</h3>
            <el-row>
                <el-col class="responsive" :span="6" v-for="(item, index) in teams" :key="index" :offset="0">
                    <card :item="item"><img :src="item.src" @click="lightbox(index)" :alt="item.alt"></card>
                </el-col>
            </el-row>
            <div class="clearfix"></div>
            <photoswipe ref="photoswipe" :items="teams"></photoswipe>
        </template>
        <!--        footer-->
        <template v-slot:footer-middle>
            <el-pagination
                    background
                    layout="prev, pager, next"
                    @current-change="handleCurrentChange"
                    :current-page="curPage"
                    :page-size="pageSize"
                    :total="total">
            </el-pagination>
        </template>
    </base-layout>
</template>

<script>
    import photoswipe from "../components/Photoswipe";
    import BaseLayout from "./BaseLayout";
    import Card from "./Card";

    export default {
        name: 'PageLayout',
        components: {
            photoswipe,
            BaseLayout,
            Card
        },
        props: {
            // e.g. "formats/", "pokemon/"
            serverPath: {
                type: String,
                required: true
            },
            // e.g. "VGC 2020" (if "formats/"), "Incineroar" (if "pokemon/")
            category: {
                type: String,
                required: true
            }
        },
        data() {
            return {
                serverUrl: '',
                // for base-layout
                loading: false,
                fail: false,
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
            getTeamsBy(serverUrl, page) {
                this.loading = true;
                this.$http.get(serverUrl, {
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
                this.getTeamsBy(this.serverUrl, v);
                this.curPage = v
            },
            // click img to view details
            lightbox(index) {
                this.$refs.photoswipe.imagePreview(index);
            },
        },
        created() {
            this.serverUrl = this.serverPath + this.category;
            this.getTeamsBy(this.serverUrl, 1);
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