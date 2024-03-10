package LiquidDB

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const colNameLiquidMember string = "LIQUIDMember"
const colNameLiquidUser string = "LIQUIDUser"
const colNameLiquidUserData string = "LIQUIDUserData"
const colNameLiquidAdmin string = "LIQUIDAdmin"

func (db *LiquidDB) GetLiquidMemberCol() *mongo.Collection {
	MongoClient := db.GetDocDb().GetClient()
	return MongoClient.Database(db.CodeName).Collection(colNameLiquidMember)
}

func (db *LiquidDB) GetLiquidUserCol() *mongo.Collection {
	MongoClient := db.GetDocDb().GetClient()
	return MongoClient.Database(db.CodeName).Collection(colNameLiquidUser)
}

func (db *LiquidDB) GetLiquidUserDataCol() *mongo.Collection {
	MongoClient := db.GetDocDb().GetClient()
	return MongoClient.Database(db.CodeName).Collection(colNameLiquidUserData)
}

func (db *LiquidDB) GetLiquidAdminCol() *mongo.Collection {
	MongoClient := db.GetDocDb().GetClient()
	return MongoClient.Database(db.CodeName).Collection(colNameLiquidAdmin)
}
