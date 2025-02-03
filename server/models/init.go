package models

import (
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB = nil
var dbInitMutex sync.Mutex

func initDb() error {
	dbInitMutex.Lock()
	defer dbInitMutex.Unlock()

	if dbConnection != nil {
		return nil
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&BlupStatistics{},
		&Cow{},
		&Genetic{},
		&GeneticIllnessStatus{},
		&GeneticIllnessData{},
		&GeneticIllness{},
		&User{},
		&Region{},
		&News{},
		&Partner{},
		&Breed{},
		&Farm{},
		&Role{},
		&Sex{},
		&Event{},
		&Grade{},
		&GradeHoz{},
		&GradeRegion{},
		&GradeCountry{},
		&EventType{},
		&Lactation{},
		&CheckMilk{},
		&DailyMilk{},
		&District{},
		&Exterior{},
		&Update{},
		&Document{},
		&UserRegisterRequest{},
		&HozRegisterRequest{},
		&HoldRegisterRequest{},
		&AdditionalInfo{},
		&DownSides{},
		&Measures{},
		&GaussianStatistics{},
	)
	dbConnection = db
	return nil
}

func GetDb() *gorm.DB {
	if dbConnection == nil {
		if err := initDb(); err != nil {
			panic(err)
		}
		// Get generic database object sql.DB to use its functions
		sqlDB, err := dbConnection.DB()
		if err != nil {
			panic(err)
		}

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(20)
		//sqlDB.SetMaxIdleConns(2)
	}
	return dbConnection
}
