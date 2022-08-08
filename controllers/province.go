package controllers

import (
	"encoding/json"
	"paninti-region-service/forms"
	"paninti-region-service/helpers"
	"paninti-region-service/types"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
	"paninti-region-service/models"
)

type ProvinceController struct {
	beego.Controller
}

// GetAllProvinces function get all the provinces
func (pc *ProvinceController) GetAllProvinces() {
	var provinces []*models.Province

	ff, _ := helpers.SetRequest(*pc.Ctx)
	provinces, err := models.GetAllProvinces(ff)

	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "Data failed to get",
			Data: types.ListProvince{
				Provinces: provinces,
			},
		}
		pc.Data["json"] = fr
	} else {
		fr := types.FormattedResponse{
			StatusCode: 200,
			Message:    "Data obtained successfully",
			Data: types.ListProvince{
				Provinces: provinces,
			},
		}
		pc.Data["json"] = fr
	}

	pc.ServeJSON()
}

// GetProvinceById gets a single province with the given id
func (pc *ProvinceController) GetProvinceById() {
	// Get the id from query string
	id, _ := strconv.Atoi(pc.Ctx.Input.Param(":id"))

	// Get province
	province, err := models.GetProvinceById(int64(id))

	// Generate response
	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 404,
			Message:    "Data not found",
			Data: types.ShowProvince{
				Province: nil,
			},
		}
		pc.Data["json"] = fr
	} else {
		fr := types.FormattedResponse{
			StatusCode: 200,
			Message:    "Data obtained successfully",
			Data: types.ShowProvince{
				Province: province,
			},
		}
		pc.Data["json"] = fr
	}

	pc.ServeJSON()
}

// AddNewProvince adds new province
func (pc *ProvinceController) AddNewProvince() {
	var p models.Province
	json.Unmarshal(pc.Ctx.Input.RequestBody, &p)
	provinceId, err := models.InsertOneProvince(&p)

	// Validation
	vr, err := forms.CheckValidationProvinceAddForm(*pc.Ctx)
	if err == nil && len(vr) > 0 {
		fr := types.FormattedResponse{
			StatusCode: 422,
			Message:    "failed to add data",
			Data: types.ShowCity{
				City: nil,
			},
			Errors: vr,
		}
		pc.Data["json"] = fr
		pc.ServeJSON()
	}

	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "failed to add data",
			Data: types.ShowProvince{
				Province: nil,
			},
		}
		pc.Data["json"] = fr
	} else {
		province, err := models.GetProvinceById(provinceId)

		if err != nil {
			fr := types.FormattedResponse{
				StatusCode: 404,
				Message:    "Failed to get inserted data",
				Data: types.ShowProvince{
					Province: nil,
				},
			}
			pc.Data["json"] = fr
		} else {
			fr := types.FormattedResponse{
				StatusCode: 200,
				Message:    "Data created successfully",
				Data: types.ShowProvince{
					Province: province,
				},
			}
			pc.Data["json"] = fr
		}
	}

	pc.ServeJSON()
}

// UpdateProvince updates an existing province
func (pc *ProvinceController) UpdateProvince() {
	// Get the id from query string
	id, _ := strconv.Atoi(pc.Ctx.Input.Param(":id"))
	var p models.Province
	json.Unmarshal(pc.Ctx.Input.RequestBody, &p)

	// Validation
	vr, err := forms.CheckValidationProvinceUpdateForm(*pc.Ctx)
	if err == nil && len(vr) > 0 {
		fr := types.FormattedResponse{
			StatusCode: 422,
			Message:    "failed to add data",
			Data: types.ShowCity{
				City: nil,
			},
			Errors: vr,
		}
		pc.Data["json"] = fr
		pc.ServeJSON()
	}

	// Update
	provinceId, err := models.UpdateProvince(int64(id), &p)
	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "Failed to update data",
			Data: types.ShowProvince{
				Province: nil,
			},
		}
		pc.Data["json"] = fr
	} else {
		province, err := models.GetProvinceById(provinceId)

		if err != nil {
			fr := types.FormattedResponse{
				StatusCode: 404,
				Message:    "Failed to get updated data",
				Data: types.ShowProvince{
					Province: nil,
				},
			}
			pc.Data["json"] = fr
		} else {
			fr := types.FormattedResponse{
				StatusCode: 200,
				Message:    "Data updated successfully",
				Data: types.ShowProvince{
					Province: province,
				},
			}
			pc.Data["json"] = fr
		}
	}

	pc.ServeJSON()
}

// DeleteProvince deletes an existing province
func (pc *ProvinceController) DeleteProvince() {
	// Get id from query string and convert it to int
	id, _ := strconv.Atoi(pc.Ctx.Input.Param(":id"))

	// Delete province
	deleted := models.DeleteProvince(int64(id))
	if !deleted {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "Failed to deleted data",
			Data: types.ShowProvince{
				Province: nil,
			},
		}
		pc.Data["json"] = fr
	} else {
		fr := types.FormattedResponse{
			StatusCode: 404,
			Message:    "Data successfully deleted",
			Data:       nil,
		}
		pc.Data["json"] = fr
	}

	pc.ServeJSON()
}
