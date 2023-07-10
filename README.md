# Anime Schedule Bot

## Operation principle

- You send the bot a link to an anime (mandatory from the animego.org website).
- The bot mentions the release time of the episode (in Japan).
- The bot checks every hour, according to the cron schedule, if a new episode of the anime has been released (with voice acting or subtitles).

## Local setup

- make env file (see .env.example)
- Perform the database migration.

```sh
npx prisma migrate dev
```

- Run 

```sh
yarn dev
```