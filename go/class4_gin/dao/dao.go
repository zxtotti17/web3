package dao

import (
	"class4_gin/models"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := &models.MysqlCfg{
		DSN:          "root:123456@tcp(localhost:3306)/zm4?charset=utf8&parseTime=True&loc=Local",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		Engine:       "InnoDB",
		Dbname:       "zm4",
	}
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.DSN, // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		models.Logger.Error("mysql connect error", zap.Error(err))
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
