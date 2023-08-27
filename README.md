# Anime Schedule Bot

## Operation principle

- You send the bot a link to an anime.
- The bot mentions the release time of the episode.
- The bot checks every half hour, according to the cron schedule, whether a new anime episode has been released (with voice acting or subtitles).

## Supported websites

### English-language sites

- [4anime.is](https://4anime.is/)

### Russian-language sites

- [animego.org](https://animego.org/)
- [amedia.online](https://amedia.online/)
- [animevost.org](https://animevost.org/)


![prewiew image](./prewiew.png)

## Local setup

- make env file (see .env.example)
- run script from Makefile

```sh
make dev 
```

## Production setup 

- run script from Makefile

```sh
make start
```

> it runs docker containers