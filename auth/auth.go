package auth

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/gatewaygo/models"
)

//VerifyToken VerifyToken
func VerifyToken(db *gorm.DB, verificationToken string) error {

	if len(verificationToken) == 0 {
		return errors.New("Verification token not found in headers")
	}

	var redoxDestination models.RedoxDestination
	return db.Preload("TargetDestination").
		Where("verificationToken = ?", verificationToken).
		First(&redoxDestination).Error

}
