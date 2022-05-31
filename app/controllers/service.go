package controller

import "github.com/gin-gonic/gin"

type StoreController interface {
	GetAll(c *gin.Context) 
	Get(c *gin.Context) 
	GetCustomersTotalEatenUniqueSnacks(c *gin.Context)
	GetTopCustomers(c *gin.Context)
}