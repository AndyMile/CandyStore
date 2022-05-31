package controller

import (
	repository "github.com/AndyMile/candyStore/repository"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	storeRepo repository.StoreRepository
}

func NewBaseHandler(storeRepo *repository.StoreRepo) *BaseHandler {
	return &BaseHandler{
		storeRepo: storeRepo,
	}
}

func (h *BaseHandler) GetAll(c *gin.Context) {
	customers, err := h.storeRepo.GetAll()
	if err != nil {
		fmt.Println("Error", err)
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

func (h *BaseHandler) Get(c *gin.Context) {
	customer, err := h.storeRepo.Get(c.Param("id"))
	if err != nil {
		fmt.Println("Error", err)
	}
	c.JSON(http.StatusOK, gin.H{"article": customer})
}

func (h *BaseHandler) GetTopCustomers(c *gin.Context) {
	customers, err := h.storeRepo.GetTopCustomers()
	if err != nil {
		fmt.Println("Error", err)
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

func (h *BaseHandler) GetCustomersTotalEatenUniqueSnacks(c *gin.Context) {
	customers, err := h.storeRepo.GetCustomersTotalEatenUniqueSnacks()
	if err != nil {
		fmt.Println("Error", err)
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}
