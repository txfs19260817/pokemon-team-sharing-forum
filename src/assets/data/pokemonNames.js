'use strict';
const pokemon = require('./pokemonNames/en.json');

const repositoryUrl = 'https://github.com/sindresorhus/pokemon';
const reportText = `Please report to ${repositoryUrl}/issues if we missed something.`;

const languages = new Set([
    'de',
    'en',
    'fr',
    'ja',
    'ko',
    'ru',
    'zh-Hans',
    'zh-Hant'
]);

function getLocalizedList(language = 'en') {
    if (language === 'en') {
        return pokemon;
    }

    if (!languages.has(language)) {
        throw new Error(`Localized list for language code '${language}' does not exist. Pull request welcome: ${repositoryUrl}`);
    }

    return require(`./pokemonNames/${language.toLowerCase()}.json`);
}

exports.all = getLocalizedList;

exports.getName = (id, language = 'en') => {
    const list = getLocalizedList(language);
    const name = list[id - 1];

    if (!name) {
        return '';
    }

    return name;
};

exports.getId = (name, language = 'en') => {
    const list = getLocalizedList(language);
    const index = list.indexOf(name);

    if (index === -1) {
        return -1;
    }

    return index + 1;
};

exports.languages = languages;
