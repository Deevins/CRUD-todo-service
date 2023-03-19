package handler

import (
	"fmt"
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid id param")
		return
	}

	var input entity.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoItem.Create(userID, listID, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})

}

func (h *Handler) getAllItems(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid id param")
		fmt.Println("error in getAllItems method")
		return
	}

	items, err := h.services.TodoItem.GetAll(userID, listID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, items)

}

func (h *Handler) getItemById(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		logrus.Error("no user ID provided")
		return

	}

	itemID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "incorrect item ID")
	}

	item, err := h.services.TodoItem.GetItemByID(userID, itemID)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(ctx *gin.Context) {}

func (h *Handler) deleteItem(ctx *gin.Context) {}
