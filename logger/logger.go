package logger

import (
	"github.com/rodiond26/gomigrator/config"
	"go.uber.org/zap"
)

func GetLogger(cfg *config.Config) (l *zap.Logger, err error) {
	l, err = zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return l, nil
}
