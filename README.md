# #covid-19⸴87➔21⸴98

This ~~simple~~ stupid bot updates a Discord channel name to match the current year progress.

Build with `go build`, then:

```
DISCORD_TOKEN=TOKEN DISCORD_CHANNEL_ID=CHANNEL_ID ./discord-channel-year-bot
```

The following environment variables are available (can be provided in a `.env` file too):
- `DISCORD_TOKEN`: the Discord bot token (required);
- `DISCORD_CHANNEL_ID`: the channel ID to update (required);
- `DISCORD_CHANNEL_PATTERN`: the channel name pattern, where `{begin}` is replaced by the beginning date year
  percentage, and `{end}`, by the end date year percentage (default `covid-{begin}➔{end}`).

The beginning date is the one of the beginning of the COVID-19 pandemic.

_Just, don't ask. (This is also a toy project to learn Go, Docker, and some k8s integration.)_
