package service

import (
	"fmt"
	"net/http"

	"github.com/kittanutp/salesrecorder/database"
	"github.com/kittanutp/salesrecorder/schema"

	"github.com/gin-gonic/gin"
)

func (db *DBController) GetSaleItems(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (db *DBController) AddSale(c *gin.Context) {
	user, err := GetUserFromCtx(c)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	var addSale schema.AddSaleSchema
	if err := c.BindJSON(&addSale); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sale := database.Sale{Price: addSale.Price, UserID: user.ID}

	var sale_items []database.SaleItem
	for _, each := range addSale.Sales {
		sale_items = append(sale_items, database.SaleItem{
			ItemID: uint(each.ItemID),
			Amount: each.Amount,
			SaleID: sale.ID,
		})
	}

	sale.Items = sale_items
	res := db.Database.Create(&sale)

	if res.Error != nil {
		c.AbortWithStatusJSON(400, fmt.Sprintf("Unable to execute the query. %v", res.Error))
		return
	}

	c.JSON(http.StatusOK, sale)
}
