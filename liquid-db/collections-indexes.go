package LiquidDB

func (db *LiquidDB) EnsureLiquidMemberIndexes() {
	db.GetDocDb().PopulateMultiIndex(
		db.CodeName,
		colNameLiquidMember,
		[]string{"account", "from_type"},
		[]int32{-1, -1},
		true,
	)
}

func (db *LiquidDB) EnsureLiquidUserIndexes() {
	db.GetDocDb().PopulateIndex(
		db.CodeName,
		colNameLiquidUser,
		"auto_id",
		-1,
		true,
	)

	db.GetDocDb().PopulateMultiIndex(
		db.CodeName,
		colNameLiquidUser,
		[]string{"from_id", "from_type"},
		[]int32{-1, -1},
		true,
	)
}

func (db *LiquidDB) EnsureLiquidAdminIndexes() {
	db.GetDocDb().PopulateIndex(
		db.CodeName,
		colNameLiquidAdmin,
		"admin",
		1,
		true,
	)
}
