package lib

import (
	"fmt"

	"github.com/fatih/color"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID    uint `gorm:"unique;primaryKey;autoIncrement"`
	Code  string
	Price uint
}

func GormTest() {

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	db, err := gorm.Open(sqlite.Open("gormTest.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	products := []*Product{
		{Code: "K-009", Price: 18},
		{Code: "K-008", Price: 19},
	}

	db.Create(products)

	// Read
	var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	fmt.Printf("  Remote: %s -- %s \n", green(product.Code), yellow(product.Price))

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Printf("  Remote: %s -- %s \n", green(product.Code), yellow(product.Price))

	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	fmt.Printf("  Remote: %s -- %s \n", green(product.Code), yellow(product.Price))

	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	fmt.Printf("  Remote: %s -- %s \n", green(product.Code), yellow(product.Price))

	// Get all records
	result := db.Find(&products)
	// SELECT * FROM users;
	fmt.Printf("  Result: %s -- %s \n", green(result.RowsAffected), yellow(result.Error))

	for _, prod := range products {
		fmt.Printf("  Remote         : %s --> %s -- %s \n", red(prod.ID), green(prod.Code), yellow(prod.Price))
	}

	// Delete - delete product
	db.Delete(&product, 1)
}
