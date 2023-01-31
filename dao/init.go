package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,  //string类型默认字段长度
		DisableDatetimePrecision:  true, //禁止datetime精度，mysql5.6之前的数据库不支持
		DontSupportRenameIndex:    true, //重命名索引，就要把索引先删除再重建,mysql 5.7 不支持
		SkipInitializeWithVersion: false,
		DontSupportRenameColumn:   true, //用change重命名列，mysql 8 之前的数据库不支持
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)  //设置连接池
	sqlDB.SetMaxOpenConns(100) //打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	//主从配置
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connWrite)},                      //写操作
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, //读操作
			Policy:   dbresolver.RandomPolicy{},
		}))

	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
