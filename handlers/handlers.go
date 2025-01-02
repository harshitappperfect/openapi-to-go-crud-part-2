package handlers

import (
	"net/http"
	"sample/db"

	"sample/models"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, name, description FROM items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Id, &item.Name, &item.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.DB.QueryRow("INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id", item.Name, item.Description).Scan(&item.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}
