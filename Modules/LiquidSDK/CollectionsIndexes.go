package LiquidSDK

func (server *LiquidServer) EnsureLiquidMemberIndexes() {
	server.GetDocDb().PopulateIndex(
		server.CodeName,
		ColNameLiquidMember,
		"account",
		-1,
		true,
	)
}

func (server *LiquidServer) EnsureLiquidUserIndexes() {
	server.GetDocDb().PopulateIndex(
		server.CodeName,
		ColNameLiquidUser,
		"auto_id",
		-1,
		true,
	)

	server.GetDocDb().PopulateMultiIndex(
		server.CodeName,
		ColNameLiquidUser,
		[]string{"from_id", "from_type"},
		[]int32{-1, -1},
		true,
	)
}

func (server *LiquidServer) EnsureLiquidAdminIndexes() {
	server.GetDocDb().PopulateIndex(
		server.CodeName,
		ColNameLiquidAdmin,
		"admin",
		1,
		true,
	)
}
