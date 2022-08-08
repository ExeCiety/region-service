package forms

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web/context"
	"paninti-region-service/helpers"
)

type SubdistrictAddForm struct {
	Name   string `valid:"Required"`
	CityId int64  `valid:"Required"`
}

type SubdistrictUpdateForm struct {
	Name   string `valid:"Required"`
	CityId int64  `valid:"Required"`
}

func CheckValidationSubdistrictAddForm(ctx context.Context) (vr map[string]string, err error) {
	validationResult := make(map[string]string)
	form := SubdistrictAddForm{}

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

func CheckValidationSubdistrictUpdateForm(ctx context.Context) (vr map[string]string, err error) {
	validationResult := make(map[string]string)
	form := SubdistrictUpdateForm{}

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
