<template>
    <div id="teamform">
        <el-dialog title="队伍表格" :visible.sync="dialogformvisible" :before-close="closeForm">
            <el-form :model="form">
                <el-form-item label="标题" :label-width="formLabelWidth">
                    <el-input v-model="form.title" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="作者" :label-width="formLabelWidth">
                    <el-input v-model="form.author" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="模式" :label-width="formLabelWidth">
                    <el-select v-model="form.format" placeholder="请选择队伍适合的对战模式">
                        <el-option v-for="item in formats" :key="item" :label="item" :value="item"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="队伍租借ID截图上传(可选)" :label-width="formLabelWidth">
                    <el-upload
                        :action=uploadUrl
                        name="upload"
                        list-type="picture-card"
                        :limit=1
                        :on-exceed="handleExceed"
                        accept="image/gif,image/jpeg,image/jpg,image/png,image/bmp"
                        :on-success="handleSuccess"
                        :on-preview="handlePictureCardPreview">
                        <i class="el-icon-plus"></i>
                    </el-upload>
                    <el-dialog :visible.sync="dialogVisible" append-to-body>
                        <img width="100%" :src="dialogImageUrl" alt="">
                    </el-dialog>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="closeForm">取 消</el-button>
                <el-button type="primary" @click="closeForm">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    export default {
        name: "teamform",
        props: {
            dialogformvisible: {
                type: Boolean,
                default: false
            }
        },
        data() {
            return {
                url:"http://127.0.0.1:8888/",
                formats: [],
                form: {
                    title: '',
                    author: '',
                    format: '',
                    rentalImgUrl:''
                },
                formLabelWidth: '120px',
                // Upload image dialog properties
                dialogImageUrl: '',
                dialogVisible: false,
            }
        },
        created() {
            this.$http.get('formats').then(r =>
                this.formats = r.data.data
            ).catch(error =>
                console.log(error)
            )
        },
        methods: {
            // for this form
            closeForm() {
                this.dialogformvisible = false;
                this.$emit("update:dialogformvisible", this.dialogformvisible)
            },
            // methods for image-upload dialog
            handleExceed(files, fileList) {
                this.$message.warning(`请删除后重新添加图片`);
            },
            handlePictureCardPreview(file) {
                this.dialogImageUrl = file.url;
                this.dialogVisible = true;
                console.log(this.dialogImageUrl)
            },
            handleSuccess(response, file, fileList){
                this.form.rentalImgUrl = response.data
                console.log(this.form.rentalImgUrl)
            }
        },
        computed: {
            uploadUrl(){
                return this.url+"api/v1/upload"
            }
        }
    }
</script>

<style scoped>

</style>