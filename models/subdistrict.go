package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"paninti-region-service/helpers"
	"paninti-region-service/types"
	"time"
)

type Subdistrict struct {
	Id        int64     `orm:"column(id);pk;auto"`
	Name      string    `orm:"column(name)"`
	CityId    int64     `orm:"column(city_id)"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp without time zone);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp without time zone);null"`
}

func (p *Subdistrict) TableName() string {
	return "region_subdistricts"
}

func init() {
	orm.RegisterModel(new(Subdistrict))
}

// GetAllSubdistricts function get all subdistricts
func GetAllSubdistricts(ff types.FilterFormat) (p []*Subdistrict, err error) {
	o := orm.NewOrm()
	var subdistricts []*Subdistrict

	qs := o.QueryTable(new(Subdistrict))
	qs, err = helpers.FilterRequest(qs, ff)
	if err != nil {
		return subdistricts, err
	}

	_, err = qs.All(&subdistricts)

	return subdistricts, nil
}

// GetSubdistrictById get a subdistrict with the given id
func GetSubdistrictById(id int64) (p *Subdistrict, err error) {
	o := orm.NewOrm()
	p = &Subdistrict{Id: id}
	if err = o.Read(p); err == nil {
		return p, nil
	}
	return p, err
}

// InsertOneSubdistrict inserts a single new subdistrict record
func InsertOneSubdistrict(subdistrict *Subdistrict) (id int64, err error) {
	o := orm.NewOrm()

	// Insert
	id, err = o.Insert(&Subdistrict{
		Id:        subdistrict.Id,
		Name:      subdistrict.Name,
		CityId:    subdistrict.CityId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	fmt.Println(err)

	return id, err
}

// UpdateSubdistrict updates an existing subdistrict
func UpdateSubdistrict(id int64, subdistrict *Subdistrict) (rid int64, err error) {
	o := orm.NewOrm()
	p := &Subdistrict{Id: id}

	// Get existing subdistrict
	if err = o.Read(p); err == nil {
		// Update
		p = &Subdistrict{
			Id:        id,
			Name:      subdistrict.Name,
			CityId:    subdistrict.CityId,
			UpdatedAt: time.Now(),
		}
		_, err := o.Update(p, "name", "city_id", "updated_at")

		return id, err
	}

	return id, err
}

// DeleteSubdistrict delete a subdistrict
func DeleteSubdistrict(id int64) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Subdistrict{Id: id})
	if err == nil {
		// Successfully
		return true
	}

	return false
}
