package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	google "github.com/rocketlaunchr/google-search"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"context"
)


func main() {
  Token := os.Getenv("TOKEN")
  dg, err := discordgo.New("Bot " + Token)
  if err != nil {
  	return
  }
  go dg.AddHandler(message)
  
  dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
  fmt.Println("Online")
  err = dg.Open()
  if err != nil{
  	return
  }
  sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// close session.
	dg.Close()
}

func message(client *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.Bot {
		return
	}
	prefix := "g!"
	if strings.HasPrefix(m.Content, prefix) {
		Content := strings.TrimLeft(m.Content, prefix)
		args := strings.Split(Content, " ")[1:]
		cmd := strings.Split(Content, " ")[0]
		
		if cmd == "search" {
			if len(args) < 1 {
				client.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
					Title: "📥 • **Please mention something!**",
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Google - Shockalicious#6576",
					},
				})
				return
			}
			ctx := context.Background()
			result, err := google.Search(ctx, strings.Join(args[0:], " "))
			if err != nil {
				fmt.Println("Search error")
				return 
			}
			var res string
			for i := 0; i < len(result); i++{
				res += fmt.Sprintf(" • [%s](%s)\n", result[i].Title, result[i].URL)
			}
			client.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title: "📬 • **Your Search Results**",
				Description: res,
				Footer: &discordgo.MessageEmbedFooter{
					Text: "Google - Shockalicious#6576",
				},
			})
		}
	}
}
