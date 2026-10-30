package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func BadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
		Message: message,
	})
}

func Unauthorized(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
		Message: message,
	})
}

func Forbidden(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusForbidden, ErrorResponse{
		Message: message,
	})
}

func NotFound(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse{
		Message: message,
	})
}

func MethodNotAllowed(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, ErrorResponse{
		Message: message,
	})
}

func Conflict(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusConflict, ErrorResponse{
		Message: message,
	})
}

func Gone(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusGone, ErrorResponse{
		Message: message,
	})
}

func UnprocessableEntity(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrorResponse{
		Message: message,
	})
}

func TooManyRequests(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusTooManyRequests, ErrorResponse{
		Message: message,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
		Message: message,
	})
}

func NotImplemented(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, ErrorResponse{
		Message: message,
	})
}

func BadGateway(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadGateway, ErrorResponse{
		Message: message,
	})
}

func ServiceUnavailable(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusServiceUnavailable, ErrorResponse{
		Message: message,
	})
}

func GatewayTimeout(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusGatewayTimeout, ErrorResponse{
		Message: message,
	})
}

func PayloadTooLarge(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, ErrorResponse{
		Message: message,
	})
}

func UnsupportedMediaType(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, ErrorResponse{
		Message: message,
	})
}

func RequestTimeout(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusRequestTimeout, ErrorResponse{
		Message: message,
	})
}

func NotAcceptable(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusNotAcceptable, ErrorResponse{
		Message: message,
	})
}
