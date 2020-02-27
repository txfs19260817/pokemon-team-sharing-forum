<template>
    <page-layout :category.sync="category" :server-path.sync="serverPath"></page-layout>
</template>

<script>
    import PageLayout from "../layouts/PageLayout";

    export default {
        name: 'Pokemon',
        components: {
            PageLayout
        },
        data() {
            return {
                category: null,
                serverPath: null
            }
        },
        created() {
            this.category = this.$route.params.pokemon;
            this.serverPath = 'pokemon/'
        },
    }
</script>

<style scoped>
</style>