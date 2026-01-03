package http

import (
	"LMS_Backend/internal/infrastructure"
	"LMS_Backend/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func NewCourseHandler(r *gin.Engine, uc usecase.CourseUsecase) {
// 	auth := r.Group("/courses")
// 	auth.Use(infrastructure.AuthMiddleware())
// 	auth.Use(infrastructure.RoleMiddleware("admin", "dosen"))

// 	auth.POST("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"message": "course created"})
// 	})
// }

type CourseHandler struct {
	uc usecase.CourseUsecase
}

func NewCourseHandler(r *gin.Engine, uc usecase.CourseUsecase) {
	h := &CourseHandler{uc}

	group := r.Group("/courses")

	// read (semua role)
	group.GET("", h.GetAll)
	group.GET("/:courseId", h.GetByID)

	// write (admin & dosen)
	group.POST("",
		infrastructure.AuthMiddleware(),
		infrastructure.RoleMiddleware("admin", "dosen"),
		h.Create,
	)
}

func (h *CourseHandler) Create(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	if err := h.uc.Create(req.Title, req.Description, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "course created"})
}

func (h *CourseHandler) GetAll(c *gin.Context) {
	data, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *CourseHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "courseId tidak valid"})
		return
	}

	data, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "course tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, data)
}
