package services

import (
	"fmt"
	"infoseciitr/slack-bot/pkg/database"
	"infoseciitr/slack-bot/pkg/log"
	"infoseciitr/slack-bot/pkg/models"
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

			mKey := models.Key{
				Owner: "anon",
				Name:  key,
			}

			database.DB.Create(&mKey)
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

	var key models.Key
	database.DB.Where("name = ?", name).First(&key)
	if key.Name == "" {
		log.GetLogger().Warn("Key not found. Assigning to anon", slog.String("Name", name))

		key = models.Key{
			Owner: username,
			Name:  name,
		}

		database.DB.Create(&key)

		return nil
	}

	database.DB.Model(&models.Key{}).Where("name = ?", name).Update("owner", username)

	log.GetLogger().Info("Transferred Keys", slog.String("Owner", username), slog.String("Name", name))

	return nil
}
