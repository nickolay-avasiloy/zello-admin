package routes

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/nickolay-avasiloy/zello-admin/api/models"
    "gorm.io/gorm"
)

func UsersRoute(router *gin.Engine, db *gorm.DB) {
    router.GET("/api/users", func(c *gin.Context) {
        var users []models.User
        result := db.Raw("SELECT * FROM users ORDER BY RANDOM() LIMIT 10").Scan(&users)
        if result.Error != nil {
            log.Printf("Error querying database: %v", result.Error)
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving users"})
            return
        }

        c.JSON(http.StatusOK, users)
    })

    router.DELETE("/api/users/:uuid", func(c *gin.Context) {
        userUUID := c.Param("uuid")
        if _, err := uuid.Parse(userUUID); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID"})
            return
        }

        result := db.Delete(&models.User{}, "uuid = ?", userUUID)
        if result.Error != nil {
            log.Printf("Error deleting user: %v", result.Error)
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user"})
            return
        }

        if result.RowsAffected == 0 {
            c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
    })
}