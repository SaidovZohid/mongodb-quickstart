package main

import (
	"github.com/SaidovZohid/mongodb-quickstart/config"
	"github.com/SaidovZohid/mongodb-quickstart/pkg/logger"
	"github.com/SaidovZohid/mongodb-quickstart/pkg/mongodb"
	"github.com/SaidovZohid/mongodb-quickstart/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// initializing logger
	logger.Init()
	log := logger.GetLogger()

	// intiliazing config
	cfg := config.Load(".")

	// making client to mongodb
	mongoClient, err := mongodb.NewClient(cfg.MongoDB.Url, cfg.MongoDB.Username, cfg.MongoDB.Password)
	if err != nil {
		log.Fatal("Error while making client to mongodb", err)
	}

	db := mongoClient.Database(cfg.MongoDB.Database)

	strg := storage.NewStorage(db)

	// user := &repo.User{
	// 	UserInfo: repo.UserInfo{
	// 		FirstName:   "Zohid",
	// 		LastName:    "Saidov",
	// 		PhoneNumber: "+998931153336",
	// 		Address:     "Abay ko'cha-2",
	// 	},
	// 	UserCourses: []repo.UserCourse{
	// 		{
	// 			ID:          primitive.NewObjectIDFromTimestamp(time.Now()),
	// 			Name:        "PHP Backend",
	// 			Code:        "4321",
	// 			Description: "PHP Backend Basics",
	// 			Color:       "Green",
	// 			ImageURL:    "zohiddev.uz/images/image.png",
	// 			CreatedAt:   time.Now(),
	// 			UpdatedAt:   time.Now(),
	// 			Published:   false,
	// 		},
	// 		{
	// 			ID:          primitive.NewObjectIDFromTimestamp(time.Now()),
	// 			Name:        "Golang Backend",
	// 			Code:        "1234",
	// 			Description: "Golang Backend Basics",
	// 			Color:       "Brown",
	// 			CreatedAt:   time.Now(),
	// 			UpdatedAt:   time.Now(),
	// 			Published:   true,
	// 		},
	// 		{
	// 			ID:          primitive.NewObjectIDFromTimestamp(time.Now()),
	// 			Name:        "Java Backend",
	// 			Code:        "4312",
	// 			Description: "Java Backend Basics",
	// 			Color:       "Red",
	// 			ImageURL:    "zohiddev.uz/images/image2.png",
	// 			CreatedAt:   time.Now(),
	// 			UpdatedAt:   time.Now(),
	// 			Published:   true,
	// 		},
	// 	},
	// }
	// err = strg.User().Create(user)
	// id, err := primitive.ObjectIDFromHex("6430efc30e3c9778d94635a1")
	// user, err := strg.User().Get(id)

	id, err := primitive.ObjectIDFromHex("6430efc30e3c9778d94635a1")
	if err != nil {
		log.Error("Failed: ", err)
	}
	
	err = strg.User().Delete(id)
	if err != nil {
		log.Error("Failed: ", err)
	}
}
