<template>
    <page-layout :category.sync="category" :server-path.sync="serverPath"></page-layout>
</template>

<script>
    import PageLayout from "./layouts/PageLayout";

    export default {
        name: 'home',
        components: {
            PageLayout
        },
        data() {
            return {
                category: "",
                serverPath: "teams"
            }
        },
    }
</script>

<style scoped>

</style>