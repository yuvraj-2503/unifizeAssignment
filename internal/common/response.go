package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(ctx *gin.Context, result interface{}) {
	ctx.JSON(http.StatusOK, result)
}

func ConflictError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusConflict, Result{
		Code:    http.StatusConflict,
		Message: message,
	})
}

func NotFound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, Result{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

func BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, Result{
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

func Unauthorized(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, Result{
		Code:    http.StatusUnauthorized,
		Message: message,
	})
}

func Forbidden(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusForbidden, Result{
		Code:    http.StatusForbidden,
		Message: message,
	})
}

func InternalError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, Result{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}
