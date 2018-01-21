package jsonTerritoryConverter

import (
	"github.com/Bebu1985/jsonTerritoryConverter/models"
	"github.com/jinzhu/gorm"
	//Blank import is needed to support sqlite in gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//CreateDatabase migrates or creates a database for the given models
func CreateDatabase(Path string, Log bool) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", Path)

	if err != nil {
		return db, err
	}

	db.LogMode(Log)
	db.AutoMigrate(&models.Servant{}, &models.AreaAction{}, &models.Area{}, &models.Group{}, &models.GroupServantJoin{})

	createOrCleanFor(db, &models.Servant{})
	createOrCleanFor(db, &models.Area{})
	createOrCleanFor(db, &models.AreaAction{})
	createOrCleanFor(db, &models.Group{})
	createOrCleanFor(db, &models.GroupServantJoin{})

	return db, nil
}

func createOrCleanFor(db *gorm.DB, model interface{}) {
	if !db.HasTable(model) {
		db.CreateTable(model)
	} else {
		db.DropTable(model)
		db.CreateTable(model)
	}
}
