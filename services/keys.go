package services

import (
	"fmt"
	"infoseciitr/slack-bot/models"
	"infoseciitr/slack-bot/pkg/database"
	"infoseciitr/slack-bot/pkg/log"
	"log/slog"
	"slices"
)

var validKeys []string = []string{"master"}

func WhoHasTheKeys() []models.Key {
	log.GetLogger().Info("Who has the keys")

	var keys []models.Key

	database.DB.Find(&keys)

	for _, key := range validKeys {
		found := false
		for _, k := range keys {
			if k.Name == key {
				found = true
				break
			}
		}
		if !found {
			log.GetLogger().Warn("Key not found. Assigning to anon", slog.String("Name", key))

			keys = append(keys, models.Key{
				Owner: "anon",
				Name:  key,
			})

			database.DB.Save(&key)
		}
	}

	for _, key := range keys {
		log.GetLogger().Info("Key:", slog.String("Owner", key.Owner), slog.String("Name", key.Name))
	}
	return keys
}

func TransferKeys(username string, name string) error {

	if !slices.Contains(validKeys, name) {
		log.GetLogger().Error("Invalid key", slog.String("Name", name))
		return fmt.Errorf("Invalid key: %s", name)
	}

	log.GetLogger().Info("Transferred Keys", slog.String("Owner", username), slog.String("Name", name))

	key := models.Key{
		Owner: username,
		Name:  name,
	}

	database.DB.Save(&key)

	return nil
}
