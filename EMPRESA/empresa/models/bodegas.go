package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Bodegas struct {
	Id                 int       `orm:"column(id_bodegas);pk;auto"`
	IdEmpresa          *Empresas `orm:"column(id_empresa);rel(fk);null"`
	Nombre             string    `orm:"column(nombre);null"`
	Direccion          string    `orm:"column(direccion);null"`
	CapacidadMaxima    int       `orm:"column(capacidad_maxima);null"`
	Activo             bool      `orm:"column(activo)"`
	FechaCreacion      time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaActualizacion time.Time `orm:"column(fecha_actualizacion);type(timestamp without time zone);null;auto_now"`
}

func (t *Bodegas) TableName() string {
	return "bodegas"
}

func init() {
	orm.RegisterModel(new(Bodegas))
}

// AddBodegas insert a new Bodegas into database and returns
// last inserted Id on success.
func AddBodegas(m *Bodegas) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBodegasById retrieves Bodegas by Id. Returns error if
// Id doesn't exist
func GetBodegasById(id int) (v *Bodegas, err error) {
	o := orm.NewOrm()
	v = &Bodegas{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "IdEmpresa")
		return v, nil
	}
	return nil, err
}

// GetAllBodegas retrieves all Bodegas matches certain condition. Returns empty list if
// no records exist
func GetAllBodegas(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Bodegas))
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Bodegas
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateBodegasById updates Bodegas by Id and returns error if
// the record to be updated doesn't exist
func UpdateBodegasById(m *Bodegas) (err error) {
	o := orm.NewOrm()
	v := Bodegas{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBodegas deletes Bodegas by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBodegas(id int) (err error) {
	o := orm.NewOrm()
	v := Bodegas{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Bodegas{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}