package http

import (
	"net/http"
	"strconv"

	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/usecase"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type SkillHandler struct {
    skillUseCase usecase.SkillUsecase
}

func NewSkillHandler(router *gin.Engine, skillUseCase usecase.SkillUsecase) {
		handler := &SkillHandler{skillUseCase}

		skillGroup := router.Group("/skills")
		{
				skillGroup.GET("/", handler.GetAllSkills)
				skillGroup.GET("/:id", handler.GetSkillByID)
				skillGroup.POST("/", handler.CreateSkill)
				skillGroup.PUT("/:id", handler.UpdateSkill)
				skillGroup.DELETE("/:id", handler.DeleteSkill)
		}
}

func (h *SkillHandler) GetAllSkills(c *gin.Context) {
	skills, err := h.skillUseCase.GetAll()
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, skills)
}

func (h *SkillHandler) GetSkillByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
      return
  }

  skill, err := h.skillUseCase.GetByID(uint(id))
  if err != nil {
      if err == gorm.ErrRecordNotFound {
          c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
          return
      } else {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
  }

  c.JSON(http.StatusOK, skill)
}

func (h *SkillHandler) CreateSkill(c *gin.Context) {
	var skill entity.Skill
	if err := c.ShouldBindJSON(&skill); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	err := h.skillUseCase.Create(&skill)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusCreated, skill)
}

func (h *SkillHandler) UpdateSkill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
      return
  }

  var skill entity.Skill
  if err := c.ShouldBindJSON(&skill); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  skill.ID = uint(id)
  err = h.skillUseCase.Update(&skill)
  if err != nil {
      if err == gorm.ErrRecordNotFound {
          c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
          return
      } else {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
  }

  c.JSON(http.StatusOK, skill)
}

func (h *SkillHandler) DeleteSkill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
      return
  }

  err = h.skillUseCase.Delete(uint(id))
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