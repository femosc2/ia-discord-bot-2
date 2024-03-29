package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	extensions "github.com/femosc2/ia-discord-bot-2/Extensions"
	entirestore "github.com/femosc2/ia-discord-bot-2/Store"
	viewmodels "github.com/femosc2/ia-discord-bot-2/viewmodels"

	config "github.com/femosc2/ia-discord-bot-2/config"
	"github.com/femosc2/ia-discord-bot-2/features/messages"

	"github.com/bwmarrin/discordgo"
)

func main() {

	fmt.Println("Adding items to store...")
	store := viewmodels.StoreViewModel{
		Schedules: extensions.StartupScrape(),
	}
	entirestore.SetStore(store)
	fmt.Println("Items added to store!")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messages.MessageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
