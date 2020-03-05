<template>
    <base-layout :loading.sync="loading" :fail.sync="fail">
        <template v-slot:default>
            <h2>{{$t('stat.title')}}</h2>
            <h3>{{$t('form.placeholder.format')}}</h3>
            <el-cascader
                    v-model="selectedFormat"
                    :options="formats"
                    :placeholder="$t('form.format')"
                    :show-all-levels="false"
                    @change="showStat">
            </el-cascader>
            <div v-show="statIsAvailable" id="pie" style="width: 800px; height:600px;"></div>
            <h3 v-show="!statIsAvailable">{{$t('stat.noData')}}</h3>
        </template>
    </base-layout>
</template>

<script>
    import BaseLayout from "../layouts/BaseLayout";
    import {Formats} from "../../assets/data/formats";
    import {ReconstructObject, SortObjectArrayByValue, IconPath} from "../../assets/utils";

    let echarts = require('echarts');

    export default {
        name: "Stat",
        components: {
            BaseLayout
        },
        data() {
            return {
                loading: false,
                fail: false,
                // echarts lib
                echarts: null,
                // selector
                formats: [],
                selectedFormat: null,
                // stat data
                statData: {},
                statIsAvailable: false,
            }
        },
        methods: {
            showStat(f) {
                this.selectedFormat = f[1];
                this.getStatByFormat(this.selectedFormat);
            },
            drawPieChart(d) {
                // http://127.0.0.1:8888/assets/pokemon-icons/2d/abomasnow.png
                this.echarts.init(document.getElementById('pie'), 'light').setOption({
                    title: {
                        text: this.selectedFormat + ' 使用率统计',
                        subtext: '数据仅包含已上传队伍',
                        left: 'center'
                    },
                    tooltip: {
                        trigger: 'item',
                        formatter: function (params, ticket, callback) {
                            let img = `<img src="` + IconPath(params.data.name, "pokemon") +
                                `" alt="` + params.data.name + `" />`;
                            let res = params.seriesName + img + "<br />" +
                                params.data.name + ': ' + params.data.value +
                                ' (' + params.percent + '%)';
                            setTimeout(function () {
                                callback(ticket, res);
                            }, 100);
                            return 'loading...';
                        }
                    },
                    series: [
                        {
                            name: this.selectedFormat + ' Usage',
                            type: 'pie',
                            radius: '55%',
                            roseType: 'angle',
                            data: d
                        }
                    ]
                });
            },
            getStatByFormat(format) {
                this.loading = true;
                this.$http.get('stat', {
                    params: {
                        format: format
                    }
                }).then(res => {
                    this.loading = false;
                    if (res.data.code === 200) {
                        this.statData = res.data.data;
                        // check object empty
                        if (Object.keys(this.statData).length === 0 && this.statData.constructor === Object) {
                            this.statIsAvailable = false;
                        } else {
                            this.statData = SortObjectArrayByValue(ReconstructObject(this.statData));
                            this.drawPieChart(this.statData);
                            this.statIsAvailable = true;
                        }
                    } else {
                        this.fail = true;
                        this.statIsAvailable = false;
                        this.$message.error("请求数据错误");
                    }
                }).catch(err => {
                    this.loading = false;
                    this.fail = true;
                    this.statIsAvailable = false;
                    this.$message.error("请求数据错误");
                })
            }
        },
        created() {
            this.formats = Formats;
            this.echarts = echarts;
        }
    }
</script>

<style scoped>

</style>