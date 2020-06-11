# pokemon-team-sharing-forum
PokéShare is a web platform that allows trainers to share awesome teams with others around the world. 

## Features
- RESTful API backend written in Go
- Vue.js-based frontend
- Convert text teams in Showdown format to an image that looks like the Rental Team Code in Pokémon Sword and Shield
- Search teams that contain specific Pokémon or under certain battle format
- Check the usage rate under specific battle format

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
