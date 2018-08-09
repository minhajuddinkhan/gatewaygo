package auth

import (
	"errors"

	"github.com/jinzhu/gorm"
	models "github.com/minhajuddinkhan/gatewaygo/models"
)

//VerifyToken VerifyToken
func VerifyToken(db *gorm.DB, verificationToken string) error {

	if len(verificationToken) == 0 {
		return errors.New("Verification token not found in headers")
	}

	rd := models.RedoxDestination{
		VerificationToken: verificationToken,
	}

	err := db.Preload("TargetDestination").First(&rd).Error
	if err != nil {
		return err
	}

	return nil

}
