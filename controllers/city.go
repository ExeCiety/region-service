package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"paninti-region-service/forms"
	"paninti-region-service/helpers"
	"paninti-region-service/models"
	"paninti-region-service/types"
	"strconv"
)

type CityController struct {
	beego.Controller
}

// GetAllCities function get all the cities
func (cc *CityController) GetAllCities() {
	var cities []*models.City

	ff, _ := helpers.SetRequest(*cc.Ctx)
	cities, err := models.GetAllCities(ff)

	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "Data failed to get",
			Data: types.ListCity{
				Cities: cities,
			},
		}
		cc.Data["json"] = fr
	} else {
		fr := types.FormattedResponse{
			StatusCode: 200,
			Message:    "Data obtained successfully",
			Data:       cities,
		}
		cc.Data["json"] = fr
	}

	cc.ServeJSON()
}

// GetCityById gets a single city with the given id
func (cc *CityController) GetCityById() {
	// Get the id from query string
	id, _ := strconv.Atoi(cc.Ctx.Input.Param(":id"))

	// Get city
	city, err := models.GetCityById(int64(id))

	// Generate response
	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 404,
			Message:    "Data not found",
			Data: types.ShowCity{
				City: nil,
			},
		}
		cc.Data["json"] = fr
	} else {
		fr := types.FormattedResponse{
			StatusCode: 200,
			Message:    "Data obtained successfully",
			Data: types.ShowCity{
				City: city,
			},
		}
		cc.Data["json"] = fr
	}

	cc.ServeJSON()
}

// AddNewCity adds new city
func (cc *CityController) AddNewCity() {
	var p models.City
	json.Unmarshal(cc.Ctx.Input.RequestBody, &p)

	// Validation
	vr, err := forms.CheckValidationCityAddForm(*cc.Ctx)
	if err == nil && len(vr) > 0 {
		fr := types.FormattedResponse{
			StatusCode: 422,
			Message:    "failed to add data",
			Data: types.ShowCity{
				City: nil,
			},
			Errors: vr,
		}
		cc.Data["json"] = fr
		cc.ServeJSON()
	}

	// Insert
	cityId, err := models.InsertOneCity(&p)

	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "failed to add data",
			Data: types.ShowCity{
				City: nil,
			},
		}
		cc.Data["json"] = fr
	} else {
		city, err := models.GetCityById(cityId)

		if err != nil {
			fr := types.FormattedResponse{
				StatusCode: 404,
				Message:    "Failed to get inserted data",
				Data: types.ShowCity{
					City: nil,
				},
			}
			cc.Data["json"] = fr
		} else {
			fr := types.FormattedResponse{
				StatusCode: 200,
				Message:    "Data created successfully",
				Data: types.ShowCity{
					City: city,
				},
			}
			cc.Data["json"] = fr
		}
	}

	cc.ServeJSON()
}

// UpdateCity updates an existing city
func (cc *CityController) UpdateCity() {
	// Get the id from query string
	id, _ := strconv.Atoi(cc.Ctx.Input.Param(":id"))

	var p models.City
	json.Unmarshal(cc.Ctx.Input.RequestBody, &p)

	// Validation
	vr, err := forms.CheckValidationCityUpdateForm(*cc.Ctx)
	if err == nil && len(vr) > 0 {
		fr := types.FormattedResponse{
			StatusCode: 422,
			Message:    "failed to update data",
			Data: types.ShowCity{
				City: nil,
			},
			Errors: vr,
		}
		cc.Data["json"] = fr
		cc.ServeJSON()
	}

	// Update
	cityId, err := models.UpdateCity(int64(id), &p)

	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "Failed to update data",
			Data: types.ShowCity{
				City: nil,
			},
		}
		cc.Data["json"] = fr
	} else {
		city, err := models.GetCityById(cityId)

		if err != nil {
			fr := types.FormattedResponse{
				StatusCode: 404,
				Message:    "Failed to get updated data",
				Data: types.ShowCity{
					City: nil,
				},
			}
			cc.Data["json"] = fr
		} else {
			fr := types.FormattedResponse{
				StatusCode: 200,
				Message:    "Data updated successfully",
				Data: types.ShowCity{
					City: city,
				},
			}
			cc.Data["json"] = fr
		}
	}

	cc.ServeJSON()
}

// DeleteCity deletes an existing city
func (cc *CityController) DeleteCity() {
	// Get id from query string and convert it to int
	id, _ := strconv.Atoi(cc.Ctx.Input.Param(":id"))

	// Delete city
	deleted := models.DeleteCity(int64(id))
	if !deleted {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "Failed to deleted data",
			Data: types.ShowCity{
				City: nil,
			},
		}
		cc.Data["json"] = fr
	} else {
		fr := types.FormattedResponse{
			StatusCode: 404,
			Message:    "Data successfully deleted",
			Data:       nil,
		}
		cc.Data["json"] = fr
	}

	cc.ServeJSON()
}
