<template>
    <div id="teamform">
        <el-dialog title="队伍表格" :visible.sync="dialogformvisible" :before-close="closeForm" width="70%"
                   :modal-append-to-body="false">
            <el-form :model="form" ref="teamFormRef" :rules="loginFormRules">
                <el-form-item label="标题" :label-width="formLabelWidth" prop="title">
                    <el-input v-model="form.title" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="作者" :label-width="formLabelWidth" prop="author">
                    <el-input v-model="form.author" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="模式" :label-width="formLabelWidth" prop="format">
                    <el-cascader
                            v-model="form.format"
                            :options="formats"
                            placeholder="请选择队伍适合的对战模式"
                            :show-all-levels="false">
                    </el-cascader>
                </el-form-item>
                <el-form-item label="第1只宝可梦" :label-width="formLabelWidth" prop="pokemon1">
                    <el-select v-model="form.pokemon1" filterable placeholder="可在此输入宝可梦名称">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="第2只宝可梦" :label-width="formLabelWidth" prop="pokemon2">
                    <el-select v-model="form.pokemon2" filterable placeholder="可在此输入宝可梦名称">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="第3只宝可梦" :label-width="formLabelWidth" prop="pokemon3">
                    <el-select v-model="form.pokemon3" filterable placeholder="可在此输入宝可梦名称">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="第4只宝可梦" :label-width="formLabelWidth" prop="pokemon4">
                    <el-select v-model="form.pokemon4" filterable placeholder="可在此输入宝可梦名称">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="第5只宝可梦" :label-width="formLabelWidth" prop="pokemon5">
                    <el-select v-model="form.pokemon5" filterable placeholder="可在此输入宝可梦名称">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="第6只宝可梦" :label-width="formLabelWidth" prop="pokemon6">
                    <el-select v-model="form.pokemon6" filterable placeholder="可在此输入宝可梦名称">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="队伍租借ID截图(可选，支持格式：png, jpg, jpeg)" :label-width="formLabelWidth"
                              prop="rentalImgUrl">
                    <el-upload
                            :action=uploadUrl
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
                </el-form-item>
                <el-form-item label="Showdown队伍文本(可选)" :label-width="formLabelWidth" prop="showdown">
                    <el-input
                            type="textarea"
                            :rows="3"
                            placeholder="请将Pokemon Showdown队伍文本粘贴至此处"
                            v-model="form.showdown"
                            maxlength="1600">
                    </el-input>
                </el-form-item>
                <el-form-item label="描述（可选）" :label-width="formLabelWidth" prop="introduction">
                    <el-input
                            type="textarea"
                            :rows="3"
                            placeholder="可以在此处简要描述该队伍（不超过200字，过长可附上外部链接）"
                            v-model="form.introduction"
                            maxlength="200"
                            show-word-limit>
                    </el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="closeForm">取 消</el-button>
                <el-button type="primary" @click="submitDialogVisible = true">确 定</el-button>
                <el-dialog title="确认提交" :visible.sync="submitDialogVisible" width="56%" append-to-body>
                    <el-alert
                            title="是否确认提交？"
                            type="warning"
                            description="请对将要提交的内容加以确认。如果输入了Showdown队伍文本，下面会显示相应的队伍预览图。如果没有上传队伍租借ID图片，这张图将作为代替显示在我们的网站上，否则会使用上传的租借ID图。"
                            :closable="false"
                            show-icon>
                    </el-alert>
                    <showdown2img :pokemonlist="parsedShowdown"></showdown2img>
                    <span slot="footer" class="dialog-footer">
                        <el-button @click="submitDialogVisible = false">取 消</el-button>
                        <el-button type="primary" @click="submitForm">确 定</el-button>
                    </span>
                </el-dialog>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import {Formats} from "../assets/data/formats";
    import {Koffing} from 'koffing';
    import showdown2img from "./Showdown2Img"
    import html2canvas from 'html2canvas';

    const PokemonNames = require('../assets/data/pokemonNames.js');

    export default {
        name: "teamform",
        props: {
            dialogformvisible: {
                type: Boolean,
                default: false
            },
            url: {
                type: String,
                default: "http://127.0.0.1:8888/"
            },
        },
        data() {
            return {
                // data to display
                formats: [],
                pokemonNames: [],
                // data to submit
                form: {
                    title: '',
                    author: '',
                    format: '', // Array, need to pop() before submit
                    pokemon1: '', // split and pick the first item
                    pokemon2: '',
                    pokemon3: '',
                    pokemon4: '',
                    pokemon5: '',
                    pokemon6: '',
                    rentalImgUrl: '',
                    showdown: '',
                    introduction: '',
                    state: 0
                },
                // form attributes
                formLabelWidth: '150px',
                loginFormRules: {
                    title: [
                        {required: true, message: '请输入标题', trigger: 'blur'},
                        {min: 1, max: 20, message: '长度在 1 到 20 个字符', trigger: 'blur'}
                    ],
                    author: [
                        {required: true, message: '请输入作者名', trigger: 'blur'},
                        {min: 1, max: 15, message: '长度在 1 到 15 个字符', trigger: 'blur'}
                    ],
                    format: [
                        {required: true, message: '请选择对战模式', trigger: 'change'}
                    ],
                    pokemon1: [
                        {required: true, message: '请选择至少一只宝可梦', trigger: 'change'}
                    ],
                    showdown: [
                        {min: 0, max: 1600, message: 'Showdown 文本过长', trigger: 'blur'}
                    ],
                    introduction: [
                        {min: 0, max: 200, message: '长度在 200 个字符以内', trigger: 'blur'}
                    ]
                },
                // Upload image dialog properties
                dialogImageUrl: '',
                dialogVisible: false,
                // submit dialog
                submitDialogVisible: false,
                // generated image
                b64code: 10086
            }
        },
        created() {
            this.formats = Formats;
            this.pokemonNames = [
                PokemonNames.all('en'),
                PokemonNames.all('ja'),
                PokemonNames.all('ko'),
                PokemonNames.all('zh-Hans'),
                PokemonNames.all('zh-Hant')
            ].reduce((r, a) =>
                a.map((v, i) => (r[i] || []).concat(v)), []
            ).map(e => e.join('/'));
        },
        methods: {
            // generate image and post
            async screenshot() {
                let that = this;
                await html2canvas(document.getElementById('preview'), {
                    useCORS: true
                }).then(async canvas => {
                    const res = await this.$http.post('uploadb64', {base64: canvas.toDataURL()});
                    that.b64code = res.data.code;
                    if (res.data.code !== 200) {
                        that.form.rentalImgUrl = ''
                    } else {
                        that.form.rentalImgUrl = res.data.data
                    }
                });
            },
            // for the form
            resetForm() {
                this.$refs.teamFormRef.resetFields();
            },
            closeForm() {
                this.dialogformvisible = false;
                this.$emit("update:dialogformvisible", this.dialogformvisible)
            },
            async submitForm() {
                this.submitDialogVisible = false;
                if (this.form.rentalImgUrl === '' && this.form.showdown.length < 200) {
                    this.$message.error("租借队伍ID图片和Showdown队伍文本至少需要提交一个");
                    return;
                }
                if (this.form.rentalImgUrl === '') {
                    await this.screenshot();
                    if (this.b64code !== 200) {
                        this.$message.error("生成的队伍预览图提交出错");
                        return;
                    }
                }

                await this.$refs.teamFormRef.validate(async valid => {
                    if (!valid) return;

                    // process this form
                    function processPokemon(pname) {
                        return (pname === '') ? '' : pname.split('/', 1)[0]
                    }

                    this.form.pokemon1 = processPokemon(this.form.pokemon1);
                    this.form.pokemon2 = processPokemon(this.form.pokemon2);
                    this.form.pokemon3 = processPokemon(this.form.pokemon3);
                    this.form.pokemon4 = processPokemon(this.form.pokemon4);
                    this.form.pokemon5 = processPokemon(this.form.pokemon5);
                    this.form.pokemon6 = processPokemon(this.form.pokemon6);
                    this.form.format = this.form.format[this.form.format.length - 1];
                    // post
                    const res = await this.$http.post('teams', this.form);
                    if (res.data.code !== 200) {
                        // FIXME: `res.data.data` is an object
                        return this.$message.error(JSON.stringify(res.data.data))
                    }
                    this.$message.success('提交成功！请耐心等待审核');
                });
                // reset the form and close this dialog
                this.dialogformvisible = false;
                this.$emit("update:dialogformvisible", this.dialogformvisible);
            },
            // methods for image-upload dialog
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
                this.form.rentalImgUrl = ''
            },
            handlePictureCardPreview(file) {
                this.dialogImageUrl = file.url;
                this.dialogVisible = true;
            },
            handleSuccess(response, file, fileList) {
                if (response.code !== 200) {
                    this.$message.error('上传失败：' + response.data);
                    return
                }
                this.form.rentalImgUrl = response.data;
            }
        },
        computed: {
            uploadUrl() {
                return this.url + "api/v1/upload"
            },
            parsedShowdown() {
                return Koffing.parse(this.form.showdown).teams[0].pokemon;
            }
        },
        components: {
            showdown2img
        }
    }
</script>

<style scoped>

</style>