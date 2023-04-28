package validation

import (
	"github.com/go-playground/validator/v10"
	idvalidator "github.com/guanguans/id-validator"
	"nft_platform/utils"
	"unicode"
)

func Mobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	return utils.ValidatorMobile(mobile)
}

func CardNo(fl validator.FieldLevel) bool {
	cardNo := fl.Field().String()
	return idvalidator.IsValid(cardNo, true)
}

func Chinese(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	for _, v := range str {
		if !unicode.Is(unicode.Han, v) {
			return false
		}
	}
	return true
}
