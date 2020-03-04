<template>
    <div>
        <teamform ref="formRef" :dialogformvisible.sync="dialogformvisible"></teamform>
        <el-menu class="el-menu-vertical-demo" background-color="#212221" text-color="#ddedf4" active-text-color="#c83c3c">
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

            <el-menu-item index="3" v-if="false">
                <i class="el-icon-menu"></i>
                <span slot="title">导航二</span>
            </el-menu-item>
            <el-menu-item index="4" v-if="false">
                <i class="el-icon-document"></i>
                <span slot="title">导航三</span>
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
            updateFormVisible() {
                this.dialogformvisible = true;
                setTimeout(this.$refs.formRef.resetForm, 100);
            },
        }
    }
</script>

<style scoped>

</style>