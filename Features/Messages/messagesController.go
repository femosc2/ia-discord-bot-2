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
	} else if m.Content == "!iahelp" {
		s.ChannelMessageSend(m.ChannelID,
			`
			_________________
			!contribute - Länk till repot för den här botten.
			!exams<17/18/19> - Se alla salstentor på ert schema - e.g !exams19
			!quote - Få ett slumpmässigt quote från en IAare!
			!addquote - Lägg till ett nytt quote! - e.g !addquote Felix nya quote
			-Felix (meddelandet måste sluta med ett bindelsträck följt av vem som sa det.)
			`)
	} else if m.Content == "!quote" {
		s.ChannelMessageSend(m.ChannelID, quotes.GetQuote())
	} else if strings.HasPrefix(m.Content, "!addquote ") {
		rawQuote := strings.Replace(m.Content, "!addquote ", "", 1)
		author := strings.SplitAfter(rawQuote, "-")
		quote := strings.Trim(rawQuote, "-"+author[1])
		s.ChannelMessageSend(m.ChannelID, quotes.PostQuote(quote, m.Author.Username, author[1]))
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
