package log

import (
	"errors"
	"infoseciitr/slack-bot/pkg/utils"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
)

var (
	logDir  = "logs"
	logFile = "./" + logDir + "/slack-bot.log"
	logger  *slog.Logger
)

// initializes the logger
func initLogger(isProd bool) *slog.Logger {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(logDir, os.ModePerm)
		if mkdirErr != nil {
			panic(errors.New("unable to create log directory: " + err.Error()))
		}
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(errors.New("unable to open log file: " + err.Error()))
	}

	level := slog.LevelDebug
	if isProd {
		level = slog.LevelInfo
	}

	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: level,
	})

	consoleHandler := tint.NewHandler(colorable.NewColorable(os.Stdout), &tint.Options{
		Level:      level,
		TimeFormat: "2006-01-02 15:04:05",
	})

	multiHandler := slogmulti.Fanout(fileHandler, consoleHandler)

	return slog.New(multiHandler)
}

// creates a new logger instance
func NewLogger() *slog.Logger {
	isProd := utils.IsProd()
	logger = initLogger(isProd)
	return logger
}

// returns the global logger instance
func GetLogger() *slog.Logger {
	return logger
}
