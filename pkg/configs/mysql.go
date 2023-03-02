package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

var DB *gorm.DB

func initDB() {
	var err error
	dsn := "root:mini-douyin@tcp(127.0.0.1:3306)/mini_douyin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Error("Mysql Connect Fail" + err.Error())
		panic(err)
	}

	// 分表中间件
	var sharedCount uint = 3
	initSharedTable(int(sharedCount))
	DB.Use(sharding.Register(sharding.Config{
		ShardingKey:    "user_id",
		NumberOfShards: sharedCount,
		// PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "relations"))

	Logger.Info("Mysql Connect and Init Shared Table Success.")
}

func initSharedTable(num int) {
	tableName := "relations"
	for i := 0; i < num; i++ {
		table := fmt.Sprintf("%s_%d", tableName, i)
		DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + ` (
			id bigint PRIMARY KEY,
			user_id bigint,
			target_id bigint,
			is_active int(1)
		)`)
	}
}
