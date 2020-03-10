<template>
    <base-layout :loading.sync="loading" :fail.sync="fail">
        <!-- The v-if must be outside otherwise "display: grid;" does not have effect. -->
        <div v-if="team">
            <div class="wrapper">
                <!--                Title-->
                <div class="box header">[{{team.format}}] {{ team.title }}</div>

                <!--                Team Preview Image-->
                <div class="box sidebar">
                    <img :src="team.rentalImgUrl" :alt="team.title" :title="team.title" width="768" height="432"/>
                </div>

                <!--                Description-->
                <div class="box sidebar2">
                    <h3>{{$t('form.description')}}</h3>
                    {{team.description}}
                </div>

                <!--                Other information-->
                <div class="box content">
                    <!--                Post By-->
                    <div class="grid-cell" style="background-color: #002766">
                        <div class="grid-sub-cell">{{$t('form.author')}}:</div>
                        <div class="grid-sub-cell">{{team.author}}</div>
                    </div>

                    <!--                    Time-->
                    <div class="grid-cell" style="background-color: #003a8c">
                        <div class="grid-sub-cell">{{$t('form.created_at')}}:</div>
                        <div class="grid-sub-cell">{{DateConversion(team.created_at)}}</div>
                    </div>

                    <!--                    Format-->
                    <div class="grid-cell" style="background-color: #0050b3">
                        <div class="grid-sub-cell">{{$t('form.format')}}:</div>
                        <div class="grid-sub-cell">
                            <router-link :to="'/formats/'+team.format">
                                <el-link style="font-size: 24px" type="primary">{{team.format}}</el-link>
                            </router-link>
                        </div>
                    </div>

                    <!--                    Pokemon icons-->
                    <div class="grid-cell" style="background-color: #003a8c">
                        <div class="grid-sub-cell">{{$t('pokemon')}}:</div>
                        <div class="grid-sub-cell">
                            <el-breadcrumb separator=" ">
                                <el-breadcrumb-item :to="'/pokemon/'+team.pokemon1">
                                    <img class="icon" :title="team.pokemon1" :alt="team.pokemon1"
                                         :src="IconPath(team.pokemon1, 'pokemon')">
                                </el-breadcrumb-item>
                                <el-breadcrumb-item :to="'/pokemon/'+team.pokemon2">
                                    <img class="icon" :title="team.pokemon1" :alt="team.pokemon2"
                                         :src="IconPath(team.pokemon2, 'pokemon')">
                                </el-breadcrumb-item>
                                <el-breadcrumb-item :to="'/pokemon/'+team.pokemon3">
                                    <img class="icon" :title="team.pokemon1" :alt="team.pokemon3"
                                         :src="IconPath(team.pokemon3, 'pokemon')">
                                </el-breadcrumb-item>
                                <el-breadcrumb-item :to="'/pokemon/'+team.pokemon4">
                                    <img class="icon" :title="team.pokemon1" :alt="team.pokemon4"
                                         :src="IconPath(team.pokemon4, 'pokemon')">
                                </el-breadcrumb-item>
                                <el-breadcrumb-item :to="'/pokemon/'+team.pokemon5">
                                    <img class="icon" :title="team.pokemon1" :alt="team.pokemon5"
                                         :src="IconPath(team.pokemon5, 'pokemon')">
                                </el-breadcrumb-item>
                                <el-breadcrumb-item :to="'/pokemon/'+team.pokemon6">
                                    <img class="icon" :title="team.pokemon1" :alt="team.pokemon6"
                                         :src="IconPath(team.pokemon6, 'pokemon')">
                                </el-breadcrumb-item>
                                <div style="display:none;">
                                    A dummy div to make the cursor appears as a pointer,
                                    when hovering over the last tag.
                                </div>
                            </el-breadcrumb>
                        </div>
                    </div>

                    <!--                    Social Platform Sharing Buttons-->
                    <div class="grid-cell" style="background-color: #002766">
                        <div class="grid-sub-cell">{{$t('form.share')}}:</div>
                        <div class="grid-sub-cell"><share :team="team"></share></div>
                    </div>
                </div>

                <!--                Showdown Importable-->
                <div class="box footer">
                    <h3>{{$t('form.showdown')}}</h3>
                    <el-input id="showdowntext" type="textarea" v-model="team.showdown" :rows="30"></el-input>
                </div>
            </div>
        </div>
    </base-layout>
</template>

<script>
    import BaseLayout from "../layouts/BaseLayout";
    import Share from "../components/Share";
    import {DateConversion, IconPath} from "../../assets/utils";

    export default {
        name: "Team",
        data() {
            return {
                loading: false,
                fail: false,
                team: null,
                DateConversion: DateConversion,
                IconPath: IconPath
            }
        },
        components: {
            BaseLayout,
            Share
        },
        watch: {
            // call again the method if the route changes
            '$route': 'fetchData'
        },
        methods: {
            // request data
            getTeamById(id) {
                this.loading = true;
                this.$http.get('teams/' + id).then(res => {
                    this.loading = false;
                    if (res.data.code === 200) {
                        this.team = res.data.data
                    } else {
                        this.fail = true;
                        this.$message.error("请求数据错误");
                    }
                }).catch(err => {
                    this.loading = false;
                    this.fail = true;
                    this.$message.error("请求数据错误");
                })
            },
        },
        created() {
            this.getTeamById(this.$route.params.id)
        }
    }
</script>

<style scoped>
    body {
        margin: 40px;
    }

    .sidebar {
        grid-area: sidebar;
    }

    .sidebar2 {
        grid-area: sidebar2;
    }


    .header {
        grid-area: header;
        background-color: #999;
    }

    .footer {
        grid-area: footer;
        background-color: #999;
    }

    .wrapper {
        border-radius: 5px;
        background-color: #2d2d2d;
        color: #444;
        display: grid;
        grid-gap: 1em;
        grid-column-gap: 4em;
        grid-template-areas: "header" "sidebar" "content" "sidebar2" "footer"
    }

    @media only screen and (min-width: 1350px) {
        .wrapper {
            grid-template-columns: 788px auto;
            grid-template-areas: "header   header" "sidebar  content" "sidebar2 sidebar2" "footer   footer";
        }
    }

    .wrapper > div img {
        max-width: 100%;
        height: auto;
        border-radius: 5px;
    }

    .box {
        background-color: #444;
        color: #fff;
        border-radius: 5px;
        padding: 10px;
        font-size: 150%;
    }

    .content {
        grid-area: content;
        display: flex;
        flex-direction: column;
        padding: 20px;
    }

    .grid-cell {
        flex: 1;
        display: flex;
        justify-content: space-between;
        border-radius: 5px;
    }

    .grid-sub-cell {
        font-size: 24px;
        line-height: 24px;
        padding: 4px;
    }

    .icon {
        width: auto;
        height: 40px;
    }
</style>