package Models

import (
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func CreateDefaultPlayerData(autoId string) {
	dateTime := time.Now()
	mainKey := bson.M{"auto_id": autoId}
	defaultGameUserData := bson.M{
		"$set": bson.M{
			"player_data": bson.M{},
			"create":      dateTime,
			"update":      dateTime,
		},
	}
	_, createDefaultError := LiquidSDK.GetServer().GetLiquidGameUserCol().UpdateOne(
		nil,
		mainKey,
		defaultGameUserData,
		options.Update().SetUpsert(true),
	)

	if createDefaultError != nil {
		Logger.SysLog.Warnf("[CMD][Login] Create Default Game User Data Failed, %s", createDefaultError)
	}
}

func CheckDefaultPlayerData(autoId string) {
	findUserFilter := bson.M{"auto_id": autoId}
	findErr := LiquidSDK.GetServer().GetLiquidGameUserCol().FindOne(
		nil,
		findUserFilter,
		options.FindOne().SetProjection(bson.D{
			{"auto_id", true},
			{"_id", true},
		}),
	)
	if findErr != nil {
		CreateDefaultPlayerData(autoId)
	}
}
