# pokemon-team-sharing-forum
PokéShare is a web platform that allows trainers to share awesome teams with others around the world. 

## Features
- RESTful API backend written in Go
- Vue.js-based frontend
- Entries: [PC](https://pokeshare.monster), [Mobile](https://m.pokeshare.monster), [Repo of the Mobile entry](https://github.com/txfs19260817/pokemon-team-sharing-forum-mobile)
- Convert text teams in Showdown format to an image that looks like the Rental Team Code in Pokémon Sword and Shield

![Convert](https://github.com/txfs19260817/pokemon-team-sharing-forum/blob/master/docs/convert.jpg)
- Search teams that contain specific Pokémon or under certain battle format

![Search](https://github.com/txfs19260817/pokemon-team-sharing-forum/blob/master/docs/search.jpg)
- Check the Pokémon usages rates under specific battle format

![Usage Rates](https://github.com/txfs19260817/pokemon-team-sharing-forum/blob/master/docs/usage.jpg)

## Mobile Devices Entry

https://m.pokeshare.monster/

https://github.com/txfs19260817/pokemon-team-sharing-forum-mobile

## Project setup
## Backend
Change directory to `server` and import the .sql file. Then, configure `app.ini` and run
```
go run main.go
```

## Frontend
```
npm install
```

### Compiles and hot-reloads for development
Setup `.env` file, then
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

## TODOs
- [ ] i18n
- [x] filter results by formats and pokemon
- [x] About Page
