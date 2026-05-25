package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type UnidadesMedida struct {
	Id                 int       `orm:"column(id_unidad_medida);pk;auto"`
	Nombre             string    `orm:"column(nombre);null"`
	Sigla              string    `orm:"column(sigla);null"`
	Activo             bool      `orm:"column(activo)"`
	FechaCreacion      time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaActualizacion time.Time `orm:"column(fecha_actualizacion);type(timestamp without time zone);null;auto_now"`
}

func (t *UnidadesMedida) TableName() string {
	return "unidades_medida"
}

func init() {
	orm.RegisterModel(new(UnidadesMedida))
}

// AddUnidadesMedida inserta una nueva UnidadMedida en la base de datos y retorna
// el Id insertado en caso de éxito.
func AddUnidadesMedida(m *UnidadesMedida) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUnidadesMedidaById obtiene una UnidadMedida por Id. Retorna error si
// el Id no existe.
func GetUnidadesMedidaById(id int) (v *UnidadesMedida, err error) {
	o := orm.NewOrm()
	v = &UnidadesMedida{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUnidadesMedida obtiene todas las UnidadesMedida que cumplen cierta condición.
// Retorna lista vacía si no existen registros.
func GetAllUnidadesMedida(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UnidadesMedida))
	// query k=v
	for k, v := range query {
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

	var l []UnidadesMedida
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

// UpdateUnidadesMedidaById actualiza una UnidadMedida por Id. Retorna error si
// el registro a actualizar no existe.
func UpdateUnidadesMedidaById(m *UnidadesMedida) (err error) {
	o := orm.NewOrm()
	v := UnidadesMedida{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUnidadesMedida elimina una UnidadMedida por Id. Retorna error si
// el registro a eliminar no existe.
func DeleteUnidadesMedida(id int) (err error) {
	o := orm.NewOrm()
	v := UnidadesMedida{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UnidadesMedida{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
