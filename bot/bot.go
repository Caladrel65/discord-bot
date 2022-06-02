package bot

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"myapp/config"
	"os"
)

var (
	id      string
	session *discordgo.Session
)

func Start() {
	// Create a new bot session
	var err error
	session, err = discordgo.New("Bot " + config.Config.Token)
	if err != nil {
		log.WithError(err).Error("error creating discordgo session")
		return
	}

	// Grab bot as a user variable
	botUser, err := session.User("@me")
	if err != nil {
		log.WithError(err).Error("error retrieving bot user")
		return
	}

	// Store the bot's user ID
	id = botUser.ID

	// Add the handler function to this session
	session.AddHandler(messageHandler)

	err = session.Open()
	if err != nil {
		log.WithError(err).Error("error opening session to Discord")
		return
	}

	log.Info("Bot is running!")
}

// messageHandler listens for messages and responds.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Return immediately if the message is from this bot
	if m.Author.ID == id {
		return
	}

	// If my brother is sending a message, respond with Felix
	// TODO: Update to his ID not mine
	if m.Author.ID == "119165799284867072" {
		//if m.Author.ID == "461078091952029716" {
		file, err := os.Open("felix.gif")
		if err != nil {
			log.WithError(err).Error("error opening gif")
			return
		}
		_, err = s.ChannelFileSend(m.ChannelID, "felix.gif", file)
		if err != nil {
			log.WithError(err).Error("error sending gif")
			return
		}
		return
	}

	// If the message says "ping", send a message that says "pong"
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		return
	}
}
