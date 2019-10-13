package cmd

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-boilerplate-project/app/types"
	"go-boilerplate-project/config"
	"strings"
)

func OpenDatabase(config *config.AppConfig) *gorm.DB {
	var dbUrl string
	if strings.EqualFold(config.Database.Dialect, "mysql") {
		databaseHost := fmt.Sprintf("%s:%d", config.Database.Host, config.Database.Port)
		dbUrl = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Database.User,
			config.Database.Password,
			databaseHost,
			config.Database.Name)
	} else {
		dbUrl = config.Database.Name
	}

	db, err := gorm.Open(config.Database.Dialect, dbUrl)
	if err != nil {
		panic(fmt.Errorf("Failed to open db: %s\n", err.Error()))
	}
	return db
}

func InitializeApplication() *types.Application {
	return &types.Application{
		Config: &config.Const,
	}
}
