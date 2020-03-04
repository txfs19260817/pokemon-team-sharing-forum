<template>
    <el-select
            id="lang"
            size="mini"
            style="width: 100px;"
            v-model="curLang"
            placeholder="Language..."
            @change="switchLang">
        <el-option
                v-for="item in languages"
                :key="item"
                :value="item">
        </el-option>
    </el-select>
</template>

<script>
    export default {
        name: "Language",
        data(){
            return{
                languages: ['zh-hans', 'en', 'ja'],
                curLang: 'zh-hans',
            }
        },
        methods:{
            switchLang(lang) {
                this.$i18n.locale = lang;
                // save lang in localStorage
                localStorage.setItem('lang', lang);
            }
        },
        created() {
            this.curLang = localStorage.getItem('lang') || 'zh-hans'
        }
    }
</script>

<style scoped>

</style>