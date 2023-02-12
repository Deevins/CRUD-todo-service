package handler

import (
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, "user id not found")
	}

	var input entity.TodoList
	if err := ctx.Bind(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	// call service method next
	id, err := h.services.TodoList.CreateList(userID, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []entity.TodoList `json:"data"`
}

func (h *Handler) getAllLists(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid id param")
	}

	list, err := h.services.TodoList.GetByID(userID, id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)

}

func (h *Handler) deleteList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid id param")
	}

	err = h.services.TodoList.Delete(userID, id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) updateList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid id param")
		return
	}

	var input entity.UpdateListInput

	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userID, id, input); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
