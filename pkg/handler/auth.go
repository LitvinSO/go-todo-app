package handler

import (
	"net/http"

	"github.com/LitvinSO/go-todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
