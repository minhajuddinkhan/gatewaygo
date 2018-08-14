package store

import (
	"github.com/jinzhu/gorm"
)

//GetToken GetToken
func (s *Store) GetToken(verificationToken string) *gorm.DB {
	return s.Database.Preload("TargetDestination").Where(`"verificationToken" = ?`, verificationToken)

}
