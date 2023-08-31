package model

//执行数据迁移

func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&BaseInfoHis{})
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&ChildrenInfo{})
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&Kinship{})
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&TestHistory{})

}
