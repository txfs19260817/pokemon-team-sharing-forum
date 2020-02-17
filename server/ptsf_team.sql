/*
 Navicat Premium Data Transfer

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : pokeshare

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 17/02/2020 22:32:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ptsf_team
-- ----------------------------
DROP TABLE IF EXISTS `ptsf_team`;
CREATE TABLE `ptsf_team`  (
  `id` int unsigned NOT NULL,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `author` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `created_at` datetime(0) DEFAULT NULL,
  `format` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `pokemon1` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `pokemon2` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `pokemon3` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `pokemon4` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `pokemon5` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `pokemon6` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `rental_img_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `showdown` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `state` tinyint(0) DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ptsf_team
-- ----------------------------
INSERT INTO `ptsf_team` VALUES (1, '扬州队', '江主席', '2020-02-16 19:40:23', 'VGC 2020', 'Duraludon', 'Whimsicott', 'Dracovish', 'Togekiss', 'Arcanine', 'Inteleon', 'http://127.0.0.1:8888/assets/teams/20200216194023000T86ff888814548e1ebaf075ad999f837f.png', 'Whimsicott @ Focus Sash  \nAbility: Prankster  \nLevel: 50  \nHappiness: 160  \nEVs: 252 HP / 4 SpA / 252 Spe  \nTimid Nature  \nIVs: 0 Atk  \n- Tailwind  \n- Dazzling Gleam  \n- Helping Hand  \n- Fake Tears  \n\nDracovish @ Choice Scarf  \nAbility: Strong Jaw  \nLevel: 50  \nEVs: 252 Atk / 4 SpD / 252 Spe  \nAdamant Nature  \n- Crunch  \n- Dragon Rush  \n- Fishious Rend  \n- Psychic Fangs  \n\nDuraludon @ Assault Vest  \nAbility: Stalwart  \nLevel: 50  \nEVs: 244 HP / 252 SpA / 12 Spe  \nModest Nature  \nIVs: 0 Atk  \n- Flash Cannon  \n- Dark Pulse  \n- Draco Meteor  \n- Thunderbolt  \n\npozor (Togekiss) @ Scope Lens  \nAbility: Super Luck  \nLevel: 50  \nEVs: 252 HP / 44 Def / 116 SpA / 4 SpD / 92 Spe  \nModest Nature  \nIVs: 0 Atk  \n- Air Slash  \n- Dazzling Gleam  \n- Follow Me  \n- Heat Wave  \n\nInteleon @ Life Orb  \nAbility: Torrent  \nLevel: 50  \nEVs: 76 HP / 4 Def / 252 SpA / 4 SpD / 172 Spe  \nTimid Nature  \nIVs: 0 Atk  \n- Air Slash  \n- Mud Shot  \n- Ice Beam  \n- Snipe Shot  \n\npes filipes (Arcanine) @ Mago Berry  \nAbility: Intimidate  \nLevel: 50  \nEVs: 252 HP / 28 Def / 4 SpA / 4 SpD / 220 Spe  \nTimid Nature  \nIVs: 0 Atk  \n- Snarl  \n- Flamethrower  \n- Will-O-Wisp  \n- Roar  \n\n', 'PS 1700', 1);
INSERT INTO `ptsf_team` VALUES (39, 'WeeZingZap', 'Labib Haq', '2020-02-17 20:23:25', 'VGC 2020', 'Weezing', 'Duraludon', 'Excadrill', 'Gyarados', 'Togedemaru', 'Braviary', 'http://127.0.0.1:8888/assets/teams/20200217202248000T03dee440a8dd1e7cacd50a32ff067566.jpg', '\n\nGyarados @ Assault Vest  \nAbility: Intimidate  \nLevel: 50  \nShiny: Yes  \nEVs: 140 HP / 196 Atk / 172 Spe  \nAdamant Nature  \n- Waterfall  \n- Bounce  \n- Power Whip  \n- Earthquake  \n\nTogedemaru @ Focus Sash  \nAbility: Lightning Rod  \nLevel: 50  \nEVs: 4 HP / 252 Atk / 252 Spe  \nJolly Nature  \n- Fake Out  \n- Zing Zap  \n- Nuzzle  \n- Spiky Shield  \n\nExcadrill @ Choice Scarf  \nAbility: Mold Breaker  \nLevel: 50  \nEVs: 4 HP / 252 Atk / 252 Spe  \nJolly Nature  \n- Earthquake  \n- Rock Slide  \n- High Horsepower  \n- Iron Head  \n\nBraviary (M) @ Sitrus Berry  \nAbility: Defiant  \nLevel: 50  \nEVs: 156 HP / 100 Atk / 252 Spe  \nJolly Nature  \n- Brave Bird  \n- Close Combat  \n- Tailwind  \n- Protect  \n\nWeezing-Galar @ Air Balloon  \nAbility: Neutralizing Gas  \nLevel: 50  \nEVs: 236 HP / 20 Def / 4 SpA / 236 SpD / 12 Spe  \nBold Nature  \nIVs: 0 Atk  \n- Clear Smog  \n- Taunt  \n- Protect  \n- Strange Steam  \n\nDuraludon @ Weakness Policy  \nAbility: Stalwart  \nLevel: 50  \nShiny: Yes  \nEVs: 252 HP / 204 SpA / 52 Spe  \nModest Nature  \nIVs: 0 Atk  \n- Dark Pulse  \n- Draco Meteor  \n- Flash Cannon  \n- Protect  \n\n', 'Labib Haq (@slippingbug) https://victoryroadvgc.com/2020/01/27/weezingzap-dallas-report/', 1);

SET FOREIGN_KEY_CHECKS = 1;
