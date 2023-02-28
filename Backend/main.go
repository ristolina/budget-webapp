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
	Id             int `json:"id"`
	Timestamp string `json:"created_at"`
}

type ExpenseEntry struct {
	Household      int `json:"household" form:"household"`
	Food       int `json:"food" form:"food"`
	Transport  int `json:"transport" form:"transport"`
	Misc   int `json:"misc" form:"misc"`
}

func postExpense(c *gin.Context) {
	var userData ExpenseEntry
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(c.PostForm("household"))
	// Query for a value based on a single row.
	result, err := db.Exec("INSERT INTO expenses (household_expenses, food_expenses, transport_expenses, misc_expenses) VALUES (?, ?, ?, ?)", userData.Household, userData.Food, userData.Transport, userData.Misc)
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(userData)
	return

}

func getExpense(c *gin.Context) {
	var res DBentry
	var expenses ExpenseEntry
	// Query for a value based on a single row.
	row := db.QueryRow("SELECT id, household_expenses, food_expenses, transport_expenses, misc_expenses, created_at FROM expenses ORDER BY id DESC LIMIT 1")
	if err := row.Scan(&res.Id, &expenses.Household, &expenses.Food, &expenses.Transport, &expenses.Misc, &res.Timestamp); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "No data",
		})
	} else {
		//total, err := strconv.Atoi(expenses.Household) + strconv.Atoi(expenses.Food) + strconv.Atoi(expenses.Transport) + strconv.Atoi(expenses.Misc)
		total := expenses.Household + expenses.Food + expenses.Transport + expenses.Misc
		c.JSON(http.StatusOK, gin.H{
			"result":  "success",
			"expenses": expenses,
			"id": res.Id,
			"timestamp": res.Timestamp,
			"total": total,
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
	router.GET("/expense", getExpense)
	router.POST("/expense", postExpense)

	router.Run("0.0.0.0:5001")
}
