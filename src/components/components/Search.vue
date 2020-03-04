<template>
    <span>
        <!--        search by format-->
        <el-cascader v-if="by==='format'"
                     v-model="selectedFormat"
                     :options="formats"
                     :placeholder="$t('form.format')"
                     :show-all-levels="false"
                     @change="searchFormat">
        </el-cascader>
        <!--        search by pokemon name-->
        <el-select v-if="by==='pokemon'"
                   v-model="selectedPokemon"
                   filterable
                   :placeholder="$t('pokemon')"
                   @change="searchPokemon">
            <el-option v-for="item in pokemonNames" :key="item" :label="item" :value="item"></el-option>
        </el-select>
    </span>
</template>

<script>
    import {Formats} from "../../assets/data/formats";

    const PokemonNames = require('../../assets/data/pokemonNames.js');

    export default {
        name: "Search",
        props: {
            by: {
                type: String,
                default: 'pokemon'
            },
        },
        data() {
            return {
                // search panel
                formats: [],
                pokemonNames: [],

                selectedFormat: null,
                selectedPokemon: null,
            }
        },
        watch: {
            '$route' (to, from) {
                this.$router.go(0);
            }
        },
        methods: {
            searchFormat(f) {
                f = f[1];
                this.$router.push('/formats/' + f)
            },
            searchPokemon(p) {
                p = p.split('/', 1)[0];
                this.$router.push('/pokemon/' + p)
            },
        },
        created() {
            this.formats = Formats;
            this.pokemonNames = PokemonNames.pmNames4Select;
        }
    }
</script>

<style scoped>

</style>