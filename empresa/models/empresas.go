package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Empresas struct {
	Id                 int       `orm:"column(id_empresa);pk;auto"`
	NombreEmpresa      string    `orm:"column(nombre_empresa)"`
	Correo             string    `orm:"column(correo);null"`
	Activo             bool      `orm:"column(activo)"`
	FechaCreacion      time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaActualizacion time.Time `orm:"column(fecha_actualizacion);type(timestamp without time zone);null;auto_now"`
}

func (t *Empresas) TableName() string {
	return "empresas"
}

func init() {
	orm.RegisterModel(new(Empresas))
}

// AddEmpresas inserta una nueva Empresa en la base de datos y retorna
// el Id insertado en caso de éxito.
func AddEmpresas(m *Empresas) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEmpresasById obtiene una Empresa por Id. Retorna error si
// el Id no existe.
func GetEmpresasById(id int) (v *Empresas, err error) {
	o := orm.NewOrm()
	v = &Empresas{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEmpresas obtiene todas las Empresas que cumplen cierta condición.
// Retorna lista vacía si no existen registros.
func GetAllEmpresas(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Empresas))
	// query k=v
	for k, v := range query {
		// reescribir notación de punto a Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// ordenamiento
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

	var l []Empresas
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// recortar campos no utilizados
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

// UpdateEmpresasById actualiza una Empresa por Id. Retorna error si
// el registro a actualizar no existe.
func UpdateEmpresasById(m *Empresas) (err error) {
	o := orm.NewOrm()
	v := Empresas{Id: m.Id}
	// verificar que el id existe en la base de datos
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEmpresas elimina una Empresa por Id. Retorna error si
// el registro a eliminar no existe.
func DeleteEmpresas(id int) (err error) {
	o := orm.NewOrm()
	v := Empresas{Id: id}
	// verificar que el id existe en la base de datos
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Empresas{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
