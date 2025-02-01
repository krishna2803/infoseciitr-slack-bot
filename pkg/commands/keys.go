package commands

import (
	"fmt"
	"infoseciitr/slack-bot/pkg/blockkit"
	"infoseciitr/slack-bot/pkg/log"
	"infoseciitr/slack-bot/services"
	"log/slog"
	"strings"

	"github.com/slack-io/slacker"
)

func HandleWhoHasTheKeys() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description: "Fetches the current key owners",
		Examples:    []string{"bot who has the keys"},
		Command:     "who has the keys",
		Handler: func(ctx *slacker.CommandContext) {
			log.GetLogger().Info("Received", slog.String("command", "who has the keys"))

			keys := services.WhoHasTheKeys()

			attachments := blockkit.FormatKeys(keys)
			_, err := ctx.Response().ReplyBlocks(attachments)

			if err != nil {
				log.GetLogger().Error("Error in HandleWhoHasTheKeys", slog.String("error", err.Error()))
			}

			var msg string
			if len(keys) == 0 {
				msg = "No keys found"
			} else {
				for _, key := range keys {
					msg += fmt.Sprintf("%s has the %s keys\n", key.Owner, key.Name)
				}
			}

			log.GetLogger().Info("Replied", slog.String("message", msg))
		},
	}
}

func HandlleTransferKeys() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description: "Sets `username` as the owner of the key `name`",
		Examples:    []string{"bot username has the name keys"},
		Command:     "{username} has the {name} keys",
		Handler: func(ctx *slacker.CommandContext) {

			username := strings.TrimSpace(strings.ToLower(ctx.Request().Param("username")))
			name := strings.TrimSpace(strings.ToLower(ctx.Request().Param("name")))

			log.GetLogger().Info("Received", slog.String("command", fmt.Sprintf("%s has the %s keys", username, name)))

			err := services.TransferKeys(username, name)

			if err != nil {
				_, err = ctx.Response().ReplyError(err)
				if err != nil {
					log.GetLogger().Error("Error in HandleWhoHasTheKeys", slog.String("error", err.Error()))
				}
				return
			}

			msg := fmt.Sprintf("%s has the %s keys", username, name)

			_, err = ctx.Response().Reply(msg)

			if err != nil {
				log.GetLogger().Error("Error in HandleWhoHasTheKeys", slog.String("error", err.Error()))
			}

			log.GetLogger().Info("Replied", slog.String("message", msg))
		},
	}
}
