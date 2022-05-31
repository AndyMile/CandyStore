package rest

import (
	controller "github.com/AndyMile/candyStore/controllers"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func RouterHandler(c *controller.BaseHandler) {
	router := gin.Default()

	router.GET("/customer/:id", c.Get)
	router.GET("/customers", c.GetAll)
	router.GET("/customers/stats", c.GetCustomersTotalEatenUniqueSnacks)
	router.GET("/customers/top", c.GetTopCustomers)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}