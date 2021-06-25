package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"server/models"
)

var (
	Client mongo.Client
)

func Connect() *mongo.Client {
	fmt.Printf("Connecting to database\n")
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Connected to database\n")
	return client
}

func GetProfessorById(id string) models.Professor {
	collection := getCollection("professors")
	cursor := collection.FindOne(context.Background(), bson.M{
		"teacherId": id,
	})
	if cursor.Err() != nil {
		log.Fatalln(cursor.Err())
	}

	var professor models.Professor

	if err := cursor.Decode(&professor); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("professor=%s", professor)
	return professor
}

func GetProfessors() []models.Professor {
	collection := getCollection("professors")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}(cursor, context.Background())

	var professors []models.Professor

	if err = cursor.All(context.Background(), &professors); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("professors=%s", professors)
	return professors
}

// Utility methods
func getDatabase(database string) *mongo.Database {
	return Client.Database(database)
}

func getDatabaseFromDefault() *mongo.Database {
	return getDatabase("valencia-rate-my-professor")
}

func getCollection(collection string) mongo.Collection {
	return *getDatabaseFromDefault().Collection(collection)
}
