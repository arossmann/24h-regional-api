package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/arossmann/24h-regional-api/entity"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/arossmann/24h-regional-api/entity"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb+srv://%s:%s@%s"
)

// GetConnection - Retrieves a client to the DocumentDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")
	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}
	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}
	return client, ctx, cancel
}

func GetAllStores() ([]*entity.Store, error) {
	var stores []*entity.Store
	configDatabase := os.Getenv("MONGODB_DATABASE")
	configCollection := os.Getenv("MONGODB_COLLECTION")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database(configDatabase)
	collection := db.Collection(configCollection)
	curser, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer curser.Close(ctx)
	err = curser.All(ctx, &stores)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	return stores, nil
}

func GetStoreByID(id string) (*entity.Store, error) {
	var store *entity.Store
	configDatabase := os.Getenv("MONGODB_DATABASE")
	configCollection := os.Getenv("MONGODB_COLLECTION")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	collection := client.Database(configDatabase).Collection(configCollection)
	objectId, _ := primitive.ObjectIDFromHex(id)
	result := collection.FindOne(ctx, bson.M{"_id": objectId})
	if result == nil {
		return nil, errors.New("could not find a store")
	}
	err := result.Decode(&store)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	return store, nil
}

func Create(c *fiber.Ctx) error {
	configDatabase := os.Getenv("MONGODB_DATABASE")
	configCollection := os.Getenv("MONGODB_COLLECTION")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	newStore := new(entity.Store)
	if err := c.BodyParser(newStore); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	// force creation of new ObjectID
	newStore.ID = ""
	insertionResult, err := client.Database(configDatabase).Collection(configCollection).InsertOne(c.Context(), newStore)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := client.Database(configDatabase).Collection(configCollection).FindOne(c.Context(), filter)
	createdStore := &entity.Store{}
	createdRecord.Decode(createdStore)
	return c.Status(201).JSON(createdStore)
}
func Delete(c *fiber.Ctx) error {
	configDatabase := os.Getenv("MONGODB_DATABASE")
	configCollection := os.Getenv("MONGODB_COLLECTION")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	storeID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}
	query := bson.D{{Key: "_id", Value: storeID}}
	result, err := client.Database(configDatabase).Collection(configCollection).DeleteOne(c.Context(), query)
	if err != nil {
		return c.SendStatus(500)
	}
	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}
	return c.Status(201).SendString("Deletion of Store with ID " + c.Params("id") + ": successful")
}

func Update(c *fiber.Ctx) error {
	configDatabase := os.Getenv("MONGODB_DATABASE")
	configCollection := os.Getenv("MONGODB_COLLECTION")

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	idParam := c.Params("id")
	storeID, err := primitive.ObjectIDFromHex(idParam)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	store := new(entity.Store)
	// Parse body into struct
	if err := c.BodyParser(store); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Find the employee and update its data
	query := bson.D{{Key: "_id", Value: storeID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: store.Name},
				{Key: "open", Value: store.Open},
				{Key: "products", Value: store.Products},
			},
		},
	}
	err = client.Database(configDatabase).Collection(configCollection).FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	// return the updated employee
	store.ID = idParam
	return c.Status(200).JSON(store)
}
