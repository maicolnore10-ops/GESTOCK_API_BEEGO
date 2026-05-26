package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

// Sistema model
	type Sistema struct {
	Id            int64  `orm:"column(id_sistema);pk;auto" json:"id"`
    NombreEmpresa string `orm:"column(nombre_empresa)" json:"NombreEmpresa"`
    Correo        string `orm:"column(correo)" json:"Correo"`
    Descripcion   string `orm:"column(descripcion)" json:"Descripcion"`
    Activo        bool   `orm:"column(activo)" json:"Activo"`
}

func (t *Sistema) TableName() string {
	return "sistema"
}

func init() {
	orm.RegisterModel(new(Sistema))
}

// AddSistema insert a new Sistema into database and returns last inserted Id on success.
func AddSistema(m *Sistema) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSistemaById retrieves Sistema by Id.
func GetSistemaById(id int) (v *Sistema, err error) {
	o := orm.NewOrm()
	// Conversión explícita de int a int64 aplicada aquí
	v = &Sistema{Id: int64(id)}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSistema retrieves all Sistema.
func GetAllSistema(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sistema))
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
		}
	}

	var l []Sistema
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

// UpdateSistemaById updates Sistema by Id
func UpdateSistemaById(m *Sistema) (err error) {
	o := orm.NewOrm()
	v := Sistema{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSistema deletes Sistema by Id
func DeleteSistema(id int) (err error) {
	o := orm.NewOrm()
	// Conversiones explícitas aplicadas aquí
	v := Sistema{Id: int64(id)}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Sistema{Id: int64(id)}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}