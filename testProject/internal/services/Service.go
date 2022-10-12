package services

import (
	"context"
	"github.com/gin-gonic/gin"
	operation "testProject/pkg"
)

type (
	OperationService interface {
		PostHandler(computer *operation.Computer)
		ReadHandler(ginCtx gin.Context, computer operation.Computer)
		DeleteHandler(ginCtx gin.Context, computer operation.Computer)
		IsValid() error
	}

	operationServiceImpl struct {
		operationRepository DbRepository
	}

	DbRepository interface {
		PostHandler(ginCtx gin.Context, computer operation.Computer)
		ReadHandler(ginCtx gin.Context, computer operation.Computer)
		DeleteHandler(ginCtx gin.Context, computer operation.Computer)
	}
)

func (d operationServiceImpl) ReadHandler(ctx context.Context, operation *interface{}) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d operationServiceImpl) PostHandler(ctx context.Context, operation *interface{}) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d operationServiceImpl) DeleteHandler(ctx context.Context, operation *interface{}) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewOperationService(operationRepository DbRepository) *operationServiceImpl {
	return &operationServiceImpl{
		operationRepository: operationRepository,
	}

}
