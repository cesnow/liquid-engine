package LiquidModels

import (
	"fmt"
	LiquidDb "github.com/cesnow/liquid-engine/liquid-db"
	"github.com/cesnow/liquid-engine/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func createDefaultUserData(autoId string) {
	dateTime := time.Now()
	mainKey := bson.M{"auto_id": autoId}
	defaultUserData := bson.M{
		"$set": bson.M{
			"player_data": bson.M{},
			"create":      dateTime,
			"update":      dateTime,
		},
	}
	_, createDefaultError := LiquidDb.GetLiquidUserDataCol().UpdateOne(
		nil,
		mainKey,
		defaultUserData,
		options.Update().SetUpsert(true),
	)

	if createDefaultError != nil {
		logger.SysLog.Warnf("[CMD][Login] Create Default User Data Failed, %s", createDefaultError)
	}
}

func SetUserData(autoId string, key string, value string) {
	dateTime := time.Now()
	mainKey := bson.M{"auto_id": autoId}
	setUserData := bson.M{
		"$set": bson.M{
			key:      value,
			"update": dateTime,
		},
	}
	_, updateErr := LiquidDb.GetLiquidUserDataCol().UpdateOne(
		nil,
		mainKey,
		setUserData,
		options.Update().SetUpsert(true),
	)
	if updateErr != nil {
		logger.SysLog.Warnf("Set User Data Failed, %s", updateErr)
	}
}

func GetUserData(autoId string, key string) (string, error) {
	mainKey := bson.M{"auto_id": autoId, key: bson.M{"exists": true}}
	var value bson.M
	findOpts := options.FindOne().SetProjection(bson.M{key: true})
	fetchErr := LiquidDb.GetLiquidUserDataCol().FindOne(nil, mainKey, findOpts).Decode(&value)
	if fetchErr != nil {
		return "", fmt.Errorf("can't find key from user data [AutoId: %s]", autoId)
	}
	return value[key].(string), nil
}

func CheckDefaultPlayerData(autoId string) {
	findUserFilter := bson.M{"auto_id": autoId}
	findErr := LiquidDb.GetLiquidUserDataCol().FindOne(
		nil,
		findUserFilter,
		options.FindOne().SetProjection(bson.D{
			{"auto_id", true},
			{"_id", true},
		}),
	)
	if findErr != nil {
		createDefaultUserData(autoId)
	}
}
