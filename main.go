package main

import (
	"errors"
	"net/http"
	"strconv"

	"customer-service/database"
	"customer-service/models"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func main() {
	database.Connect()

	// Tạo bảng và cấu hình auto increment
	database.DB.AutoMigrate(&models.Customer{})

	// Thiết lập sequence cho PostgreSQL
	var count int64
	database.DB.Model(&models.Customer{}).Count(&count)
	if count == 0 {
		database.DB.Exec("ALTER SEQUENCE customers_id_seq RESTART WITH 10000000;")
	}

	router := gin.Default()

	// Endpoint hiển thị tên Minh Kong
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Minh dep trai")
	})

	router.GET("/customers", GetCustomers)
	router.GET("/customers/:id", GetCustomer)
	router.POST("/customers", CreateCustomer)
	router.PUT("/customers/:id", UpdateCustomer)
	router.DELETE("/customers/:id", DeleteCustomer)

	router.Run(":8080")
}

// Cập nhật các hàm xử lý
func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	if result := database.DB.Find(&customers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func GetCustomer(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID phải là số nguyên dương"})
		return
	}

	var customer models.Customer
	if result := database.DB.First(&customer, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy khách hàng"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func CreateCustomer(c *gin.Context) {
	var input models.CustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customer{
		Name:  input.Name,
		Phone: input.Phone,
		Email: input.Email,
	}

	if result := database.DB.Create(&customer); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func UpdateCustomer(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID phải là số nguyên dương"})
		return
	}

	var customer models.Customer
	if result := database.DB.First(&customer, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy khách hàng"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Save(&customer); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID phải là số nguyên dương"})
		return
	}

	if result := database.DB.Delete(&models.Customer{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa khách hàng thành công"})
}
