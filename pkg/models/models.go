package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title           string             `bson:"title" json:"title"`
	Company         string             `bson:"company" json:"company"`
	Location        string             `bson:"location" json:"location"`
	Description     string             `bson:"description" json:"description"`
	Salary          string             `bson:"salary,omitempty" json:"salary,omitempty"`
	JobType         string             `bson:"job_type,omitempty" json:"job_type,omitempty"`
	ExperienceLevel string             `bson:"experience_level,omitempty" json:"experience_level,omitempty"`
	RequiredSkills  []string           `bson:"required_skills,omitempty" json:"required_skills,omitempty"`
	PostedDate      time.Time          `bson:"posted_date" json:"posted_date"`
	URL             string             `bson:"url" json:"url"`
	LastUpdated     time.Time          `bson:"last_updated" json:"last_updated"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username     string             `bson:"username" json:"username"`
	PasswordHash string             `bson:"password_hash" json:"-"`
	Email        string             `bson:"email" json:"email"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	LastLogin    time.Time          `bson:"last_login,omitempty" json:"last_login,omitempty"`
	Role         string             `bson:"role" json:"role"`
}
