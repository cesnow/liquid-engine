package LiquidModels

import (
	LiquidDb "github.com/cesnow/liquid-engine/liquid-db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"strconv"
	"time"
)

type LiquidUser struct {
	AutoId      string `json:"auto_id" bson:"auto_id"`
	InviteCode  string `json:"invite_code" bson:"invite_code"`
	Fingerprint string `json:"fingerprint" bson:"fingerprint"`

	FromType  string `json:"from_type" bson:"from_type"`
	FromId    string `json:"from_id" bson:"from_id"`
	FromToken string `json:"from_token" bson:"from_token"`

	IsDeactivate bool `json:"is_deactivate" bson:"is_deactivate"`

	Create time.Time `json:"create" bson:"create"`
	Update time.Time `json:"update" bson:"update"`
}

func FindLiquidGuestUser(autoId string) *LiquidUser {
	filter := bson.M{"auto_id": autoId}
	var user *LiquidUser
	findUserErr := LiquidDb.GetLiquidUserCol().FindOne(nil, filter).Decode(&user)
	if findUserErr != nil {
		return nil
	}
	return user
}

func CreateLiquidUser(fromType string, fromId string) *LiquidUser {
	autoId := GetAutoID()

	if autoId == "" {
		return nil
	}

	dateTime := time.Now()
	inviteCode := getAutoIdToInviteCode(autoId)
	fingerprint := ""

	if fromId == "" {
		fromId = autoId
	}

	liquidUser := &LiquidUser{
		AutoId:       autoId,
		InviteCode:   inviteCode,
		Fingerprint:  fingerprint,
		FromType:     fromType,
		FromId:       fromId,
		IsDeactivate: false,
		Create:       dateTime,
		Update:       dateTime,
	}

	_, insertErr := LiquidDb.GetLiquidUserCol().InsertOne(nil, liquidUser)
	if insertErr != nil {
		return nil
	}

	createDefaultUserData(autoId)

	liquidUser = FindLiquidGuestUser(autoId)
	return liquidUser
}

func FindLiquidUserFromType(fromType string, fromId string) *LiquidUser {
	filter := bson.M{"from_id": fromId, "from_type": fromType}
	var user *LiquidUser
	findUserErr := LiquidDb.GetLiquidUserCol().FindOne(nil, filter).Decode(&user)
	if findUserErr != nil {
		return nil
	}
	return user
}

func FindLiquidUserByAutoId(AutoId string, InviteCode string) *LiquidUser {
	filter := bson.M{"auto_id": AutoId, "invite_code": InviteCode}
	var user *LiquidUser
	findUserErr := LiquidDb.GetLiquidUserCol().FindOne(nil, filter).Decode(&user)
	if findUserErr != nil {
		return nil
	}
	return user
}

func BindLiquidUser(AutoId string, FromId string, FromType string, FromToken string) (*mongo.UpdateResult, error) {
	filter := bson.M{"auto_id": AutoId}
	setBindData := bson.M{
		"$set": bson.M{
			"from_id":    FromId,
			"from_type":  FromType,
			"from_token": FromToken,
		},
	}
	setBindResult, setBindResultErr := LiquidDb.GetLiquidUserCol().UpdateOne(
		nil,
		filter,
		setBindData,
	)
	if setBindResultErr != nil {
		return nil, setBindResultErr
	}
	return setBindResult, nil
}

var convertTable = [...]string{
	"U", "V", "W", "X", "Y",
	"A", "B", "C", "D", "E",
	"F", "G", "H", "I", "J",
	"L", "M", "N", "Z", "K",
	"P", "Q", "R", "S", "T",
}

func getAutoIdToInviteCode(autoId string) string {
	inviteCode := ""
	InviteCodeList := make([]int, 0)
	autoIdInt, _ := strconv.ParseInt(autoId, 10, 64)
	rand.Seed(autoIdInt + time.Now().UnixNano())
	newId := strconv.FormatInt(time.Now().Unix()+rand.Int63n(time.Now().Unix()), 10)
	newId = newId + autoId
	newId2Int, _ := strconv.Atoi(newId[len(newId)-13:])
	resultToBase := decimalToBase(InviteCodeList, newId2Int)
	for _, base := range resultToBase {
		inviteCode += convertTable[base]
	}
	return inviteCode
}

func decimalToBase(baseList []int, decimal int) []int {
	base := len(convertTable)
	baseList = append(baseList, decimal%base)
	div := decimal / base
	if div == 0 {
		return baseList
	}
	return decimalToBase(baseList, div)
}
