package handler

import (
	"net/http"
	"strconv"

	"github.com/a1exander256/todo/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	var input models.TodoItem
	if err := c.BindJSON(&input); err != nil {
		h.responseError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.TodoItem.Create(input)
	if err != nil {
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.log.Debugf("item %d added", id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.responseError(c, http.StatusBadRequest, err.Error())
		return
	}
	item, err := h.service.TodoItem.GetById(itemId)
	if err != nil {
		h.responseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}
