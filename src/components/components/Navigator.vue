<template>
    <div>
        <teamform ref="formRef" :dialogformvisible.sync="dialogformvisible"></teamform>
        <el-menu class="el-menu-vertical-demo" background-color="#212221" text-color="#ddedf4"
                 active-text-color="#c83c3c" :collapse="true">

            <el-menu-item index="1" @click="updateFormVisible">
                <i class="el-icon-upload2"></i>
                <span slot="title">{{$t('navigation.share')}}</span>
            </el-menu-item>

            <el-submenu index="2">
                <template slot="title">
                    <i class="el-icon-search"></i>
                    <span>{{$t('navigation.search')}}</span>
                </template>
                <el-menu-item-group :title="$t('navigation.pokemon')">
                    <el-menu-item index="2-1">
                        <search by="pokemon"></search>
                    </el-menu-item>
                </el-menu-item-group>
                <el-menu-item-group :title="$t('navigation.format')">
                    <el-menu-item index="2-2">
                        <search by="format"></search>
                    </el-menu-item>
                </el-menu-item-group>
            </el-submenu>

            <el-menu-item index="3" @click="goBoard" disabled>
                <i class="el-icon-data-board"></i>
                <span slot="title">{{$t('navigation.board')}}</span>
            </el-menu-item>

            <el-menu-item index="4" @click="goStat">
                <i class="el-icon-pie-chart"></i>
                <span slot="title">{{$t('navigation.stat')}}</span>
            </el-menu-item>
        </el-menu>
    </div>
</template>

<script>
    import Search from "./Search";
    import teamform from '../components/Submit'

    export default {
        name: "Navigator",
        components: {
            Search,
            teamform
        },
        data() {
            return {
                dialogformvisible: false
            }
        },
        methods: {
            // submit
            updateFormVisible() {
                this.dialogformvisible = true;
                setTimeout(this.$refs.formRef.resetForm, 100);
            },
            // go to bulletin board page
            goBoard() {
                this.$router.push('/board')
            },
            // go to stat page
            goStat() {
                this.$router.push('/stat')
            }
        }
    }
</script>

<style scoped>

</style>