package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hamgit.ir/novin-backend/trader-bot/config"
)

func Init(conf *config.Config) *gorm.DB {
	zap.L().Info("Establishing connection with core database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.C().Mysql.Username,
		config.C().Mysql.Password, config.C().Mysql.Host, config.C().Mysql.Port, config.C().Mysql.Schema)
	var dblogger logger.Interface

	if conf.App.Environment == config.EnvProduction {
		dblogger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dblogger})
	if err != nil {
		zap.L().Error("error while creating connection with database", zap.Error(err))
		return nil
	}

	zap.L().Info("Connection with core database established successfully")
	return db
}
