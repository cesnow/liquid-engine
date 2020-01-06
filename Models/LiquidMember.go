package Models

import (
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateLiquidMember(member bson.M) (*mongo.InsertOneResult, error) {
	insertResult, err := LiquidSDK.GetServer().GetLiquidMemberCol().InsertOne(nil, member)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}
