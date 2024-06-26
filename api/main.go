package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
    runMigrations()

	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	// Set up CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

    router.GET("/api/ping", func(c *gin.Context) {
        var response string
        err := db.QueryRow("SELECT 'pong'").Scan(&response)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Database query error"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": response})
    })

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

        var results = []map[string]interface{}{}
        if query == "" {
            c.JSON(http.StatusOK, results)
            return
        }

        tsQuery := strings.Join(strings.Fields(query), " | ")
        searchQuery := fmt.Sprintf("%%%s%%", query)
        rows, err := db.Query(
            `SELECT username, company, phone_number
             FROM users
             WHERE tsv @@ to_tsquery($1)
                OR username ILIKE $2
                OR company ILIKE $2
                OR phone_number ILIKE $2`,
            tsQuery, searchQuery)
        if err != nil {
            log.Printf("Error querying database: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Search query error"})
            return
        }
        defer rows.Close()

        for rows.Next() {
            var username, company, phoneNumber string
            if err := rows.Scan(&username, &company, &phoneNumber); err != nil {
                log.Printf("Error scanning row: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"message": "Error scanning results"})
                return
            }
            result := map[string]interface{}{
                "username":     username,
                "company":      company,
                "phoneNumber":  phoneNumber,
            }
            results = append(results, result)
        }

        if err = rows.Err(); err != nil {
            log.Printf("Error iterating over rows: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Error iterating over rows"})
            return
        }

        c.JSON(http.StatusOK, results)
    })



	router.Run(":8080")
}
