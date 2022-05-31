package repository

import (
	"github.com/pmylund/sortutil"
	"strconv"
	model "github.com/AndyMile/candyStore/models"
	"gorm.io/gorm"
)

var customers []model.Customer
var customersTotalEatenUniqueSnacks []model.CustomerTotalEatenUniqueSnacks

var topCustomers []model.TopCustomer

type StoreRepo struct {
	db *gorm.DB
}

func NewStoreRepo(db *gorm.DB) *StoreRepo {
	return &StoreRepo{
		db: db,
	}
}

func (r *StoreRepo) GetAll() ([]model.Customer, error) {
	r.db.Find(&customers)
	return customers, nil
}

func (r *StoreRepo) Get(id string) (model.Customer, error) {
	customerId, _ := strconv.Atoi(id)
	a := model.Customer{
		Id: customerId,
	}
	r.db.First(&a)
	return a, nil
}

//Customer with a total number of eaten unique snaks
func (r *StoreRepo) GetCustomersTotalEatenUniqueSnacks() ([]model.CustomerTotalEatenUniqueSnacks, error) {
	r.db.Table("customers").Select("SUM(eaten) as TotalEaten, name, candy").Group("name,candy").Find(&customersTotalEatenUniqueSnacks)
	return customersTotalEatenUniqueSnacks, nil
}

func (r *StoreRepo) GetTopCustomers() ([]model.TopCustomer, error) {
	customersTotalEatenUniqueSnacks, _ := r.GetCustomersTotalEatenUniqueSnacks()
	
	totalEatenWithFavouriteSnack := make(map[string]model.TotalEatenWithFavouriteSnack)
	//calculate total eaten snack and favourite snack
	for _, v := range customersTotalEatenUniqueSnacks {
		if _, ok := totalEatenWithFavouriteSnack[v.Name]; ok {
			c := totalEatenWithFavouriteSnack[v.Name]

			if totalEatenWithFavouriteSnack[v.Name].FavouriteSnack.Total < v.TotalEaten  {
				c.FavouriteSnack = model.FavouriteSnack {
					Total: v.TotalEaten,
					Candy: v.Candy,
				}
			}

			totalEatenWithFavouriteSnack[v.Name] = model.TotalEatenWithFavouriteSnack {
				Name: v.Name,
				TotalEaten: totalEatenWithFavouriteSnack[v.Name].TotalEaten + v.TotalEaten,
				FavouriteSnack: c.FavouriteSnack,
			}
		} else {
			totalEatenWithFavouriteSnack[v.Name] = model.TotalEatenWithFavouriteSnack {
				Name: v.Name,
				TotalEaten: v.TotalEaten,
				FavouriteSnack: model.FavouriteSnack {
					 Total: v.TotalEaten, 
					 Candy: v.Candy,
				},
			}
		}
	}

	//prepare slice of TopCustomer
	topCustomers = []model.TopCustomer{}
	for _, v := range totalEatenWithFavouriteSnack {
		customer := model.TopCustomer {
			Name: v.Name,
			FavouriteSnack: v.FavouriteSnack.Candy,
			TotalEaten: v.TotalEaten,
		}
		topCustomers = append(topCustomers, customer)
	}
	sortutil.DescByField(topCustomers, "TotalEaten")

	return topCustomers, nil
}