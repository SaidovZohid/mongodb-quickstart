package storage

import (
	"github.com/SaidovZohid/mongodb-quickstart/storage/mongodb"
	"github.com/SaidovZohid/mongodb-quickstart/storage/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type StorageI interface {
	User() repo.UserStorageI
}

type StoragePg struct {
	userRepo repo.UserStorageI
}

func NewStorage(db *mongo.Database) StorageI {
	return &StoragePg{
		userRepo: mongodb.NewUser(db),
	}
}

func (s *StoragePg) User() repo.UserStorageI {
	return s.userRepo
}