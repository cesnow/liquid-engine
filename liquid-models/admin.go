package LiquidModels

import (
	"github.com/cesnow/liquid-engine/liquid-db"
	"github.com/cesnow/liquid-engine/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

type LiquidAdminConfCounterSetting struct {
	Admin   string `json:"admin" bson:"admin"`
	Counter int    `json:"counter" bson:"counter"`
}

const initAutoId = 1000000

func GetAutoID() string {
	filter := bson.M{"admin": "auto_id"}
	update := bson.M{"$inc": bson.M{"counter": 1}}

	var newAutoIdDoc *LiquidAdminConfCounterSetting
	newAutoIdDocErr := LiquidDB.GetLiquidAdminCol().FindOneAndUpdate(
		nil,
		filter,
		update,
		options.FindOneAndUpdate().SetUpsert(true),
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&newAutoIdDoc)
	if newAutoIdDocErr != nil {
		logger.SysLog.Errorf("[CMD][Login] Get New Auto ID Failed, %s", newAutoIdDocErr)
		return ""
	}

	uid := strconv.Itoa(initAutoId + newAutoIdDoc.Counter)
	logger.SysLog.Infof("[CMD][Login] New Auto ID Created : UID(%s)", uid)
	return uid
}
