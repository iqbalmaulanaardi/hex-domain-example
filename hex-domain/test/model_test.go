package test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/listing"
	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/models"
	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/storage/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var server = database.Storage{}
var userInstance = models.User{}

var lister *listing.Service

func TestMain(m *testing.M) {
	s, _ := database.NewStorage("hex-example-test")
	list := listing.NewService(s)
	lister = &list
	Database()
	os.Exit(m.Run())
}
func Database() (*database.Storage, error) {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "password", "127.0.0.1", "3306", "hex-example-test")
	server.DB, err = gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", "mysql")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", "hex-example-test")
	}
	return &server, nil
}
func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}
func seedOneUser() (models.User, error) {
	now := time.Now()
	createdAt := now.Unix()
	refreshUserTable()

	user := models.User{
		ID:        1,
		Nickname:  "foobar",
		Email:     "iqbalmaulana.ardi@gmail.com",
		Password:  "foobar",
		CreatedAt: createdAt,
		UpdatedAt: 0,
	}
	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	log.Printf("Successfully seed table")
	return user, nil
}
