package products

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt time.Time `json:"CreatedAt" bson:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" bson:"updated_at"`
	Price  float64 `json:"Price" bson:"
	title string `json:"title" bson:title"
}

func CreateProduct(c *fiber.Ctx) *primitive.Product {
	fmt.Println("Creating Product...")
	product := Product {
		ID: primitive.NewObjectID(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        Price: 5.99,
        title: "Pen",
    }

	if err := c.BodyParser(&product); err != nil { 
		return err 
	}

	client, err := db.GetMongoClient() 

	if err != nil {
		return err
	}

	collection := client.Database(db.Database).Collection.products
}

func GetAllProducts(c *fiber.Ctx) {

}