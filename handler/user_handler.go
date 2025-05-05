
package handler

import (
    "net/http"
    "persona-core/model"
    "persona-core/service"
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    Service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
    return &UserHandler{Service: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdUser, err := h.Service.CreateUser(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, createdUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := h.Service.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}
