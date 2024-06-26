package main

import (
    "fmt"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"

	"github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

type User struct {
    UUID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uuid"`
    Username    string `gorm:"uniqueIndex" json:"username"`
    Company     string `json:"company"`
    PhoneNumber string `json:"phoneNumber"`
}

func main() {
    connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
    db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    })
    if err != nil {
        log.Fatal(err)
    }

    // Run migrations
    runMigrations()

    router := gin.Default()

	// Set up CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

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

        var results []User = []User{}
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



	router.Run(":8080")
}
