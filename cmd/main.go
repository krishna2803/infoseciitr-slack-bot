package main

import (
	"context"
	"infosec/key-bot/pkg/commands"
	"infosec/key-bot/pkg/database"
	"infosec/key-bot/pkg/log"
	"infosec/key-bot/utils"

	"github.com/slack-io/slacker"
)

func registerCommands(cmdgroup *slacker.CommandGroup) {
	cmdgroup.AddCommand(commands.HandlePing())
}

func initialize() *slacker.Slacker {
	logger := log.NewLogger()
	err := database.Init()
	if err != nil {
		logger.Error(err.Error())
	}

	slackBotToken := utils.GetDotEnvValue("SLACK_BOT_TOKEN")
	slackAppToken := utils.GetDotEnvValue("SLACK_APP_TOKEN")

	bot := slacker.NewClient(
		slackBotToken,
		slackAppToken,
		slacker.WithLogger(log.GetLogger()),
		slacker.WithBotMode(slacker.BotModeIgnoreNone),
	)

	cmdgroup := bot.AddCommandGroup("bot")

	registerCommands(cmdgroup)

	return bot
}

func main() {
	bot := initialize()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.GetLogger().Error(err.Error())
	}
}
