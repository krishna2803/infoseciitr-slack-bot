package main

import (
	"context"
	"infoseciitr/slack-bot/pkg/commands"
	"infoseciitr/slack-bot/pkg/database"
	"infoseciitr/slack-bot/pkg/log"
	"infoseciitr/slack-bot/pkg/utils"

	"github.com/slack-io/slacker"
)

func registerCommands(cmdgroup *slacker.CommandGroup) {
	cmdgroup.AddCommand(commands.HandlePing())
	cmdgroup.AddCommand(commands.HandleWhoHasTheKeys())
	cmdgroup.AddCommand(commands.HandlleTransferKeys())
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
		slacker.WithBotMode(slacker.BotModeIgnoreAll),
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
