package dao

import (
	"fmt"
	"gin_shop/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Address{},
			&model.Category{},
			&model.Carousel{},
			&model.Cart{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.Product{},
			&model.ProductImg{},
			&model.Admin{},
			&model.BasePage{})
	if err != nil {
	}
	fmt.Println("err ", err)
	return
}
