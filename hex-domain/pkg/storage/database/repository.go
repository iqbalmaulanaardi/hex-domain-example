package database

import (
	"fmt"
	"log"

	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Storage struct {
	DB *gorm.DB
}
type TransactionRepo interface {
	GetUsersRepo() ([]models.User, error)
}

func NewStorage(dbname string) (*Storage, error) {
	var err error
	s := new(Storage)
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "password", "127.0.0.1", "3306", dbname)

	s.DB, err = gorm.Open("mysql", DBURL)
	if err != nil {
		log.Fatalf("cannot create connection: %v", err)
	}
	fmt.Println("Connected to db")
	return s, nil
}
