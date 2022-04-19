package main

import (
	"log"
	"os"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/bot/extras/middlewares"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/joho/godotenv"
)

type Bot struct {
	Ctx *bot.Context
}

func (bot *Bot) Setup(sub *bot.Subcommand) {
	sub.AddMiddleware(bot.Test, middlewares.GuildOnly(bot.Ctx))
}

func (b *Bot) Test(m *gateway.MessageCreateEvent, a string) (string, error) {
	if a == "" {
		return "szto", nil
	}
	return a, nil
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	token := os.Getenv("TOKEN")

	cmds := &Bot{}
	bot.Run(token, cmds, func(c *bot.Context) error {
		c.HasPrefix = bot.NewPrefix(",")
		c.SilentUnknown.Command = true

		return nil
	})
}
