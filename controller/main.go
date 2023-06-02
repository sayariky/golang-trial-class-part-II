package controller

import (
	"net/http"
	"time"
	"trial-class/config"
	"trial-class/entity"
	"trial-class/service"

	"github.com/gin-gonic/gin"
)

// disini kita akan mendefinisikan bentuk data yang akan dikirimkan ke client
type OrderRequest struct {
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ProductID    uint   `json:"product_id"`
}

// @Summary Get Product
// @Description Get list of all available Products
// @Tags Product
// @Produce json
// @Success 200 {array} entity.Product
// @Router /products [get]
func HandlerGetProducts(ctx *gin.Context) {
	var products []entity.Product
	err := config.DB.Find(&products).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// @Summary Post Orders
// @Description Create New Order
// @Tags Order
// @Produce json
// @Success 200 {array} entity.Order
// @Router /Orders [post]
func HandlerPostOrder(ctx *gin.Context) {
	// ingin menerima data dari client
	var orderBody OrderRequest
	err := ctx.ShouldBindJSON(&orderBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			//custom error message
			gin.H{"message": "invalid request body"})
		return
	}

	var product entity.Product
	result := config.DB.Where("id = ?", orderBody.ProductID).First(&product)

	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			//custom error message
			gin.H{"message": "product not found"})
		return
	}

	// ini belum pakai validate
	newOrder := entity.Order{
		BuyerEmail:   orderBody.BuyerEmail,
		BuyerAddress: orderBody.BuyerAddress,
		ProductID:    product.ID,
		// ID:           orderBody.ProductID,
		OrderDate: time.Now(),
	}

	err = config.DB.Create(&newOrder).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	// trigger mailer service untuk kirim email
	service.SendMail(newOrder.BuyerEmail, newOrder.BuyerAddress, product.Name)

	ctx.JSON(http.StatusOK, "Suskes membuat order")
}

// @Summary Get Orders
// @Description Get list of all available Orders
// @Tags Order
// @Produce json
// @Success 200 {array} entity.Order
// @Router /orders [get]
func HandlerGetOrders(ctx *gin.Context) {
	var orders []entity.Order

	// err := config.DB.Find(&orders).Error
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	if err := config.DB.Preload("Product").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
