package store

import (
	"github.com/jinzhu/gorm"
)

//Store Store
type Store struct {
	Database *gorm.DB
}

//NewPostgresStore NewPostgresStore
func NewPostgresStore(connString string) (*Store, error) {
	conn, err := gorm.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &Store{Database: conn}, nil
}
