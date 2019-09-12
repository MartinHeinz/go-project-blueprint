package test_data

import (
	"fmt"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"strings"
)

func init() {
	// the test may be started from the home directory or a subdirectory
	err := config.LoadConfig("/config") // on host use absolute path
	if err != nil {
		panic(err)
	}
	config.Config.DB, config.Config.DBErr = gorm.Open("sqlite3", ":memory:")
	config.Config.DB.Exec("PRAGMA foreign_keys = ON") // SQLite defaults to `foreign_keys = off'`
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	config.Config.DB.AutoMigrate(&models.User{})
}

func ResetDB() *gorm.DB {
	config.Config.DB.DropTableIfExists(&models.User{}) // Note: Order matters
	config.Config.DB.AutoMigrate(&models.User{})
	if err := runSQLFile(config.Config.DB, getSQLFile()); err != nil {
		panic(fmt.Errorf("error while initializing test database: %s", err))
	}
	return config.Config.DB
}

func getSQLFile() string {
	return "/test_data/db.sql" // on host use absolute path
}

func GetTestCaseFolder() string {
	return "/test_data/test_case_data" // on host use absolute path
}

func runSQLFile(db *gorm.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if result := db.Exec(line); result.Error != nil {
			fmt.Println(line)
			return result.Error
		}
	}
	return nil
}
