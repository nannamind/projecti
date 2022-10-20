package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nannamind/sa-65/entity"
)

// POST /select_genders
func CreateNutritionist(c *gin.Context) {

	var nutritionist entity.Nutritionist
	var jobduties entity.JobDuties
	var admin entity.Admin
	var gender entity.Gender

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Nutritionist
	if err := c.ShouldBindJSON(&nutritionist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", nutritionist.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// 10: ค้นหา jobduties ด้วย id
	if tx := entity.DB().Where("id = ?", nutritionist.JobDutiesID).First(&jobduties); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jobduties not found"})
		return
	}

	// 11: ค้นหา admin ด้วย id
	if tx := entity.DB().Where("id = ?", nutritionist.AdminID).First(&admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admin not found"})
		return
	}
	// 12: สร้าง Nutritionist
	wv := entity.Nutritionist{
		JobDuties:  jobduties,             // โยงความสัมพันธ์กับ Entity JobDuties
		Gender:       gender,                  // โยงความสัมพันธ์กับ Entity Gender
		Admin:    admin,               // โยงความสัมพันธ์กับ Entity Admin
		DOB: nutritionist.DOB, // ตั้งค่าฟิลด์ DOB
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": wv})
}

// GET /nutritionist/:id
func GetNutritionist(c *gin.Context) {
	var nutritionist entity.Nutritionist
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&nutritionist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nutritionist not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nutritionist})
}

// GET /nutritionists
func ListNutritionists(c *gin.Context) {
	var nutritionists []entity.Nutritionist
	if err := entity.DB().Preload("JobDuties").Preload("Admin").Preload("Gender").Raw("SELECT * FROM nutritionists").Find(&nutritionists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nutritionists})
}

// DELETE /nutritionists/:id
func DeleteNutritionist(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM nutritionists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nutritionist not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /nutritionists
func UpdateNutritionist(c *gin.Context) {
	var nutritionist entity.Nutritionist
	if err := c.ShouldBindJSON(&nutritionist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", nutritionist.ID).First(&nutritionist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nutritionist not found"})
		return
	}

	if err := entity.DB().Save(&nutritionist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nutritionist})
}