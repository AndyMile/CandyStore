package model

type Customer struct {
	Id int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Candy string `json:"candy"`
	Eaten int64 `json:"eaten"`
}

type CustomerTotalEatenUniqueSnacks struct {
	Name string `json:"name"`
	Candy string `json:"candy"`
	TotalEaten int64 `json:"TotalEaten"`
}

type TopCustomer struct {
	Name string 
	FavouriteSnack string
	TotalEaten int64
}

type TotalEatenWithFavouriteSnack struct {
	Name string
	TotalEaten int64
	FavouriteSnack FavouriteSnack
}

type FavouriteSnack struct {
	Candy string
	Total int64
}