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
	"strings"
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
	var professors []models.Professor

	var cursor *mongo.Cursor
	var err error
	for _, collectionName := range getAllCollectionsContaining("professors") {
		collection := getCollection(collectionName)
		cursor, err = collection.Find(context.Background(), bson.M{})
		if err != nil {
			log.Fatalln(err)
		}

		var currentProfessors []models.Professor
		if err = cursor.All(context.Background(), &currentProfessors); err != nil {
			log.Fatalln(err)
		}
		for _, professor := range currentProfessors {
			fmt.Printf("professor=%s\n", professor)
		}
		professors = append(professors, currentProfessors...)
	}
	defer cursor.Close(context.Background())

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

func getAllCollectionsContaining(containing string) []string {
	names, err := getDatabaseFromDefault().ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		log.Fatalln(err)
	}
	var collectionNames []string
	for _, name := range names {
		if strings.Contains(name, containing) {
			collectionNames = append(collectionNames, name)
		}
	}
	return collectionNames
}
