<template>
    <base-layout :loading.sync="loading" :fail.sync="fail" :show-nav="false">
        <template v-slot:header></template>
        <template v-slot:default>
            <div v-if="team" class="content">
                <h2>{{ team.title }}</h2>
                <el-tabs type="border-card">
                    <!--                    Main information Tab-->
                    <el-tab-pane class="info" label="Info">
                        <table>
                            <tbody>
                            <tr>
                                <td>队伍预览图</td>
                                <td><img :src="team.rentalImgUrl" alt="team.title" width="768" height="432"/></td>
                            </tr>
                            <tr>
                                <td>分享</td>
                                <td><share :team="team"></share></td>
                            </tr>
                            <tr>
                                <td>{{$t('form.author')}}</td>
                                <td>{{team.author}}</td>
                            </tr>
                            <tr>
                                <td>上传时间</td>
                                <td>{{DateConversion(team.created_at)}}</td>
                            </tr>
                            <tr>
                                <td>对战模式</td>
                                <td>
                                    <router-link :to="'/formats/'+team.format">
                                        <el-tag type="info">{{team.format}}</el-tag>
                                    </router-link>
                                </td>
                            </tr>
                            <tr>
                                <td>宝可梦</td>
                                <td>
                                    <el-breadcrumb separator=" ">
                                        <el-breadcrumb-item :to="'/pokemon/'+team.pokemon1">
                                            <el-tag type="" v-if="team.pokemon1">{{team.pokemon1}}</el-tag>
                                        </el-breadcrumb-item>
                                        <el-breadcrumb-item :to="'/pokemon/'+team.pokemon2">
                                            <el-tag type="success" v-if="team.pokemon2">{{team.pokemon2}}</el-tag>
                                        </el-breadcrumb-item>
                                        <el-breadcrumb-item :to="'/pokemon/'+team.pokemon3">
                                            <el-tag type="info" v-if="team.pokemon3">{{team.pokemon3}}</el-tag>
                                        </el-breadcrumb-item>
                                        <el-breadcrumb-item :to="'/pokemon/'+team.pokemon4">
                                            <el-tag type="warning" v-if="team.pokemon4">{{team.pokemon4}}</el-tag>
                                        </el-breadcrumb-item>
                                        <el-breadcrumb-item :to="'/pokemon/'+team.pokemon5">
                                            <el-tag type="danger" v-if="team.pokemon5">{{team.pokemon5}}</el-tag>
                                        </el-breadcrumb-item>
                                        <el-breadcrumb-item :to="'/pokemon/'+team.pokemon6">
                                            <el-tag effect="dark" v-if="team.pokemon6">{{team.pokemon6}}</el-tag>
                                        </el-breadcrumb-item>
                                        <div style="display:none;">
                                            A dummy div to make the cursor appears as a pointing hand,
                                            when hovering over the last tag.
                                        </div>
                                    </el-breadcrumb>
                                </td>
                            </tr>
                            <tr>
                                <td>描述</td>
                                <td>{{team.description}}</td>
                            </tr>
                            </tbody>
                        </table>
                    </el-tab-pane>
                    <!--                    Showdown Tab-->
                    <el-tab-pane class="showdown" label="Showdown">
                        <el-input id="showdowntext" type="textarea" v-model="team.showdown" :rows="40"></el-input>
                    </el-tab-pane>
                </el-tabs>
            </div>
        </template>
        <template v-slot:footer></template>
    </base-layout>
</template>

<script>
    import BaseLayout from "../layouts/BaseLayout";
    import Share from "../components/Share";
    import {DateConversion} from "../../assets/utils";

    export default {
        name: "Team",
        data() {
            return {
                loading: false,
                fail: false,
                team: null,
                DateConversion: DateConversion
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
    .info {
        display: flex;
        justify-content: center;
    }

    table {
        border-collapse: collapse;
    }

    table, td, th {
        border: 1px solid gainsboro;
    }
</style>