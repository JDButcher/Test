package rest

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	service "testProject/internal/services"
	logger "testProject/pkg"
)

type operationHandler struct {
	operationService service.OperationService
}

func NewOperationHandler(operationService service.OperationService) *operationHandler {
	return &operationHandler{
		operationService: operationService,
	}
}

func (oh *operationHandler) PostHandler(ginCtx gin.Context, db *sql.DB) {
	ctx := ginCtx.Request.Context()
	log := logger.Logger(ctx)

	computer := &logger.Computer{}
	err := ginCtx.ShouldBindJSON(computer)
	if err != nil {
		log.WithError(err).Error("could not deserialize request args")
		ginCtx.Status(http.StatusBadRequest)
		return
	}

	err = computer.IsValid()
	if err != nil {
		log.WithError(err)
		ginCtx.Status(http.StatusBadRequest)
		return
	}

	oh.operationService.PostHandler(computer)

	// ToDo
	// if PC counter for User > 3 send message to admin

	ctx = context.WithValue(ctx, "ok", "ok")
	log = logger.Logger(ctx)
}

func (oh *operationHandler) ReadHandler(ginCtx gin.Context, db *sql.DB) {
	// ToDo

}

func (oh *operationHandler) DeleteHandler(ginCtx gin.Context, db *sql.DB) {
	// ToDo
	return
}
