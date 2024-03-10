package LiquidModels

import (
	LiquidDB "github.com/cesnow/liquid-engine/liquid-db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateLiquidMember(member bson.M) (*mongo.InsertOneResult, error) {
	insertResult, err := LiquidDB.GetLiquidMemberCol().InsertOne(nil, member)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}
