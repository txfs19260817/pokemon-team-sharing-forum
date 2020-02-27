<template>
    <base-layout :loading.sync="loading" :fail.sync="fail">
        <template v-slot:header>
            <a id="submit" href="#" @click="updateFormVisible">Share</a>
            <teamform ref="formRef" :dialogformvisible.sync="dialogformvisible"></teamform>
        </template>
        <template v-slot:default>
            <!-- Search -->
            <div class="search">
                <el-breadcrumb separator="or">
                    <h2>Search by</h2>
                    <el-breadcrumb-item>
                        <el-cascader
                                v-model="selectedFormat"
                                :options="formats"
                                :placeholder="$t('form.format')"
                                :show-all-levels="false"
                                @change="searchFormat">
                        </el-cascader>
                    </el-breadcrumb-item>
                    <el-breadcrumb-item>
                        <el-select v-model="selectedPokemon" filterable :placeholder="$t('pokemon')"
                                   @change="searchPokemon">
                            <el-option v-for="item in pokemonNames" :key="item" :label="item" :value="item">
                            </el-option>
                        </el-select>
                    </el-breadcrumb-item>
                </el-breadcrumb>
            </div>

            <!-- Display results -->
            <el-row>
                <el-col class="responsive" :span="6" v-for="(item, index) in teams" :key="index" :offset="0">
                    <card :item="item"><img :src="item.src" @click="lightbox(index)" :alt="item.alt"></card>
                </el-col>
            </el-row>
            <div class="clearfix"></div>
            <photoswipe ref="photoswipe" :items="teams"></photoswipe>
        </template>
        <template v-slot:footer>
            <!-- left: button to open the drawer -->
            <img id="about" src="../../public/about.png" width="62" height="62" @click.stop="drawervisible = true"
                 alt="about"/>
            <el-drawer
                    title="About"
                    :size="'60%'"
                    :destroy-on-close="true"
                    :visible.sync="drawervisible"
                    :direction="'ltr'">
                <div class="about" v-html="$t('about')"></div>
            </el-drawer>
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
            <el-select id="lang" size="mini" style="width: 100px;" v-model="curLang" @change="switchLang">
                <el-option v-for="item in languages" :key="item" :value="item"></el-option>
            </el-select>
        </template>
    </base-layout>
</template>

<script>
    import teamform from './components/Submit'
    import photoswipe from "./components/Photoswipe";
    import BaseLayout from "./layouts/BaseLayout";
    import Card from "./layouts/Card";
    import {Formats} from "../assets/data/formats";

    const PokemonNames = require('../assets/data/pokemonNames.js');

    export default {
        name: 'home',
        components: {
            teamform,
            photoswipe,
            BaseLayout,
            Card
        },
        data() {
            return {
                // for base-layout
                loading: false,
                fail: false,
                // lang
                languages: ['zh-hans', 'en', 'ja'],
                curLang: 'zh-hans',
                // dialog and drawer visible
                dialogformvisible: false,
                dialogshowdowntextvisible: false,
                drawervisible: false,
                // page
                total: 1,
                pageSize: 8,
                curPage: 1,
                // temp data
                teams: [],
                showdownText: '',
                // search panel
                formats: [],
                pokemonNames: [],
                selectedPokemon: null,
                selectedFormat: null,
            }
        },
        methods: {
            // submit form dialog
            updateFormVisible() {
                this.dialogformvisible = true;
                setTimeout(this.$refs.formRef.resetForm, 100);
            },
            // request data
            getTeams(page) {
                this.loading = true;
                this.$http.get('teams', {
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
                this.getTeams(v);
                this.curPage = v
            },
            // click img to view details
            lightbox(index) {
                this.$refs.photoswipe.imagePreview(index);
            },
            switchLang(lang) {
                this.$i18n.locale = lang;
                // save lang in localStorage
                localStorage.setItem('lang', lang);
            },
            // search panel
            searchFormat(f) {
                f = f[1];
                this.$router.push('formats/' + f)
            },
            searchPokemon(p) {
                p = p.split('/', 1)[0];
                this.$router.push('pokemon/' + p)
            },
        },
        created() {
            this.getTeams(1);
            this.formats = Formats;
            this.pokemonNames = PokemonNames.pmNames4Select;
            this.curLang = localStorage.getItem('lang') || 'zh-hans'
        },
    }
</script>

<style scoped>
    @import url("https://fonts.lug.ustc.edu.cn/css?family=Montserrat");
    /* black: #212221 blue: #1181b2 lightblue:#ddedf4 purple:#44449b white:ECF5FF*/
    /*Share button*/
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

    /*About Page*/
    div.about {
        height: 600px;
        overflow: auto;
        margin: 0px 40px;
        font-size: 16px;
        line-height: 20px;
        padding: 0;
    }

    /* Gallery */
    div {
        box-sizing: border-box;
    }

    .responsive {
        padding: 6px 6px;
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