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
            <el-container>
                <el-aside v-if="showNav" width="3%">
                    <navigator></navigator>
                </el-aside>
                <el-main v-loading="loading" element-loading-background="rgba(56, 56, 56, 0.9)">
                    <div v-if="loading" class="loading">
                        <h2>Loading...</h2>
                    </div>
                    <div v-if="fail" class="loading">
                        <h2>Sobble's Sitrus Berry has been took away :'(</h2>
                        <div class="error-image">
                            <img src="../../../public/error.jpg" alt="error"/>
                        </div>
                    </div>
                    <slot></slot>
                </el-main>
            </el-container>
            <el-footer style="height: 62px;">
                <slot name="footer-left">
                    <el-tooltip class="item" effect="dark" content="About" placement="top-start">
                        <router-link to='/about'>
                            <img id="about" src="../../../public/about.png" width="62" height="62" alt="about"/>
                        </router-link>
                    </el-tooltip>
                </slot>
                <slot name="footer-middle">
                    <!-- dummy element-->
                    <div style="height: 62px; width: 62px;"></div>
                </slot>
                <slot name="footer-right">
                    <language></language>
                </slot>
            </el-footer>
        </el-container>
    </el-container>
</template>

<script>
    import Language from "../components/Language";
    import Navigator from "../components/Navigator";

    export default {
        name: "BaseLayout",
        components: {
            Language,
            Navigator
        },
        props: {
            loading: {
                type: Boolean,
                default: false
            },
            fail: {
                type: Boolean,
                default: false
            },
            showNav: {
                type: Boolean,
                default: true
            }
        }
    }
</script>

<style scoped>
    .error-image {
        text-align: center;
    }
</style>