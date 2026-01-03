package http

import (
	"net/http"

	"LMS_Backend/internal/infrastructure"
	"LMS_Backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC usecase.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, authUC usecase.AuthUsecase) {
	handler := &AuthHandler{authUC}
	group := r.Group("/auth")

	group.POST("/register", handler.Register)
	group.POST("/login", handler.Login)

	group.GET("/me", infrastructure.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user_id": c.GetUint("user_id"),
			"role":    c.GetString("role"),
		})
	})

}


func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.authUC.Register(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "register berhasil"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := h.authUC.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login berhasil",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
