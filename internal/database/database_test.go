package database

import (
	"testing"
	"time"

	"github.com/paipeline/LinkedInScraper/pkg/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertJob(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		jobsCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		job := &models.Job{
			Title:       "Software Engineer",
			Company:     "Tech Corp",
			Location:    "Remote",
			Description: "Exciting opportunity for a software engineer",
			PostedDate:  time.Now(),
			URL:         "https://example.com/job",
		}

		err := InsertJob(job)
		assert.NoError(t, err)
	})

	mt.Run("error", func(mt *mtest.T) {
		jobsCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		job := &models.Job{
			Title:       "Software Engineer",
			Company:     "Tech Corp",
			Location:    "Remote",
			Description: "Exciting opportunity for a software engineer",
			PostedDate:  time.Now(),
			URL:         "https://example.com/job",
		}

		err := InsertJob(job)
		assert.Error(t, err)
	})
}

func TestFindJobByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		jobsCollection = mt.Coll
		expectedJob := models.Job{
			ID:          primitive.NewObjectID(),
			Title:       "Software Engineer",
			Company:     "Tech Corp",
			Location:    "Remote",
			Description: "Exciting opportunity for a software engineer",
			PostedDate:  time.Now(),
			URL:         "https://example.com/job",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedJob.ID},
			{Key: "title", Value: expectedJob.Title},
			{Key: "company", Value: expectedJob.Company},
			{Key: "location", Value: expectedJob.Location},
			{Key: "description", Value: expectedJob.Description},
			{Key: "posted_date", Value: expectedJob.PostedDate},
			{Key: "url", Value: expectedJob.URL},
		}))

		job, err := FindJobByID(expectedJob.ID.Hex())
		assert.NoError(t, err)
		assert.Equal(t, expectedJob.ID, job.ID)
		assert.Equal(t, expectedJob.Title, job.Title)
		assert.Equal(t, expectedJob.Company, job.Company)
	})

	mt.Run("not found", func(mt *mtest.T) {
		jobsCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))

		job, err := FindJobByID(primitive.NewObjectID().Hex())
		assert.Error(t, err)
		assert.Nil(t, job)
	})
}

func TestInsertUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		usersCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		user := &models.User{
			Username:     "testuser",
			PasswordHash: "hashedpassword",
			Email:        "test@example.com",
			CreatedAt:    time.Now(),
			Role:         "user",
		}

		err := InsertUser(user)
		assert.NoError(t, err)
	})

	mt.Run("error", func(mt *mtest.T) {
		usersCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		user := &models.User{
			Username:     "testuser",
			PasswordHash: "hashedpassword",
			Email:        "test@example.com",
			CreatedAt:    time.Now(),
			Role:         "user",
		}

		err := InsertUser(user)
		assert.Error(t, err)
	})
}

func TestFindUserByUsername(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		usersCollection = mt.Coll
		expectedUser := models.User{
			ID:           primitive.NewObjectID(),
			Username:     "testuser",
			PasswordHash: "hashedpassword",
			Email:        "test@example.com",
			CreatedAt:    time.Now(),
			Role:         "user",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedUser.ID},
			{Key: "username", Value: expectedUser.Username},
			{Key: "password_hash", Value: expectedUser.PasswordHash},
			{Key: "email", Value: expectedUser.Email},
			{Key: "created_at", Value: expectedUser.CreatedAt},
			{Key: "role", Value: expectedUser.Role},
		}))

		user, err := FindUserByUsername(expectedUser.Username)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, expectedUser.Username, user.Username)
		assert.Equal(t, expectedUser.Email, user.Email)
	})

	mt.Run("not found", func(mt *mtest.T) {
		usersCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))

		user, err := FindUserByUsername("nonexistentuser")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
