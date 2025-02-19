package http

import (
	"net/http"

	"strconv"

	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/usecase"

	"github.com/gin-gonic/gin"
)

type FormalEducationHandler struct {
    eduUseCase usecase.FormalEducationUsecase
}

func NewFormalEducationHandler(router *gin.Engine, eduUseCase usecase.FormalEducationUsecase) {
    handler := &FormalEducationHandler{eduUseCase}

    eduGroup := router.Group("/formal-educations")
    {
        eduGroup.GET("/", handler.GetAllEducations)
        eduGroup.GET("/:id", handler.GetEducationByID)
        eduGroup.POST("/", handler.CreateEducation)
        eduGroup.PUT("/:id", handler.UpdateEducation)
        eduGroup.DELETE("/:id", handler.DeleteEducation)
    }
}

func (h *FormalEducationHandler) GetAllEducations(c *gin.Context) {
    educations, err := h.eduUseCase.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, educations)
}

func (h *FormalEducationHandler) GetEducationByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    education, err := h.eduUseCase.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, education)
}

func (h *FormalEducationHandler) CreateEducation(c *gin.Context) {
    var education entity.FormalEducation
    if err := c.ShouldBindJSON(&education); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.eduUseCase.Create(&education); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, education)
}

func (h *FormalEducationHandler) UpdateEducation(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var education entity.FormalEducation
    if err := c.ShouldBindJSON(&education); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    education.ID = uint(id)

    if err := h.eduUseCase.Update(&education); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, education)
}

func (h *FormalEducationHandler) DeleteEducation(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.eduUseCase.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Education deleted successfully"})
}
