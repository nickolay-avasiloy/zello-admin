package routes

import (
    "fmt"
    "log"
    "net/http"
    "net/url"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/nickolay-avasiloy/zello-admin/api/models"
    "gorm.io/gorm"
)

func SearchRoute(router *gin.Engine, db *gorm.DB) {
    router.GET("/api/search", func(c *gin.Context) {
        rawQuery := c.Query("q")
        query, err := url.QueryUnescape(rawQuery)
        if err != nil {
            log.Printf("Error decoding query parameter: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameter"})
            return
        }

        // Trim spaces from the query
        query = strings.TrimSpace(query)

        var results []models.User = []models.User{}
        if query == "" {
            c.JSON(http.StatusOK, results)
            return
        }

        tsQuery := strings.Join(strings.Fields(query), " | ")
        searchQuery := fmt.Sprintf("%%%s%%", query)

        result := db.Raw(
            `SELECT uuid, username, company, phone_number
             FROM users
             WHERE tsv @@ to_tsquery(?)
                OR username ILIKE ?
                OR company ILIKE ?
                OR phone_number ILIKE ?`,
            tsQuery, searchQuery, searchQuery, searchQuery).Scan(&results)
        if result.Error != nil {
            log.Printf("Error querying database: %v", result.Error)
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Search query error"})
            return
        }

        c.JSON(http.StatusOK, results)
    })
}