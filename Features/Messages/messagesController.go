package messages

import (
	"github.com/bwmarrin/discordgo"
	quotes "github.com/femosc2/ia-discord-bot-2/Features/Quotes"
)

// MessageCreate The bot sends a message in the discord channel
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "!contribute":
		s.ChannelMessageSend(m.ChannelID, "https://github.com/femosc2/ia-discord-bot-2/")
	case "!help":
		s.ChannelMessageSend(m.ChannelID, "Detta kommandot kan man använda för att se alla tillgängliga kommando så fort jag lägger till ett par")
	case "!quote":
		s.ChannelMessageSend(m.ChannelID, quotes.GetQuote())
	}
}
