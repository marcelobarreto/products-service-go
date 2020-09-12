package main

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, dbError = gorm.Open(sqlite.Open("development.db"), &gorm.Config{})

type Store struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Open bool   `gorm:"default:true"`
}

type StoreSerialized struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Open bool   `gorm:"default:true"`
}

type Product struct {
	gorm.Model
	ID        uint      `json:"id",gorm:"primaryKey"`
	Code      string    `json:"code",gorm:"index"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Discount  float64   `json:"discount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	StoreID   uint      `json:"store_id",gorm:"index"`
	Store     Store     `json:"store"`
}

type ProductSerialized struct {
	ID       uint            `json:"id"`
	Code     string          `json:"code"`
	Name     string          `json:"name"`
	Price    float64         `json:"price"`
	Discount float64         `json:"discount"`
	StoreID  uint            `json:"store_id"`
	Store    StoreSerialized `json:"store"`
}

// TableName - Sets StoreSerialized TableName
func (StoreSerialized) TableName() string {
	return "stores"
}

// TableName - Sets ProductSerialized TableName
func (ProductSerialized) TableName() string {
	return "products"
}

func (p *ProductSerialized) Save() *ProductSerialized {
	db.Create(p)
	db.Table("stores").Find(&p.Store, p.StoreID)
	return p
}

func (p Product) Find(id string) *ProductSerialized {
	var product ProductSerialized
	result := db.Preload("Store").First(&product, id)
	err := result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &product
}

func main() {
	app := fiber.New()
	var store StoreSerialized
	db.FirstOrCreate(&store, Store{Name: "Foo Store", Open: true})

	if dbError != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Product{})

	app.Get("/products", func(c *fiber.Ctx) {
		var products []ProductSerialized
		db.Preload("Store").Find(&products)
		c.Status(200).JSON(products)
	})

	app.Get("/products/:id", func(c *fiber.Ctx) {
		product := Product.Find(Product{}, c.Params("id"))
		if product != nil {
			c.Status(200).JSON(product)
		} else {
			c.Status(404)
		}
	})

	app.Post("/products", func(c *fiber.Ctx) {
		product := &ProductSerialized{}
		if err := c.BodyParser(product); err != nil {
			log.Println(err)
			c.Status(422).JSON(map[string]interface{}{"error": "Could not create a new product"})
			return
		}
		product.Save()
		c.Status(201).JSON(product)
	})

	app.Put("/products/:id", func(c *fiber.Ctx) {
		if product := Product.Find(Product{}, c.Params("id")); product != nil {
			if err := c.BodyParser(product); err != nil {
				log.Println(err)
				c.Status(422).JSON(map[string]interface{}{"error": "Could not update product"})
				return
			}
			db.Save(&product)
			c.Status(200).JSON(product)
		} else {
			c.Status(404)
		}
	})

	app.Delete("/products/:id", func(c *fiber.Ctx) {
		db.Unscoped().Delete(&Product{}, c.Params("id"))
		c.Status(200)
	})

	app.Listen(3000)
}
