package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main () {
	token := "test" 
	dg , err := discordgo.New("BOT" + token) 
	if err != nil {
		fmt.Println("err : %v", err)
		return
	}
	fmt.Println("Bot is now running. Press Ctrl+C to exit.");
}