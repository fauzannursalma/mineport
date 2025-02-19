package http

import (
	"net/http"
	"strconv"

	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/usecase"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
    projectUseCase usecase.ProjectUsecase
}

func NewProjectHandler(router *gin.Engine, projectUseCase usecase.ProjectUsecase) {
		handler := &ProjectHandler{projectUseCase}

		projectGroup := router.Group("/projects")
		{
				projectGroup.GET("/", handler.GetAllProjects)
				projectGroup.GET("/:id", handler.GetProjectByID)
				projectGroup.POST("/", handler.CreateProject)
				projectGroup.PUT("/:id", handler.UpdateProject)
				projectGroup.DELETE("/:id", handler.DeleteProject)
		}
}

func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	projects, err := h.projectUseCase.GetAll()
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
      return
  }

  project, err := h.projectUseCase.GetByID(uint(id))
  if err != nil {
      if err == gorm.ErrRecordNotFound {
          c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
          return
      } else {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
  }

  c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var project entity.Project
	if err := c.ShouldBindJSON(&project); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	err := h.projectUseCase.Create(&project)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
      return
  }

  var project entity.Project
  if err := c.ShouldBindJSON(&project); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  project.ID = uint(id)
  err = h.projectUseCase.Update(&project)
  if err != nil {
      if err == gorm.ErrRecordNotFound {
          c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
          return
      } else {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
  }

  c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
      return
  }

  err = h.projectUseCase.Delete(uint(id))
  if err != nil {
      if err == gorm.ErrRecordNotFound {
          c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
          return
      } else {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
  }

  c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}