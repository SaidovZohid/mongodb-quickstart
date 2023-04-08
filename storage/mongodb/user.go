package mongodb

import (
	"context"
	"errors"

	"github.com/SaidovZohid/mongodb-quickstart/pkg/mongodb"
	"github.com/SaidovZohid/mongodb-quickstart/storage/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserAlreadyExists = errors.New("user with such email already exists")
	ErrUserNotFound      = errors.New("user doesn't exists")
)

type userRepo struct {
	db *mongo.Collection
}

func NewUser(db *mongo.Database) repo.UserStorageI {
	return &userRepo{
		db: db.Collection(users),
	}
}

func (d *userRepo) Create(user *repo.User) error {
	res, err := d.db.InsertOne(context.Background(), user)
	if err != nil {
		if mongodb.IsDuplicate(err) {
			return ErrUserAlreadyExists
		}
		return err
	}

	user.Id = res.InsertedID.(primitive.ObjectID)

	return nil
}

func (d *userRepo) Get(id primitive.ObjectID) (*repo.User, error) {
	var user repo.User

	if err := d.db.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (d *userRepo) Delete(id primitive.ObjectID) error {
	_, err := d.db.DeleteOne(context.Background(), bson.M{"_id": id})

	return err
}
