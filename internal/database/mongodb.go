package database

import (
	"context"
	"time"

	"github.com/ppeng24/LinkedInScraper/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client          *mongo.Client
	jobsCollection  *mongo.Collection
	usersCollection *mongo.Collection
)

func InitMongoDB(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	db := client.Database("linkedin_scraper")
	jobsCollection = db.Collection("jobs")
	usersCollection = db.Collection("users")

	return nil
}

func InsertJob(job *models.Job) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := jobsCollection.InsertOne(ctx, job)
	return err
}

func FindJobByID(id string) (*models.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var job models.Job
	objID, _ := primitive.ObjectIDFromHex(id)
	err := jobsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func InsertUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := usersCollection.InsertOne(ctx, user)
	return err
}

func FindUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
