package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kittanutp/salesrecorder/database"
	"github.com/kittanutp/salesrecorder/schema"
)

func (db *DBController) GetItems(c *gin.Context) {
	var items []database.Item
	res := db.Database.Find(&items)
	if res.Error != nil {
		log.Fatalf("Unable to execute the query. %v", res.Error)
	}
	c.JSON(http.StatusOK, items)
}

func (db *DBController) GetItemsByUser(c *gin.Context) {
	user, err := GetUserFromCtx(c)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	var items []database.Item

	res := db.Database.Find(&items, "user_id = ?", user.ID)

	if res.Error != nil {
		log.Fatalf("Unable to execute the query. %v", res.Error)
	}
	c.JSON(http.StatusOK, items)
}

func (db *DBController) CreateItem(c *gin.Context) {
	user, err := GetUserFromCtx(c)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	var requestedItem schema.Item

	if err := c.BindJSON(&requestedItem); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	item := database.Item{
		Name:   requestedItem.Name,
		Cost:   requestedItem.Cost,
		UserID: user.ID,
	}
	res := db.Database.Create(&item)

	if res.Error != nil {
		log.Fatalf("Unable to execute the query. %v", res.Error)
	}

	log.Printf("Inserted a single record %v", item.ID)
	c.JSON(http.StatusCreated, item)
}
