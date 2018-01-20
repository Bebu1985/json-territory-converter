package jsonTerritoryConverter

import (
	"github.com/Bebu1985/jsonTerritoryConverter/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateDatabase(Path string, Log bool) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", Path)

	if err != nil {
		return db, err
	}

	db.LogMode(Log)
	db.AutoMigrate(&models.Servant{}, &models.AreaAction{}, &models.Area{}, &models.Group{})
	if !db.HasTable(&models.Servant{}) {
		db.CreateTable(&models.Servant{})
	}
	if !db.HasTable(&models.AreaAction{}) {
		db.CreateTable(&models.AreaAction{})
	}
	if !db.HasTable(&models.Area{}) {
		db.CreateTable(&models.Area{})
	}
	if !db.HasTable(&models.Group{}) {
		db.CreateTable(&models.Group{})
	}

	return db, nil
}
