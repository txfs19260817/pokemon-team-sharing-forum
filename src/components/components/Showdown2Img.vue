<template>
    <div id="preview" class="team-outer" v-show="true">
        <ul>
            <li v-for="(pm, index) in pokemon" :id="'pm'+index">
                <div class="team-pokemon">
                    <div class="team-pokemon-info">
                        <div class="team-pokemon-info-basic">
                            <div class="team-pokemon-info-basic-pokemon-icon">
                                <img class="pokemon-icon" :alt="pm.name" :src="iconPath(pm.name, pmIconBaseUrl)"/>
                            </div>
                            <div class="team-pokemon-info-basic-item-icon">
                                <img class="item-icon" :alt="pm.item" :src="iconPath(pm.item, itemIconBaseUrl)"/>
                            </div>
                            <div class="team-pokemon-info-basic-type-icons">
                                <img v-for="type in pokedex[processStr(pm.name)].types"
                                     :src="iconPath(type, typeIconBaseUrl)" :alt="type"/>
                            </div>
                        </div>
                        <div class="team-pokemon-info-name">{{pm.name + ' Lv. ' + (pm.level ? pm.level : 100)}}</div>
                        <div class="team-pokemon-info-ability">{{pm.ability}}</div>
                        <div class="team-pokemon-info-item">{{(pm.item ? pm.item : '')}}</div>
                    </div>
                    <div class="team-pokemon-moves">
                        <div class="team-pokemon-move" v-for="m in pm.moves">
                            <img :alt="m" :src="iconPath(m, movetypeIconBaseUrl)"/>
                            <div class="move-name">{{m}}</div>
                        </div>
                    </div>
                </div>
            </li>
        </ul>
    </div>
</template>

<script>
    import {BattlePokedex} from "../../assets/data/pokedex"
    import {BattleMovedex} from "../../assets/data/moves"

    export default {
        name: "showdown2img",
        props: {
            pokemonlist: {
                type: Array
            }
        },
        data() {
            return {
                pokedex: {},
                movedex: {},
                pokemon: this.pokemonlist,
                pmIconBaseUrl: process.env.VUE_APP_ICONS + '2d/',
                itemIconBaseUrl: process.env.VUE_APP_ICONS + 'items/',
                typeIconBaseUrl: process.env.VUE_APP_ICONS + 'types/',
                movetypeIconBaseUrl: process.env.VUE_APP_ICONS + 'movetypes/',
                ext: '.png',
            }
        },
        watch: {
            pokemonlist() {
                this.pokemon = this.pokemonlist;
            }
        },
        methods: {
            processStr(s) {
                return s.replace(/[-\s\[\].']/g, '').toLowerCase();
            },
            iconPath(name, baseUrl) {
                name = this.processStr(name);
                if (baseUrl.substr(baseUrl.length - 10) === "movetypes/") {
                    name = this.movedex[name].type.toLowerCase();
                    return baseUrl + name + this.ext
                }
                if (baseUrl.substr(baseUrl.length - 6) === "types/") {
                    let lang = "eng";
                    let filename = ["type", name, lang];
                    return baseUrl + filename.join('-') + this.ext
                }
                return baseUrl + name + this.ext
            }
        },
        created() {
            this.pokedex = BattlePokedex;
            this.movedex = BattleMovedex;
        }
    }
</script>

<style scoped>
    @import url('https://fonts.lug.ustc.edu.cn/css?family=IBM+Plex+Sans&display=swap');

    .team-outer {
        margin: 0;
        background: url(../../../public/background.jpg) no-repeat center;
        background-size: 100%;
        width: 1024px;
        height: 576px;
        text-align: center;
        font-family: 'IBM Plex Sans', sans-serif;
    }

    .team-outer ul {
        padding: 0;
        overflow: hidden;
    }

    .team-outer ul li {
        width: 46%;
        height: 144px;
        float: left;
        margin: 10px 20px;
        box-sizing: border-box;
        list-style: none
    }

    #pm0, #pm2, #pm4 {
        padding-left: 43px;
    }

    #pm1, #pm3, #pm5 {
        padding-right: 43px;
    }

    .team-outer ul li .team-pokemon {
        height: inherit;
        background: rgb(0, 174, 181);
        display: flex;
        flex-shrink: 0;
    }

    .team-outer ul li .team-pokemon .team-pokemon-info {
        flex: 0 0 50%;
        overflow: hidden;
        display: flex;
        flex-direction: column;
    }

    .team-outer ul li .team-pokemon .team-pokemon-info .team-pokemon-info-basic {
        flex: 0 0 40%;
        overflow: hidden;
        display: flex;
    }

    .team-outer ul li .team-pokemon .team-pokemon-info .team-pokemon-info-basic .team-pokemon-info-basic-pokemon-icon {
        flex: 0 0 25%;
        position: relative;
    }

    .pokemon-icon {
        position: absolute;
        right: 0%;
        bottom: 10%;
        /* width: 70%; */
    }

    .team-outer ul li .team-pokemon .team-pokemon-info .team-pokemon-info-basic .team-pokemon-info-basic-item-icon {
        flex: 0 0 15%;
        position: relative;
    }

    .item-icon {
        position: absolute;
        left: 0%;
        bottom: 10%;
    }

    .team-outer ul li .team-pokemon .team-pokemon-info .team-pokemon-info-basic .team-pokemon-info-basic-type-icons {
        flex: 0 0 60%;
        overflow: hidden;
    }

    .team-outer ul li .team-pokemon .team-pokemon-info .team-pokemon-info-basic .team-pokemon-info-basic-type-icons img {
        width: 85%;
    }

    .team-outer ul li .team-pokemon .team-pokemon-info .team-pokemon-info-name {
        flex: 0 0 20%;
        text-align: left;
        text-indent: 4px;
        color: #FFFFFF;
        font-size: 16px;
        font-weight: bold;
        /* 文字垂直居中 */
        align-items: center;
        display: -webkit-flex;
    }

    .team-outer ul li .team-pokemon .team-pokemon-info .team-pokemon-info-ability {
        flex: 0 0 20%;
        text-align: left;
        text-indent: 4px;
        color: #FFFFFF;
        font-size: 16px;
        font-weight: bold;
        /* 文字垂直居中 */
        align-items: center;
        display: -webkit-flex;
    }

    .team-outer ul li .team-pokemon .team-pokemon-info .team-pokemon-info-item {
        flex: 0 0 20%;
        text-align: left;
        text-indent: 4px;
        color: #FFFFFF;
        font-size: 16px;
        font-weight: bold;
        /* 文字垂直居中 */
        align-items: center;
        display: -webkit-flex;
    }

    .team-outer ul li .team-pokemon .team-pokemon-moves {
        flex: 0 0 50%;
        overflow: hidden;
        background-color: white;
        display: flex;
        flex-direction: column;
    }

    .team-outer ul li .team-pokemon .team-pokemon-moves .team-pokemon-move {
        flex: 0 0 25%;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: flex-start;
    }

    .team-outer ul li .team-pokemon .team-pokemon-moves .team-pokemon-move img {
        margin: 0 2px;
        width: 15%;
    }

    .move-name {
        font-size: 16px;
        font-weight: bold;
    }
</style>