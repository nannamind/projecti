package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nannamind/sa-65/entity"
)

// POST /jobdutiess
func CreateJobDuties(c *gin.Context) {
	var jobduties entity.JobDuties
	if err := c.ShouldBindJSON(&jobduties); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&jobduties).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": jobduties})
}

// GET /jobduties/:id
func GetJobDuties(c *gin.Context) {
	var jobduties entity.JobDuties
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM jobdutiess WHERE id = ?", id).Scan(&jobduties).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobduties})
}

// GET /jobdutiess
func ListJobDutiess(c *gin.Context) {
	var jobdutiess []entity.JobDuties
	if err := entity.DB().Raw("SELECT * FROM jobdutiess").Scan(&jobdutiess).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobdutiess})
}

// DELETE /jobdutiess/:id
func DeleteJobDuties(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM jobdutiess WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jobduties not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /jobdutiess
func UpdateJobDuties(c *gin.Context) {
	var jobduties entity.JobDuties
	if err := c.ShouldBindJSON(&jobduties); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", jobduties.ID).First(&jobduties); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jobduties not found"})
		return
	}

	if err := entity.DB().Save(&jobduties).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobduties})
}
