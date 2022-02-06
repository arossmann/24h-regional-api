package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/arossmann/24h-regional-api/entity"
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
	//fmt.Println("Mongo Connection: " + connectionURI)
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

	fmt.Println("Connected to MongoDB!")
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

func GetStoreByID(id primitive.ObjectID) (*entity.Store, error) {
	var store *entity.Store

	configDatabase := os.Getenv("MONGODB_DATABASE")
	configCollection := os.Getenv("MONGODB_COLLECTION")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database(configDatabase)
	collection := db.Collection(configCollection)
	result := collection.FindOne(ctx, bson.D{})
	if result == nil {
		return nil, errors.New("Could not find a store")
	}
	err := result.Decode(&store)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	log.Printf("Store: %v", store)
	return store, nil
}

/*func GetProducts()([]string, error){
	var products = []string
	configDatabase := os.Getenv("MONGODB_DATABASE")
	configCollection := os.Getenv("MONGODB_COLLECTION")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database(configDatabase)
	collection := db.Collection(configCollection)
	result, err := collection.Distinct(ctx,"products", bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range result {
		fmt.Println(value)
	}
	return result, nil

}*/

func Create(store *entity.Store) (primitive.ObjectID, error) {
	configDatabase := os.Getenv("MONGODB_DATABASE")
	configCollection := os.Getenv("MONGODB_COLLECTION")

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	store.ID = primitive.NewObjectID()

	result, err := client.Database(configDatabase).Collection(configCollection).InsertOne(ctx, store)
	if err != nil {
		log.Printf("could not create store: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func Update(store *entity.Store) (*entity.Store, error) {
	var updatedStore *entity.Store

	configDatabase := os.Getenv("MONGODB_DATABASE")
	configCollection := os.Getenv("MONGODB_COLLECTION")

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	update := bson.M{
		"$set": store,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}

	err := client.Database(configDatabase).Collection(configCollection).FindOneAndUpdate(ctx, bson.M{"_id": store.ID}, update, &opt).Decode(&updatedStore)
	if err != nil {
		log.Printf("Could not save Store: %v", err)
		return nil, err
	}
	return updatedStore, nil
}
