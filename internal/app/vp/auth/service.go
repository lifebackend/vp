package auth

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DeviceNotFound = "device not found"

type Service struct {
	collection *mongo.Collection
}

type Auth struct {
	DeviceID string `bson:"deviceID"`
	Password string `bson:"password"`
}

func NewService(db *mongo.Database) *Service {
	return &Service{
		collection: db.Collection("users"),
	}
}

func (s *Service) Check(ctx context.Context, deviceID string, password string) error {
	r := s.collection.FindOne(ctx, bson.D{{"deviceID", deviceID}, {"password", password}})

	var auth Auth

	err := r.Decode(&auth)
	if err != nil {
		return err
	}

	if auth.DeviceID != deviceID {
		return errors.New(DeviceNotFound)
	}

	return nil
}
