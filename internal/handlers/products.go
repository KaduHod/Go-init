package handlers

import (
	"time"
	"fmt"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"api/internal/db"
)

type Product struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt time.Time `json:"CreatedAt" bson:"created_at"`
	UpdatedAt time.Time `json:"UpdatedAt" bson:"updated_at"`
	Price float32 `json:"Price" bson:Price"`
	title string `json:"title" bson:title"`
}

func CreateProduct(c *fiber.Ctx) error {
	fmt.Println("Creating Product...")

	product := Product {
		ID: primitive.NewObjectID(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        Price: 5.99,
        title: "Pen",
    }

	if err := c.BodyParser(&product); err != nil { 
		panic(err) 
	}

	fmt.Println(product)

	client, err := db.GetMongoClient() 

	if err != nil {
		panic(err)
	}

	collection := client.Database(db.Database).Collection(string(db.ProductsCollection))

	_, err = collection.InsertOne(context.TODO(), product)

	if err != nil {
		panic(err)
	}

	return c.JSON(product)
}

func GetAllProducts(c *fiber.Ctx) {

}