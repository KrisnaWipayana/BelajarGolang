package adminController

import (
	"encoding/json"
	"net/http"

	"github.com/KrisnaWipayana/BelajarGO/GolangAPI/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(c *gin.Context) {

	var user []model.User

	model.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Show(c *gin.Context) {

	var user model.User
	id := c.Param("id")

	if err := model.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Add(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} //mengambil body JSON untuk dimasukkan ke dalam struct product

	model.DB.Create(&user)                     //memasukkan ke dalam database
	c.JSON(http.StatusOK, gin.H{"user": user}) //menampilkan respon JSON
}

func Update(c *gin.Context) {

	var user model.User //menampung value yang didapat dari JSON

	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate user"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil di-update"})
}

func Delete(c *gin.Context) {

	var user model.User

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if model.DB.Delete(&user, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}
