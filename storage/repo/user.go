package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStorageI interface {
	Create(u *User) error
	Delete(id primitive.ObjectID) error
	Get(id primitive.ObjectID) (*User, error)
}

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	UserInfo    UserInfo           `bson:"user_info"`
	UserCourses []UserCourse       `bson:"user_courses"`
}

type UserInfo struct {
	FirstName   string `bson:"first_name"`
	LastName    string `bson:"last_name"`
	PhoneNumber string `bson:"phone_number"`
	Address     string `bson:"address"`
}

type UserCourse struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Code        string             `bson:"code,omitempty"`
	Description string             `bson:"description,omitempty"`
	Color       string             `bson:"color,omitempty"`
	ImageURL    string             `bson:"imageUrl"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty"`
	Published   bool               `bson:"published,omitempty"`
}
