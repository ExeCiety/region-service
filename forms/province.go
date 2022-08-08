package forms

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web/context"
	"paninti-region-service/helpers"
)

type ProvinceAddForm struct {
	Name string `valid:"Required"`
}

type ProvinceUpdateForm struct {
	Name string `valid:"Required"`
}

func CheckValidationProvinceAddForm(ctx context.Context) (vr map[string]string, err error) {
	validationResult := make(map[string]string)
	form := ProvinceAddForm{}

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

func CheckValidationProvinceUpdateForm(ctx context.Context) (vr map[string]string, err error) {
	validationResult := make(map[string]string)
	form := ProvinceUpdateForm{}

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
