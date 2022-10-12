package pkg

import (
	"context"
)

type (
	DbRepository interface {
		ReadHandler(ctx context.Context, operation *ReadHandler) (string, error)
	}

	OperationService interface {
		ReadHandler(ctx context.Context, operation *ReadHandler) (string, error)
	}

	operationServiceImpl struct {
		dbRepsitory DbRepository
	}

	ReadHandler struct {
		Id string `json:"id"`
	}
)

func NewOperatonService(dbRepository DbRepository) *operationServiceImpl {
	return &operationServiceImpl{
		dbRepsitory: dbRepository,
	}
}

func (os *operationServiceImpl) ReadHandler(ctx context.Context, operation *ReadHandler) (string, error) {
	return "nil", nil
}
