package handler

import "github.com/gin-gonic/gin"

func (h *Handler) responseError(c *gin.Context, statusCode int, message string) {
	h.log.Debug(message)
	c.AbortWithStatusJSON(statusCode, map[string]interface{}{"message": message})
}
