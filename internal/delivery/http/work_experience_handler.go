package http

import (
	"net/http"
	"strconv"

	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/usecase"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type WorkExperience struct {
    workExperienceUseCase usecase.SkillUsecase
}

func NewWorkExperience(router *gin.Engine, workExperienceUseCase usecase.SkillUsecase) {
		handler := &WorkExperience{workExperienceUseCase}

		workExperienceGroup := router.Group("/workExperiences")
		{
				workExperienceGroup.GET("/", handler.GetAllSkills)
				workExperienceGroup.GET("/:id", handler.GetSkillByID)
				workExperienceGroup.POST("/", handler.CreateSkill)
				workExperienceGroup.PUT("/:id", handler.UpdateSkill)
				workExperienceGroup.DELETE("/:id", handler.DeleteSkill)
		}
}

func (h *WorkExperience) GetAllSkills(c *gin.Context) {
	workExperiences, err := h.workExperienceUseCase.GetAll()
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, workExperiences)
}

func (h *WorkExperience) GetSkillByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workExperience ID"})
      return
  }

  workExperience, err := h.workExperienceUseCase.GetByID(uint(id))
  if err != nil {
      if err == gorm.ErrRecordNotFound {
          c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
          return
      } else {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
  }

  c.JSON(http.StatusOK, workExperience)
}

func (h *WorkExperience) CreateSkill(c *gin.Context) {
	var workExperience entity.Skill
	if err := c.ShouldBindJSON(&workExperience); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	err := h.workExperienceUseCase.Create(&workExperience)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusCreated, workExperience)
}

func (h *WorkExperience) UpdateSkill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workExperience ID"})
      return
  }

  var workExperience entity.Skill
  if err := c.ShouldBindJSON(&workExperience); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  workExperience.ID = uint(id)
  err = h.workExperienceUseCase.Update(&workExperience)
  if err != nil {
      if err == gorm.ErrRecordNotFound {
          c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
          return
      } else {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
  }

  c.JSON(http.StatusOK, workExperience)
}

func (h *WorkExperience) DeleteSkill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workExperience ID"})
      return
  }

  err = h.workExperienceUseCase.Delete(uint(id))
  if err != nil {
      if err == gorm.ErrRecordNotFound {
          c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
          return
      } else {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
  }

  c.JSON(http.StatusOK, gin.H{"message": "Skill deleted"})
}