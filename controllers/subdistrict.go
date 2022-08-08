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

type SubdistrictController struct {
	beego.Controller
}

// GetAllSubdistricts function get all the subdistricts
func (sc *SubdistrictController) GetAllSubdistricts() {
	var subdistricts []*models.Subdistrict

	ff, _ := helpers.SetRequest(*sc.Ctx)
	subdistricts, err := models.GetAllSubdistricts(ff)

	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "Data failed to get",
			Data: types.ListSubdistrict{
				Subdistricts: subdistricts,
			},
		}
		sc.Data["json"] = fr
	} else {
		fr := types.FormattedResponse{
			StatusCode: 200,
			Message:    "Data obtained successfully",
			Data: types.ListSubdistrict{
				Subdistricts: subdistricts,
			},
		}
		sc.Data["json"] = fr
	}

	sc.ServeJSON()
}

// GetSubdistrictById gets a single subdistrict with the given id
func (sc *SubdistrictController) GetSubdistrictById() {
	// Get the id from query string
	id, _ := strconv.Atoi(sc.Ctx.Input.Param(":id"))

	// Get subdistrict
	subdistrict, err := models.GetSubdistrictById(int64(id))

	// Generate response
	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 404,
			Message:    "Data not found",
			Data: types.ShowSubdistrict{
				Subdistrict: nil,
			},
		}
		sc.Data["json"] = fr
	} else {
		fr := types.FormattedResponse{
			StatusCode: 200,
			Message:    "Data obtained successfully",
			Data: types.ShowSubdistrict{
				Subdistrict: subdistrict,
			},
		}
		sc.Data["json"] = fr
	}

	sc.ServeJSON()
}

// AddNewSubdistrict adds new subdistrict
func (sc *SubdistrictController) AddNewSubdistrict() {
	var p models.Subdistrict
	json.Unmarshal(sc.Ctx.Input.RequestBody, &p)

	// Validation
	vr, err := forms.CheckValidationSubdistrictAddForm(*sc.Ctx)
	if err == nil && len(vr) > 0 {
		fr := types.FormattedResponse{
			StatusCode: 422,
			Message:    "failed to add data",
			Data: types.ShowSubdistrict{
				Subdistrict: nil,
			},
			Errors: vr,
		}
		sc.Data["json"] = fr
		sc.ServeJSON()
	}

	// Create
	subdistrictId, err := models.InsertOneSubdistrict(&p)
	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "failed to add data",
			Data: types.ShowSubdistrict{
				Subdistrict: nil,
			},
		}
		sc.Data["json"] = fr
	} else {
		subdistrict, err := models.GetSubdistrictById(subdistrictId)

		if err != nil {
			fr := types.FormattedResponse{
				StatusCode: 404,
				Message:    "Failed to get inserted data",
				Data: types.ShowSubdistrict{
					Subdistrict: nil,
				},
			}
			sc.Data["json"] = fr
		} else {
			fr := types.FormattedResponse{
				StatusCode: 200,
				Message:    "Data created successfully",
				Data: types.ShowSubdistrict{
					Subdistrict: subdistrict,
				},
			}
			sc.Data["json"] = fr
		}
	}

	sc.ServeJSON()
}

// UpdateSubdistrict updates an existing subdistrict
func (sc *SubdistrictController) UpdateSubdistrict() {
	// Get the id from query string
	id, _ := strconv.Atoi(sc.Ctx.Input.Param(":id"))

	var p models.Subdistrict
	json.Unmarshal(sc.Ctx.Input.RequestBody, &p)

	// Validation
	vr, err := forms.CheckValidationSubdistrictUpdateForm(*sc.Ctx)
	if err == nil && len(vr) > 0 {
		fr := types.FormattedResponse{
			StatusCode: 422,
			Message:    "failed to update data",
			Data: types.ShowSubdistrict{
				Subdistrict: nil,
			},
			Errors: vr,
		}
		sc.Data["json"] = fr
		sc.ServeJSON()
	}

	// Update
	subdistrictId, err := models.UpdateSubdistrict(int64(id), &p)
	if err != nil {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "Failed to update data",
			Data: types.ShowSubdistrict{
				Subdistrict: nil,
			},
		}
		sc.Data["json"] = fr
	} else {
		subdistrict, err := models.GetSubdistrictById(subdistrictId)

		if err != nil {
			fr := types.FormattedResponse{
				StatusCode: 404,
				Message:    "Failed to get updated data",
				Data: types.ShowSubdistrict{
					Subdistrict: nil,
				},
			}
			sc.Data["json"] = fr
		} else {
			fr := types.FormattedResponse{
				StatusCode: 200,
				Message:    "Data updated successfully",
				Data: types.ShowSubdistrict{
					Subdistrict: subdistrict,
				},
			}
			sc.Data["json"] = fr
		}
	}

	sc.ServeJSON()
}

// DeleteSubdistrict deletes an existing subdistrict
func (sc *SubdistrictController) DeleteSubdistrict() {
	// Get id from query string and convert it to int
	id, _ := strconv.Atoi(sc.Ctx.Input.Param(":id"))

	// Delete
	deleted := models.DeleteSubdistrict(int64(id))
	if !deleted {
		fr := types.FormattedResponse{
			StatusCode: 500,
			Message:    "Failed to deleted data",
			Data: types.ShowSubdistrict{
				Subdistrict: nil,
			},
		}
		sc.Data["json"] = fr
	} else {
		fr := types.FormattedResponse{
			StatusCode: 404,
			Message:    "Data successfully deleted",
			Data:       nil,
		}
		sc.Data["json"] = fr
	}

	sc.ServeJSON()
}
