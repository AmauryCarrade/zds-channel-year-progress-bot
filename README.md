# #covid-19⸴87➔21⸴98

This ~~simple~~ stupid bot updates a Discord channel name to match the current year progress.

Build with `go build`, then:

```
./discord-channel-year-bot -t TOKEN -c CHANNEL_ID
```

where `TOKEN` is the bot token, and CHANNEL_ID is the channel to update. The channel name format
is not (yet?) configurable: it's always `covid-{begin}➔{end}`, as in the title, where `{begin}` is
the date of the beginning of the COVID-19 pandemic.

Just, don't ask. (This is also a toy project to learn Go, Docker, and some k8s integration.)
