package Models

import (
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

type LiquidAdminConfCounterSetting struct {
	Admin   string `json:"admin" bson:"admin"`
	Counter int    `json:"counter" bson:"counter"`
}

const initAutoId = 1000001

func GetAutoID() string {
	filter := bson.M{"admin": "auto_id"}
	update := bson.M{"$inc": bson.M{"counter": 1}}

	var newAutoIdDoc *LiquidAdminConfCounterSetting
	newAutoIdDocErr := LiquidSDK.GetServer().GetLiquidAdminCol().FindOneAndUpdate(
		nil,
		filter,
		update,
		options.FindOneAndUpdate().SetUpsert(true),
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&newAutoIdDoc)
	if newAutoIdDocErr != nil {
		Logger.SysLog.Errorf("[CMD][Login] Get New Auto ID Failed, %s", newAutoIdDocErr)
		return ""
	}

	uid := strconv.Itoa(initAutoId + newAutoIdDoc.Counter)
	Logger.SysLog.Infof("[CMD][Login] New Auto ID Created : UID(%s)", uid)
	return uid
}
