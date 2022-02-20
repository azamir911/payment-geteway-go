package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Serve() {

	engine := gin.Default()
	engine.GET("/index", index)
	engine.GET("/payments", getAllInvoice)
	engine.GET("/payment/:id", getInvoice)
	engine.POST("/payment", postInvoice)
	engine.POST("/close", closeChannels)

	engine.Run("localhost:8080")

}

func index(context *gin.Context) {

}

func getAllInvoice(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "getAll")
}

func getInvoice(context *gin.Context) {
	id := context.Param("id")
	msg := fmt.Sprintf("getInvoice for %v", id)
	context.IndentedJSON(http.StatusOK, msg)
}

func postInvoice(context *gin.Context) {
	context.IndentedJSON(http.StatusCreated, "Create new invoice")
}

func closeChannels(context *gin.Context) {
	context.IndentedJSON(http.StatusCreated, "Close channels")
}
