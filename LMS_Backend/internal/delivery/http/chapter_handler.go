package http

import (
	"LMS_Backend/internal/infrastructure"
	"LMS_Backend/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChapterHandler struct {
	uc usecase.ChapterUsecase
}

func NewChapterHandler(r *gin.Engine, uc usecase.ChapterUsecase) {
	h := &ChapterHandler{uc}

	// list chapter per course (public)
	r.GET("/courses/:courseId/chapters", h.GetByCourse)

	// get chapter by id (public)
	r.GET("/chapters/:id", h.GetByID)

	// create chapter (admin/dosen)
	r.POST("/courses/:courseId/chapters",
		infrastructure.AuthMiddleware(),
		infrastructure.RoleMiddleware("admin", "dosen"),
		h.Create,
	)
	
}

func (h *ChapterHandler) Create(c *gin.Context) {
	courseID, errJSON := strconv.Atoi(c.Param("courseId"))
	
	if errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course_id tidak valid"})
		return
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	errREQ := c.ShouldBindJSON(&req)
	if errREQ != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errREQ.Error()})
		return
	}

	errCREATE := h.uc.Create(uint(courseID), req.Title, req.Content)
	if errCREATE != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCREATE.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "chapter created"})
}

func (h *ChapterHandler) GetByCourse(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "courseId tidak valid"})
		return
	}

	data, errJSON := h.uc.GetByCourse(uint(courseID))
	if errJSON != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errJSON.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *ChapterHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id tidak valid"})
		return
	}

	data, errJSON := h.uc.GetByID(uint(id))	
	if errJSON != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "chapter tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, data)
}