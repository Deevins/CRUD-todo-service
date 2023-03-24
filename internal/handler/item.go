package handler

import (
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) createItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	listID := uuid.Must(uuid.FromString(ctx.Param("id")))

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

	listID := uuid.Must(uuid.FromString(ctx.Param("id")))

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
	itemID := uuid.Must(uuid.FromString(ctx.Param("id")))

	item, err := h.services.TodoItem.GetItemByID(userID, itemID)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(ctx *gin.Context) {}

func (h *Handler) deleteItem(ctx *gin.Context) {}
