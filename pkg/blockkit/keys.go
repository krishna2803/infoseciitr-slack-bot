package blockkit

import (
	"infoseciitr/slack-bot/pkg/models"

	"github.com/slack-go/slack"
)

func FormatKeys(keys []models.Key) []slack.Block {
	attachments := []slack.Block{}
	attachments = append(attachments,
		slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, "Keys:", false, false),
			nil,
			nil),
	)
	for _, keyOwner := range keys {
		keyOwnerText := "*" + keyOwner.Owner + "* has the " + keyOwner.Name + " keys :key:"
		keyOwnerTextObject := slack.NewTextBlockObject(slack.MarkdownType, keyOwnerText, false, false)
		attachments = append(attachments, slack.NewDividerBlock())
		attachments = append(attachments, slack.NewSectionBlock(
			keyOwnerTextObject,
			nil,
			nil),
		)
	}
	attachments = append(attachments, slack.NewDividerBlock())
	return attachments
}
