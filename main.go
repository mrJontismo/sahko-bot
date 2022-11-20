package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"

	"sahko-bot/price"

	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	Token := os.Getenv("TOKEN")

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal("Error creating discord session: ", err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	prefix := "!sahko"
	prices := price.GetPrice()

	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case prefix + " nyt":
		s.ChannelMessageSend(m.ChannelID, prices.Now + " snt/kWh")
	case prefix + " alin":
		s.ChannelMessageSend(m.ChannelID, prices.Min + " snt/kWh")
	case prefix + " ylin":
		s.ChannelMessageSend(m.ChannelID, prices.Max + " snt/kWh")
	case prefix + " vk":
		s.ChannelMessageSend(m.ChannelID, prices.Avg + " snt/kWh")
	case prefix + " kk":
		s.ChannelMessageSend(m.ChannelID, prices.Avg_28 + " snt/kWh")
	}
}