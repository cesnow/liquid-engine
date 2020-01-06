package LiquidSDK

import "go.mongodb.org/mongo-driver/mongo"

const ColNameLiquidMember string = "LIQUIDMember"
const ColNameLiquidUser string = "LIQUIDUser"
const ColNameLiquidGameUser string = "LIQUIDGameUser"
const ColNameLiquidAdmin string = "LIQUIDAdmin"

func (server *LiquidServer) GetLiquidMemberCol() *mongo.Collection {
	MongoClient := server.GetDocDb().GetClient()
	return MongoClient.Database(server.CodeName).Collection(ColNameLiquidMember)
}

func (server *LiquidServer) GetLiquidUserCol() *mongo.Collection {
	MongoClient := server.GetDocDb().GetClient()
	return MongoClient.Database(server.CodeName).Collection(ColNameLiquidUser)
}

func (server *LiquidServer) GetLiquidGameUserCol() *mongo.Collection {
	MongoClient := server.GetDocDb().GetClient()
	return MongoClient.Database(server.CodeName).Collection(ColNameLiquidGameUser)
}

func (server *LiquidServer) GetLiquidAdminCol() *mongo.Collection {
	MongoClient := server.GetDocDb().GetClient()
	return MongoClient.Database(server.CodeName).Collection(ColNameLiquidAdmin)
}
