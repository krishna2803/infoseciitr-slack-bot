package commands

import (
	"infoseciitr/slack-bot/pkg/log"
	"log/slog"

	"github.com/slack-io/slacker"
)

func HandlePing() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description: "Pings the bot",
		Command:     "ping",
		Examples:    []string{"bot ping"},
		Handler: func(ctx *slacker.CommandContext) {
			log.GetLogger().Info("Received", slog.String("command", "ping"))

			_, err := ctx.Response().Reply("pong :table_tennis_paddle_and_ball:")

			if err != nil {
				log.GetLogger().Error("Error in HandlePing", slog.String("error", err.Error()))
			}

			log.GetLogger().Info("Replied", slog.String("message", "pong"))
		},
	}
}
