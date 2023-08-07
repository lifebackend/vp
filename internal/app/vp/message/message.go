package message

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mapTypesRegExp map[string]*regexp.Regexp
	steps          map[string]string
	mapFields      map[string][]string
)

type Service struct {
	collection *mongo.Collection
	logger     *logrus.Entry
}

func NewService(db *mongo.Database, logger *logrus.Entry) *Service {
	collection := db.Collection("messages")
	return &Service{collection: collection, logger: logger}
}

func (s *Service) Save(ctx context.Context, deviceID string, from string, typeMsg string, msg string) error {
	t, m := s.parseMessage(msg)

	if ok, tp := isTwoStep(t); ok {
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
		}
		set := bson.M{}
		if _, ok := m["balance"]; ok {
			set = bson.M{"$set": bson.M{"msg.balance": m["balance"]}}
		}

		if _, ok := m["from"]; ok {
			set = bson.M{"$set": bson.M{"msg.from": m["from"]}}
		}

		err := s.collection.
			FindOneAndUpdate(
				ctx,
				bson.D{
					{"deviceID", deviceID},
					{"typeMsg", typeMsg},
					{"type", tp},
					{"msg.amount", m["amount"]},
				}, set, &opt).Err()

		if !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}

	}

	b := bson.D{
		{"deviceID", deviceID},
		{"from", from},
		{"typeMsg", typeMsg},
		{"msg", m},
		{"type", t},
		{"datetime", time.Now().Unix()},
	}

	_, err := s.collection.InsertOne(ctx, b)

	return err
}

func isTwoStep(tp string) (bool, string) {
	stp, ok := steps[tp]
	return ok, stp
}

func (s *Service) mapDataToField(m map[string]string, data []string, fields []string) {
	var mx sync.Mutex

	mx.Lock()

	for i, f := range fields {
		s.logger.Info("map", "field:", f, "data:", data[i])
		m[f] = data[i]
	}
	mx.Unlock()
}

func getFieldsByType(tp string) []string {
	if r, ok := mapFields[tp]; ok {
		return r
	}
	return nil
}

// sanitize push message
func prepareMessage(msg string) string {
	msg = strings.ReplaceAll(msg, "\n", "")
	msg = strings.ReplaceAll(msg, "\u00a0", " ")
	return msg
}

func (s *Service) parseMessage(msg string) (string, map[string]string) {
	msg = prepareMessage(msg)

	var mx sync.Mutex
	m := make(map[string]string)
	for k, r := range mapTypesRegExp {
		if ok := r.MatchString(msg); ok {
			data := r.FindAllStringSubmatch(msg, -1)
			fields := getFieldsByType(k)

			s.mapDataToField(m, data[0], fields)

			return k, m
		}
	}
	mx.Lock()
	m["body"] = msg
	mx.Unlock()
	return TypeUnknown, m
}
