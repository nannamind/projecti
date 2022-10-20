package controller

import (
	"github.com/nannamind/sa-65/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)


// GET /admins
// List all admins
func ListAdmins(c *gin.Context) {
	var admins []entity.Admin
	if err := entity.DB().Raw("SELECT * FROM admins").Scan(&admins).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admins})
}

// GET /admin/:id
// Get admin by id
func GetAdmin(c *gin.Context) {
	var admin entity.Admin
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admin})
}

// PATCH /admins
func UpdateAdmin(c *gin.Context) {
	var admin entity.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", admin.ID).First(&admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admin not found"})
		return
	}

	if err := entity.DB().Save(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admin})
}

// DELETE /admins/:id
func DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM admins WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admin not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.Admin{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}