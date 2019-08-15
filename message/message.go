package message

import "github.com/bwmarrin/discordgo"

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!help" {
		s.ChannelMessageSend(m.ChannelID, "Detta kommandot kan man använda för att se alla tillgängliga kommando så fort jag lägger till ett par")
	}

	switch m.Content {
	case "!contribute":
		s.ChannelMessageSend(m.ChannelID, "https://github.com/femosc2/ia-discord-bot-2/")
	case "!help":
		s.ChannelMessageSend(m.ChannelID, "Detta kommandot kan man använda för att se alla tillgängliga kommando så fort jag lägger till ett par")
	}
}
