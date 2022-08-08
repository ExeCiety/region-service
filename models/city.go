package models

import (
	"github.com/beego/beego/v2/client/orm"
	"paninti-region-service/helpers"
	"paninti-region-service/types"
	"time"
)

type City struct {
	Id         int64     `orm:"column(id);pk;auto"`
	Type       string    `orm:"column(type)";valid:"Required"`
	Name       string    `orm:"column(name)"`
	ProvinceId int64     `orm:"column(province_id)"`
	CreatedAt  time.Time `orm:"column(created_at);type(timestamp without time zone);null"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(timestamp without time zone);null"`
}

func (c *City) TableName() string {
	return "region_cities"
}

func init() {
	orm.RegisterModel(new(City))
}

// GetAllCities function get all cities
func GetAllCities(ff types.FilterFormat) (c []*City, err error) {
	o := orm.NewOrm()
	var cities []*City

	qs := o.QueryTable(new(City))
	qs, err = helpers.FilterRequest(qs, ff)
	if err != nil {
		return cities, err
	}

	_, err = qs.All(&cities)

	return cities, err
}

// GetCityById get a city with the given id
func GetCityById(id int64) (p *City, err error) {
	o := orm.NewOrm()
	p = &City{Id: id}
	if err = o.Read(p); err == nil {
		return p, nil
	}
	return p, err
}

// InsertOneCity inserts a single new city record
func InsertOneCity(city *City) (id int64, err error) {
	o := orm.NewOrm()

	// Insert
	id, err = o.Insert(&City{
		Id:         city.Id,
		Type:       city.Type,
		Name:       city.Name,
		ProvinceId: city.ProvinceId,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	return id, err
}

// UpdateCity updates an existing city
func UpdateCity(id int64, city *City) (rid int64, err error) {
	o := orm.NewOrm()
	p := &City{Id: id}

	// Get existing city
	if err = o.Read(p); err == nil {
		// Update
		p = &City{
			Id:         id,
			Type:       city.Type,
			Name:       city.Name,
			ProvinceId: city.ProvinceId,
			UpdatedAt:  time.Now(),
		}
		_, err := o.Update(p, "type", "name", "province_id", "updated_at")

		return id, err
	}

	return id, err
}

// DeleteCity delete a city
func DeleteCity(id int64) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&City{Id: id})
	if err == nil {
		// Successfully
		return true
	}

	return false
}
