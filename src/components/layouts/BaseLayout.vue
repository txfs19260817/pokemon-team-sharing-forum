<template>
    <el-container class="home-container">
        <el-header style="height: 180px;">
            <header>
                <router-link to='/home'>
                    <img id="logo" src="../../../public/logo.png" alt="Pokemon Team Sharing Forum"/>
                </router-link>
                <slot name="header"></slot>
            </header>
        </el-header>
        <el-container>
            <el-main>
                <div v-if="loading" class="loading">
                    <h2>Loading...</h2>
                </div>
                <div v-if="fail" class="loading">
                    <h2>Sobble's Sitrus Berry has been took away :'(</h2>
                    <div class="error-image">
                        <img src="../../../public/error.jpg" alt="error" />
                    </div>
                </div>
                <slot></slot>
            </el-main>
            <el-footer style="height: 62px;">
                <slot name="footer"></slot>
            </el-footer>
        </el-container>
    </el-container>
</template>

<script>
    export default {
        name: "BaseLayout",
        props: {
            loading: {
                type: Boolean,
                default: false
            },
            fail: {
                type: Boolean,
                default: false
            },
        }
    }
</script>

<style scoped>
    .error-image {
        text-align: center;
    }
</style>