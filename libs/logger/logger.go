package logger

import (
	"os"

	"github.com/apsdehal/go-logger"
)

type appLogger struct {
	*logger.Logger
}

func NewAppLogger() (*appLogger, error) {
	log, err := logger.New("DreonECommerce", 1, os.Stdout)
	if err != nil {
		return nil, err
	}
	log.SetFormat("{module: %{module}, level: %{level}, message: %{message}}")

	return &appLogger{
		Logger: log,
	}, nil
}
