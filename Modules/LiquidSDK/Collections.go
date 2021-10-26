package LiquidSDK

import "go.mongodb.org/mongo-driver/mongo"

const ColNameLiquidMember string = "LIQUIDMember"
const ColNameLiquidUser string = "LIQUIDUser"
const ColNameLiquidUserData string = "LIQUIDUserData"
const ColNameLiquidAdmin string = "LIQUIDAdmin"

func (server *LiquidServer) GetLiquidMemberCol() *mongo.Collection {
	MongoClient := server.GetDocDb().GetClient()
	return MongoClient.Database(server.CodeName).Collection(ColNameLiquidMember)
}

func (server *LiquidServer) GetLiquidUserCol() *mongo.Collection {
	MongoClient := server.GetDocDb().GetClient()
	return MongoClient.Database(server.CodeName).Collection(ColNameLiquidUser)
}

func (server *LiquidServer) GetLiquidUserDataCol() *mongo.Collection {
	MongoClient := server.GetDocDb().GetClient()
	return MongoClient.Database(server.CodeName).Collection(ColNameLiquidUserData)
}

func (server *LiquidServer) GetLiquidAdminCol() *mongo.Collection {
	MongoClient := server.GetDocDb().GetClient()
	return MongoClient.Database(server.CodeName).Collection(ColNameLiquidAdmin)
}
