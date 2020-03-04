<template>
    <div id="teamform">
        <el-dialog :title="$t('form.form')" :visible.sync="dialogformvisible" :before-close="closeForm" width="70%"
                   :modal-append-to-body="false" append-to-body>
            <el-form :model="form" ref="form" :rules="loginFormRules">
                <el-form-item :label="$t('form.title')" :label-width="formLabelWidth" prop="title">
                    <el-input v-model="form.title" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item :label="$t('form.author')" :label-width="formLabelWidth" prop="author">
                    <el-input v-model="form.author" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item :label="$t('form.format')" :label-width="formLabelWidth" prop="format">
                    <el-cascader
                            v-model="form.format"
                            :options="formats"
                            :placeholder="$t('form.placeholder.format')"
                            :show-all-levels="false">
                    </el-cascader>
                </el-form-item>
                <el-form-item :label="$t('form.pokemon1')" :label-width="formLabelWidth" prop="pokemon1">
                    <el-select v-model="form.pokemon1" filterable :placeholder="$t('form.placeholder.pokemon')">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item :label="$t('form.pokemon2')" :label-width="formLabelWidth" prop="pokemon2">
                    <el-select v-model="form.pokemon2" filterable :placeholder="$t('form.placeholder.pokemon')">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item :label="$t('form.pokemon3')" :label-width="formLabelWidth" prop="pokemon3">
                    <el-select v-model="form.pokemon3" filterable :placeholder="$t('form.placeholder.pokemon')">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item :label="$t('form.pokemon4')" :label-width="formLabelWidth" prop="pokemon4">
                    <el-select v-model="form.pokemon4" filterable :placeholder="$t('form.placeholder.pokemon')">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item :label="$t('form.pokemon5')" :label-width="formLabelWidth" prop="pokemon5">
                    <el-select v-model="form.pokemon5" filterable :placeholder="$t('form.placeholder.pokemon')">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item :label="$t('form.pokemon6')" :label-width="formLabelWidth" prop="pokemon6">
                    <el-select v-model="form.pokemon6" filterable :placeholder="$t('form.placeholder.pokemon')">
                        <el-option
                                v-for="item in pokemonNames"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item
                        :label="$t('form.rental') + ' (' + $t('form.placeholder.optional') + ', ' + $t('form.placeholder.uploadFormat') + ')'"
                        :label-width="formLabelWidth" prop="rentalImgUrl">
                    <upload :imgurl.sync="form.rentalImgUrl"></upload>
                </el-form-item>
                <el-form-item :label="$t('form.showdown') + ' (' + $t('form.placeholder.optional') + ')'"
                              :label-width="formLabelWidth" prop="showdown">
                    <el-input
                            type="textarea"
                            :rows="3"
                            :placeholder="$t('form.placeholder.showdown')"
                            v-model="form.showdown"
                            maxlength="1600">
                    </el-input>
                </el-form-item>
                <el-form-item :label="$t('form.description') + ' (' + $t('form.placeholder.optional') + ')'"
                              :label-width="formLabelWidth" prop="description">
                    <el-input
                            type="textarea"
                            :rows="3"
                            :placeholder="$t('form.placeholder.description')"
                            v-model="form.description"
                            maxlength="300"
                            show-word-limit>
                    </el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" icon="el-icon-circle-close" circle @click="closeForm"></el-button>
                <el-button type="primary" icon="el-icon-circle-check" circle
                           @click="submitDialogVisible = true"></el-button>
                <!--                Confirm dialog-->
                <el-dialog :title="$t('form.alert.dialogTitle')" :visible.sync="submitDialogVisible" width="56%"
                           append-to-body :modal="false">
                    <el-alert
                            :title="$t('form.alert.alertTitle')"
                            type="warning"
                            :description="$t('form.alert.description')"
                            :closable="false"
                            show-icon>
                    </el-alert>
                    <showdown2img :pokemonlist="parsedShowdown"></showdown2img>
                    <span slot="footer" class="dialog-footer">
                        <el-button type="info" icon="el-icon-circle-close" circle
                                   @click="submitDialogVisible = false"></el-button>
                        <el-button type="primary" icon="el-icon-circle-check" circle @click="submitForm"></el-button>
                    </span>
                </el-dialog>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import {Formats} from "../../assets/data/formats";
    import {Koffing} from 'koffing';
    import showdown2img from "./Showdown2Img"
    import upload from "./Upload";
    import html2canvas from 'html2canvas';

    const PokemonNames = require('../../assets/data/pokemonNames.js');

    export default {
        name: "teamform",
        props: {
            dialogformvisible: {
                type: Boolean,
                default: false
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
                    description: '',
                    state: 0
                },
                // form attributes
                formLabelWidth: '150px',
                loginFormRules: {
                    title: [
                        {required: true, message: '请输入标题', trigger: 'blur'},
                        {min: 1, max: 50, message: '长度在 1 到 50 个字符', trigger: 'blur'}
                    ],
                    author: [
                        {required: true, message: '请输入作者名', trigger: 'blur'},
                        {min: 1, max: 30, message: '长度在 1 到 30 个字符', trigger: 'blur'}
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
                    description: [
                        {min: 0, max: 1000, message: '长度在 1000 个字符以内', trigger: 'blur'}
                    ]
                },
                // submit dialog
                submitDialogVisible: false,
                // generated team preview image: return code
                b64code: 10086
            }
        },
        created() {
            this.formats = Formats;
            this.pokemonNames = PokemonNames.pmNames4Select;
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
                this.$refs.form.resetFields();
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

                await this.$refs.form.validate(async valid => {
                    if (!valid) {
                        this.$message.error('提交内容有误，请修改')
                        return;
                    }
                    this.$message.warning('提交中，请耐心等待不要重复提交');

                    // process form data
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
        },
        computed: {
            parsedShowdown() {
                return Koffing.parse(this.form.showdown).teams[0].pokemon;
            }
        },
        components: {
            upload,
            showdown2img
        }
    }
</script>

<style scoped>

</style>