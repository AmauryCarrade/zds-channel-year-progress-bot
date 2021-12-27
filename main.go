package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"time"
)

var (
	Token   string
	Channel string
)

func init() {
	flag.StringVar(&Token, "t", "", "Discord bot token")
	flag.StringVar(&Channel, "c", "", "Discord channel ID")
	flag.Parse()

	if Token == "" || Channel == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		fmt.Println("Unable to load location", err)
		return
	}

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error while creating Discord session", err)
		return
	}

	channel, err := dg.Channel(Channel)
	if err != nil {
		fmt.Println("Error while loading channel", err)
		return
	}

	edit := discordgo.ChannelEdit{
		Name: GenerateChannelName(
			"covid-{begin}➔{end}",
			time.Date(2019, time.November, 16, 0, 0, 0, 0, loc),
			time.Now().In(loc),
		),
		// We must specify the position, and cannot use the dg.ChannelEdit shortcut method,
		// else the channel is moved at the top every time.
		Position: channel.Position,
	}

	_, err = dg.ChannelEditComplex(Channel, &edit)
	if err != nil {
		fmt.Println("Error while editing channel", err)
		return
	}
}
