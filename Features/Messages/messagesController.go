package messages

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	quotes "github.com/femosc2/ia-discord-bot-2/Features/Quotes"
	entirestore "github.com/femosc2/ia-discord-bot-2/Store"
)

// MessageCreate The bot sends a message in the discord channel
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	store := entirestore.GetStore()

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
		rawQuote := strings.Replace(m.Content, "!addquote ", "", 1)
		author := strings.SplitAfter(rawQuote, "-")
		quote := strings.Trim(rawQuote, "-"+author[1])
		s.ChannelMessageSend(m.ChannelID, quotes.PostQuote(quote, author[1], m.Author.Username))
	} else if m.Content == "!exams17" {
		for _, element := range store.Schedules.IA17 {
			s.ChannelMessageSend(m.ChannelID, "________________________")
			s.ChannelMessageSend(m.ChannelID, element)
		}
		s.ChannelMessageSend(m.ChannelID, "________________________")
	} else if m.Content == "!exams18" {
		for _, element := range store.Schedules.IA18 {
			s.ChannelMessageSend(m.ChannelID, "________________________")
			s.ChannelMessageSend(m.ChannelID, element)
		}
		s.ChannelMessageSend(m.ChannelID, "________________________")
	} else if m.Content == "!exams19" {
		for _, element := range store.Schedules.IA19 {
			s.ChannelMessageSend(m.ChannelID, "________________________")
			s.ChannelMessageSend(m.ChannelID, element)
		}
		s.ChannelMessageSend(m.ChannelID, "________________________")
	}
}
