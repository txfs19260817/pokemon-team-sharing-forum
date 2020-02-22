<template>
    <el-card :body-style="{ padding: '0px' }" shadow="hover">
        <!-- team preview image -->
        <slot></slot>
        <!-- showdown text dialog -->
        <el-dialog
                title="Showdown"
                :modal-append-to-body="false"
                :visible.sync="dialogshowdowntextvisible"
                width="30%">
            <el-input id="showdowntext" type="textarea" v-model="showdownText" :rows="6"></el-input>
            <span slot="footer" class="dialog-footer">
                <el-button icon="el-icon-document-copy" circle @click="copyButton"></el-button>
                <el-button type="primary" icon="el-icon-circle-check" circle
                           @click="dialogshowdowntextvisible = false"></el-button>
            </span>
        </el-dialog>
        <!-- description below -->
        <div class="bottom">
            <router-link :to="'/team/' + item.id">
                <span>{{item.alt}}</span>
            </router-link>
            <div class="info clearfix">
                <el-breadcrumb separator="/">
                    <el-breadcrumb-item :to="'/formats/'+item.format">[{{item.format}}]</el-breadcrumb-item>
                    <el-breadcrumb-item>{{$t('form.author')}}: {{item.author}}</el-breadcrumb-item>
                    <el-breadcrumb-item>
                        <time>{{DateConversion(item.created_at)}}</time>
                    </el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <el-row>
                <el-col :span="3" :offset="21">
                    <a href="#">
                        <img id="showdown" src="../../public/showdown.png" class="s-button"
                             @click.stop="showdownButton(item.showdown)" alt="showdown"/>
                    </a>
                </el-col>
            </el-row>
        </div>
    </el-card>
</template>

<script>
    import {DateConversion} from "../assets/utils";

    export default {
        name: "Card",
        props: {
            item: {
                type: Object,
                required: true
            },
        },
        data() {
            return {
                dialogshowdowntextvisible: false,
                showdownText: '',
                DateConversion: DateConversion
            }
        },
        methods: {
            // show showdown text
            showdownButton(text) {
                if (text.length > 0) {
                    this.showdownText = text;
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
        }
    }
</script>

<style scoped>
    time {
        font-size: 13px;
        color: #999;
    }

    a {
        text-decoration: none
    }

    .bottom {
        overflow: auto;
        padding: 14px;
        height: 140px;
        font-size: 14px;
    }

    .info {
        margin-top: 13px;
        line-height: 12px;
    }

    .s-button {
        padding: 0;
        float: right;
    }

    .el-row img {
        width: 100%;
        display: block;
    }
</style>