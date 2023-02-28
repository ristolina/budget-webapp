package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type DBentry struct {
	id             int
	household_expenses       int
	food_expenses       int
	transport_expenses  int
	misc_expenses   int
	created_at string
}

func postExpense(c *gin.Context) {
	var userData DBentry

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Query for a value based on a single row.
	result, err := db.Exec("INSERT INTO expenses (household_expenses, food_expenses, transport_expenses, misc_expenses) VALUES (?, ?, ?, ?)", userData.household_expenses, userData.food_expenses, userData.transport_expenses, userData.misc_expenses)
	if err != nil {
		fmt.Println(err)
		fmt.Println("fel!")
	}
	fmt.Println(result)
	return

}

func getExpense(c *gin.Context) {
	var res DBentry
	// Query for a value based on a single row.
	row := db.QueryRow("SELECT id, household_expenses, food_expenses, transport_expenses, misc_expenses, created_at FROM expenses ORDER BY id DESC LIMIT 1")
	if err := row.Scan(&res.id, &res.household_expenses, &res.food_expenses, &res.transport_expenses, &res.misc_expenses, &res.created_at); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "No data",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":  "success",
			"message": res,
		})
	}
}

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to DB!")
	router := gin.Default()
	router.GET("expense", getExpense)
	router.POST("/expense", postExpense)

	router.Run("0.0.0.0:5001")
}
