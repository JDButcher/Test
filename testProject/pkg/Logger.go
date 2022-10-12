package pkg

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type contextKey string

const BackendOperation contextKey = "backend operation"

func Logger(ctx context.Context) *log.Entry {
	logger := log.WithField(string(BackendOperation), string(BackendOperation))
	if v := ctx.Value(BackendOperation); v != nil {
		logger = logger.WithField(string(BackendOperation), v)
	}
	return logger
}
