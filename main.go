package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	// Don't care about error as if it doesn't load it's fine as long as regular env vars are used.
	err := godotenv.Load(".env")

	token := os.Getenv("DISCORD_TOKEN")
	channelId := os.Getenv("DISCORD_CHANNEL_ID")
	channelPattern := os.Getenv("DISCORD_CHANNEL_PATTERN")

	if channelPattern == "" {
		channelPattern = "covid-{begin}➔{end}"
	}

	if token == "" || channelId == "" {
		log.Fatal("Both DISCORD_TOKEN and DISCORD_CHANNEL_ID environment variables are required (.env supported).")
	}

	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		fmt.Println("Unable to load location", err)
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error while creating Discord session", err)
		return
	}

	channel, err := dg.Channel(channelId)
	if err != nil {
		fmt.Println("Error while loading channel", err)
		return
	}

	edit := discordgo.ChannelEdit{
		Name: GenerateChannelName(
			channelPattern,
			time.Date(2019, time.November, 16, 0, 0, 0, 0, loc),
			time.Now().In(loc),
		),
		// We must specify the position, and cannot use the dg.ChannelEdit shortcut method,
		// else the channel is moved at the top every time.
		Position: channel.Position,
	}

	_, err = dg.ChannelEditComplex(channelId, &edit)
	if err != nil {
		fmt.Println("Error while editing channel", err)
		return
	}
}
