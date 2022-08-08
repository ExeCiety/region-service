// Package routers
// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"paninti-region-service/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/provinces",
			// get all provinces
			beego.NSRouter("/", &controllers.ProvinceController{}, "get:GetAllProvinces"),

			// get a province with id
			beego.NSRouter("/:id", &controllers.ProvinceController{}, "get:GetProvinceById"),

			// add new province
			beego.NSRouter("/", &controllers.ProvinceController{}, "post:AddNewProvince"),

			// update an existing province
			beego.NSRouter("/:id", &controllers.ProvinceController{}, "put:UpdateProvince"),

			// delete a province
			beego.NSRouter("/:id", &controllers.ProvinceController{}, "delete:DeleteProvince"),
		),
		beego.NSNamespace("/cities",
			// get all provinces
			beego.NSRouter("/", &controllers.CityController{}, "get:GetAllCities"),

			// get a province with id
			beego.NSRouter("/:id", &controllers.CityController{}, "get:GetCityById"),

			// add new province
			beego.NSRouter("/", &controllers.CityController{}, "post:AddNewCity"),

			// update an existing province
			beego.NSRouter("/:id", &controllers.CityController{}, "put:UpdateCity"),

			// delete a province
			beego.NSRouter("/:id", &controllers.CityController{}, "delete:DeleteCity"),
		),
		beego.NSNamespace("/subdistricts",
			// get all subdistricts
			beego.NSRouter("/", &controllers.SubdistrictController{}, "get:GetAllSubdistricts"),

			// get a subdistrict with id
			beego.NSRouter("/:id", &controllers.SubdistrictController{}, "get:GetSubdistrictById"),

			// add new subdistrict
			beego.NSRouter("/", &controllers.SubdistrictController{}, "post:AddNewSubdistrict"),

			// update an existing subdistrict
			beego.NSRouter("/:id", &controllers.SubdistrictController{}, "put:UpdateSubdistrict"),

			// delete a subdistrict
			beego.NSRouter("/:id", &controllers.SubdistrictController{}, "delete:DeleteSubdistrict"),
		),
	)

	beego.AddNamespace(ns)
}
