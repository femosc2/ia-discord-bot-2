package messages

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	quotes "github.com/femosc2/ia-discord-bot-2/Features/Quotes"
)

// MessageCreate The bot sends a message in the discord channel
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!contribute" {
		s.ChannelMessageSend(m.ChannelID, "https://github.com/femosc2/ia-discord-bot-2/")
	} else if m.Content == "!help" {
		s.ChannelMessageSend(m.ChannelID, "Detta kommandot kan man använda för att se alla tillgängliga kommando så fort jag lägger till ett par")
	} else if m.Content == "!quote" {
		s.ChannelMessageSend(m.ChannelID, quotes.GetQuote())
	} else if strings.HasPrefix(m.Content, "!addquote ") {
		quote := strings.Replace(m.Content, "!addquote ", "", 1)
		s.ChannelMessageSend(m.ChannelID, quotes.PostQuote(quote, "Felix", m.Author.Username))
	}
}
