package forms

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web/context"
	"paninti-region-service/helpers"
)

type CityAddForm struct {
	Type       string `valid:"Required"`
	Name       string `valid:"Required"`
	ProvinceId int64  `valid:"Required"`
}

type CityUpdateForm struct {
	Type       string `valid:"Required"`
	Name       string `valid:"Required"`
	ProvinceId int64  `valid:"Required"`
}

func CheckValidationCityAddForm(ctx context.Context) (vr map[string]string, err error) {
	validationResult := make(map[string]string)
	form := CityAddForm{}

	json.Unmarshal(ctx.Input.RequestBody, &form)
	valid := validation.Validation{}

	b, err := valid.Valid(&form)
	if err != nil {
		return validationResult, err
	}

	if !b {
		validationResult = helpers.GetValidationResult(valid)
	}

	return validationResult, nil
}

func CheckValidationCityUpdateForm(ctx context.Context) (vr map[string]string, err error) {
	validationResult := make(map[string]string)
	form := CityUpdateForm{}

	json.Unmarshal(ctx.Input.RequestBody, &form)
	valid := validation.Validation{}

	b, err := valid.Valid(&form)
	if err != nil {
		return validationResult, err
	}

	if !b {
		validationResult = helpers.GetValidationResult(valid)
	}

	return validationResult, nil
}
