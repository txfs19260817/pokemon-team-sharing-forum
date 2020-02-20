<template>
    <div id="upload">
        <el-upload
                :action=destinationUrl
                name="upload"
                list-type="picture-card"
                accept="image/jpeg,image/jpg,image/png"
                :limit=1
                :drag="true"
                :before-upload="handleBeforeUpload"
                :on-exceed="handleExceed"
                :on-remove="handleRemove"
                :on-success="handleSuccess"
                :on-preview="handlePictureCardPreview">
            <i class="el-icon-plus"></i>
        </el-upload>
        <el-dialog :visible.sync="dialogVisible" append-to-body>
            <img width="100%" :src="dialogImageUrl" alt="Team Preview">
        </el-dialog>
    </div>
</template>

<script>
    export default {
        name: "upload",
        props:{
            // uploaded rental image url
            imgurl: {
                type: String,
                default:''
            },
        },
        data() {
            return {
                // server upload api
                destinationUrl: process.env.VUE_APP_UPLOAD,
                // preview url
                dialogImageUrl: '',
                // preview dialog visible
                dialogVisible: false,
            }
        },
        methods: {
            handleBeforeUpload(file) {
                let test = /^image\/(jpeg|png|jpg)$/.test(file.type);
                const isLt2M = file.size / 1024 / 1024 < 2;
                if (!test) {
                    this.$message.error('上传图片格式有误，支持格式：png, jpg, jpeg');
                    return false
                }
                if (!isLt2M) {
                    this.$message.error('上传图片大小不能超过 2MB');
                    return false
                }
                return test && isLt2M
            },
            handleExceed(files, fileList) {
                this.$message.warning(`请删除后重新添加图片`);
            },
            handleRemove(file, fileList) {
                this.$emit('update:imgurl', '')
            },
            handleSuccess(response, file, fileList) {
                if (response.code !== 200) {
                    this.$message.error('上传失败：' + response.data);
                    return
                }
                this.$emit('update:imgurl', response.data)
            },
            handlePictureCardPreview(file) {
                this.dialogImageUrl = file.url;
                this.dialogVisible = true;
            },
        }
    }
</script>

<style scoped>

</style>