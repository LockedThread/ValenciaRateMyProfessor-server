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
	"server/schema"
	"time"
)

var (
	Client mongo.Client
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	Client = *client
	Setup()
}

func Setup() {
	// TODO: Setup some database setup methods
}

func GetProfessorById(id string) schema.Professor {
	collection, ctx := getCollection("professors")
	cursor, err := collection.Find(ctx, bson.M{
		"teacherId": id,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer cursor.Close(ctx)

	var professor schema.Professor

	if err = cursor.All(ctx, &professor); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("professor=%s", professor)
	return professor
}

func GetProfessors() []schema.Professor {
	collection, ctx := getCollection("professors")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	defer cursor.Close(ctx)

	var professors []schema.Professor

	if err = cursor.All(ctx, &professors); err != nil {
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
	return Client.Database("valencia-rate-my-professor")
}

func getCollection(collection string) (mongo.Collection, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return *getDatabaseFromDefault().Collection(collection), ctx
}
