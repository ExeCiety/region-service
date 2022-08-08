package models

import (
	"github.com/beego/beego/v2/client/orm"
	"paninti-region-service/helpers"
	"paninti-region-service/types"
	"time"
)

type Province struct {
	Id        int64     `orm:"column(id);pk;auto"`
	Name      string    `orm:"column(name)"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp without time zone);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp without time zone);null"`
}

func (p *Province) TableName() string {
	return "region_provinces"
}

func init() {
	orm.RegisterModel(new(Province))
}

// GetAllProvinces function get all provinces
func GetAllProvinces(ff types.FilterFormat) (p []*Province, err error) {
	o := orm.NewOrm()
	var provinces []*Province

	qs := o.QueryTable(new(Province))
	qs, err = helpers.FilterRequest(qs, ff)
	if err != nil {
		return provinces, err
	}

	_, err = qs.All(&provinces)

	return provinces, nil
}

// GetProvinceById get a province with the given id
func GetProvinceById(id int64) (p *Province, err error) {
	o := orm.NewOrm()
	p = &Province{Id: id}
	if err = o.Read(p); err == nil {
		return p, nil
	}
	return p, err
}

// InsertOneProvince inserts a single new province record
func InsertOneProvince(province *Province) (id int64, err error) {
	o := orm.NewOrm()

	// Insert
	id, err = o.Insert(&Province{Id: province.Id, Name: province.Name, CreatedAt: time.Now(), UpdatedAt: time.Now()})

	return id, err
}

// UpdateProvince updates an existing province
func UpdateProvince(id int64, province *Province) (rid int64, err error) {
	o := orm.NewOrm()
	p := &Province{Id: id}

	// Get existing province
	if err = o.Read(p); err == nil {
		// Update
		p = &Province{Id: id, Name: province.Name, UpdatedAt: time.Now()}
		_, err := o.Update(p, "name", "updated_at")

		return id, err
	}

	return id, err
}

// DeleteProvince delete a province
func DeleteProvince(id int64) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Province{Id: id})
	if err == nil {
		// Successfully
		return true
	}

	return false
}
