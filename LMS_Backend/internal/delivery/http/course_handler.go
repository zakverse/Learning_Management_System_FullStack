package http

import (
	"LMS_Backend/internal/infrastructure"
	"LMS_Backend/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



type CourseHandler struct {
	uc usecase.CourseUsecase
}

func NewCourseHandler(r *gin.Engine, uc usecase.CourseUsecase) {
	h := &CourseHandler{uc}

	group := r.Group("/courses")

	// read (semua role)
	group.GET("", h.GetAll)
	group.GET("/:courseId", h.GetByID)

	group.PUT("/:id",
	infrastructure.AuthMiddleware(),
	infrastructure.RoleMiddleware("admin", "dosen"),
	h.Update,
)

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

func (h *CourseHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "id tidak valid"})
		return
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "body tidak valid"})
		return
	}

	err = h.uc.Update(uint(id), req.Title, req.Description)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "course updated"})
}
