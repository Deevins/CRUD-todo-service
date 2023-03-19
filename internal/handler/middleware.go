package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	Authorization = "Authorization"
	userCtx       = "userID"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(Authorization)
	if header == "" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerSplit := strings.Split(header, " ")
	if len(headerSplit) != 2 || headerSplit[0] != "Bearer" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerSplit[1]) == 0 {
		NewErrorResponse(ctx, http.StatusUnauthorized, "token is empty")
		return
	}

	//parse Token
	userID, err := h.services.Authorization.ParseToken(headerSplit[1])
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userID)
}

func getUserID(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idToInt, ok := id.(int) // trying to parse id from "any" to "int" type
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id is not of valid type")
		return 0, errors.New("user id is not of valid type")
	}
	return idToInt, nil
}
