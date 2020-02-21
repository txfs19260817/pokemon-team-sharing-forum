<template>
    <base-layout :loading.sync="loading" :fail.sync="fail">
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
                                <td>作者</td>
                                <td>{{team.author}}</td>
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
                                    <el-tag type="" v-if="team.pokemon1">{{team.pokemon1}}</el-tag>
                                    <el-tag type="success" v-if="team.pokemon2">{{team.pokemon2}}</el-tag>
                                    <el-tag type="info" v-if="team.pokemon3">{{team.pokemon3}}</el-tag>
                                    <el-tag type="warning" v-if="team.pokemon4">{{team.pokemon4}}</el-tag>
                                    <el-tag type="danger" v-if="team.pokemon5">{{team.pokemon5}}</el-tag>
                                    <el-tag type="" v-if="team.pokemon6">{{team.pokemon6}}</el-tag>
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
                        <el-input
                                id="showdowntext"
                                type="textarea"
                                v-model="team.showdown"
                                :rows="28">
                        </el-input>
                    </el-tab-pane>
                </el-tabs>
            </div>
        </template>
        <template v-slot:footer></template>
    </base-layout>
</template>

<script>
    import BaseLayout from "./BaseLayout";

    export default {
        name: "Team",
        data() {
            return {
                loading: false,
                fail: false,
                team: null,
            }
        },
        components: {
            BaseLayout
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