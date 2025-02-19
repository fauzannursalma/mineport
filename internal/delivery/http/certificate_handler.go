package http

import (
	"net/http"

	"strconv"

	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CertificateHandler struct {
    certificateUseCase usecase.CertificateUsecase
}

func NewCertificateHandler(router *gin.Engine, certificateUseCase usecase.CertificateUsecase) {
    handler := &CertificateHandler{certificateUseCase}

    certificateGroup := router.Group("/formal-certificatecations")
    {
        certificateGroup.GET("/", handler.GetAllEducations)
        certificateGroup.GET("/:id", handler.GetEducationByID)
        certificateGroup.POST("/", handler.CreateEducation)
        certificateGroup.PUT("/:id", handler.UpdateEducation)
        certificateGroup.DELETE("/:id", handler.DeleteEducation)
    }
}

func (h *CertificateHandler) GetAllEducations(c *gin.Context) {
    certificatecations, err := h.certificateUseCase.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, certificatecations)
}

func (h *CertificateHandler) GetEducationByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    certificatecation, err := h.certificateUseCase.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, certificatecation)
}

func (h *CertificateHandler) CreateEducation(c *gin.Context) {
    var certificatecation entity.Certificate
    if err := c.ShouldBindJSON(&certificatecation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.certificateUseCase.Create(&certificatecation); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, certificatecation)
}

func (h *CertificateHandler) UpdateEducation(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var certificatecation entity.Certificate
    if err := c.ShouldBindJSON(&certificatecation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    certificatecation.ID = uint(id)

    if err := h.certificateUseCase.Update(&certificatecation); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, certificatecation)
}

func (h *CertificateHandler) DeleteEducation(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.certificateUseCase.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Education deleted successfully"})
}
