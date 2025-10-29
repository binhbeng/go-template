package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/binhbeng/goex/internal/pkg/errors"
	// log "github.com/binhbeng/goex/internal/pkg/logger"
	r "github.com/binhbeng/goex/internal/pkg/response"
	// "go.uber.org/zap"
)

type Api struct {
	errors.Error
}

func (api Api) Success(c *gin.Context, data ...any) {
	response := r.Resp()
	if data != nil {
		response.WithDataSuccess(c, data[0])
		return
	}
	response.Success(c)
}

func (api Api) FailCode(c *gin.Context, code int, data ...any) {
	response := r.Resp()
	if data != nil {
		response.WithData(data[0]).FailCode(c, code)
		return
	}
	response.FailCode(c, code)
}

func (api Api) Fail(c *gin.Context, code int, message string, data ...any) {
	response := r.Resp()
	if data != nil {
		response.WithData(data[0]).FailCode(c, code, message)
		return
	}
	response.FailCode(c, code, message)
}

func (api Api) Err(c *gin.Context, e error) {
	businessError, err := api.AsBusinessError(e)
	if err != nil {
		// log.Logger.Warn("Unknown error:", zap.Any("Error reason:", err))
		api.FailCode(c, errors.ServerError)
		return
	}

	api.Fail(c, businessError.GetCode(), businessError.GetMessage())
}
