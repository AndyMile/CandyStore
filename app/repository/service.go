package repository

import model "github.com/AndyMile/candyStore/models"


type StoreRepository interface {
	GetAll() ([]model.Customer, error)
	Get(id string) (model.Customer, error)
	GetTopCustomers() ([]model.TopCustomer, error)
	GetCustomersTotalEatenUniqueSnacks() ([]model.CustomerTotalEatenUniqueSnacks, error)
}