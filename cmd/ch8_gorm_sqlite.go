package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func action(stepName string, product Product) {
	fmt.Println(stepName, product.Code, product.Price)
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Creat
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	// 官网 demo 此处执行报错，提示无法找数据
	db.First(&product, 1)
	action("Read by index", product)

	db.First(&product, "code=?", "D42")
	action("Read by code index", product)

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	action("update price to 200", product)

	// Update 多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	action("updates", product)
	db.Model(&product).Updates(map[string]interface{}{"Price": 300, "Code": "A42"})
	action("updates", product)

	// delete
	db.Delete(&product, "code=?", product.Code)
	// 执行数据语句后，product 对象依旧保留数据，效果和执行 update 不同，执行 update 会同步更新实际的 product 对象
	action("delete", product)
}
