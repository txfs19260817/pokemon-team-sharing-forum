<template>
    <div id="teamform">
        <el-dialog title="队伍表格" :visible.sync="dialogformvisible" :before-close="closeForm">
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
                            list-type="picture-card" h
                            :limit=1
                            :before-upload="handleBeforeUpload"
                            :on-exceed="handleExceed"
                            accept="image/jpeg,image/jpg,image/png"
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
                            v-model.lazy="form.showdown"
                            maxlength="1600">
                    </el-input>
                </el-form-item>
                <el-form-item label="简介（可选）" :label-width="formLabelWidth" prop="introduction">
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
                <el-button type="primary" @click="submitForm">确 定</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import {Formats} from "../assets/data/formats";
    import {Koffing} from 'koffing';

    const PokemonNames = require('../assets/data/pokemonNames.js');

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
                url: "http://127.0.0.1:8888/",
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
                        {min: 1, max: 10, message: '长度在 1 到 20 个字符', trigger: 'blur'}
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
            // for this form dialog
            resetForm() {
                this.$refs.teamFormRef.resetFields();
            },
            closeForm() {
                this.dialogformvisible = false;
                this.$emit("update:dialogformvisible", this.dialogformvisible)
            },
            submitForm() {
                this.$refs.teamFormRef.validate(async valid => {
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
                    this.form.format = this.form.format[0];
                    // post
                    const res = await this.$http.post('teams', this.form);
                    console.log(res);
                    if (res.data.code !== 200) {
                        // FIXME: `res.data.data` is an object
                        return this.$message.error(res.data.data)
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
            handlePictureCardPreview(file) {
                this.dialogImageUrl = file.url;
                this.dialogVisible = true;
                console.log(this.dialogImageUrl)
            },
            handleSuccess(response, file, fileList) {
                this.form.rentalImgUrl = response.data;
                console.log(this.form.rentalImgUrl);
            }
        },
        computed: {
            uploadUrl() {
                return this.url + "api/v1/upload"
            },
            // TODO: rename uploaded image
            // TODO: parse showdown -> generate css -> css to img
            parseShowdown() {
                let parsedTeam = Koffing.parse(this.form.showdown);
                return parsedTeam.teams[0].pokemon
            },
        },
    }
</script>

<style scoped>

</style>