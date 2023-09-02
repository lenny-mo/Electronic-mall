package dao

import "eletronicMall/model"

// migration 数据库迁移
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Carousel{},
		&model.Category{},
		&model.Notice{},
		&model.Favorite{},
		&model.Order{},
		&model.ProductImg{},
		&model.Cart{},
		&model.Address{},
		&model.Admin{},
		&model.BasePage{},
	)
	if err != nil {
		panic(err)
	}
}
