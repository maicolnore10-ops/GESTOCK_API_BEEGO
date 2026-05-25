package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Ajustes struct {
	Id                int        `orm:"column(id_ajuste);pk;auto"`
	IdProducto        *Productos `orm:"column(id_producto);rel(fk)"` // Sin tag JSON. Go usará "IdProducto" por defecto.
	CantidadEsperada  int        `orm:"column(cantidad_esperada);null"`
	CantidadReal      int        `orm:"column(cantidad_real);null"`
	Diferencia        int        `orm:"column(diferencia);null"`
	Motivo            string     `orm:"column(motivo);null"`
	Activo            bool       `orm:"column(activo);null"`
	FechaCreacion     time.Time  `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaModificacion time.Time  `orm:"column(fecha_modificacion);type(timestamp without time zone);null;auto_now"`
}

func (t *Ajustes) TableName() string {
	return "ajustes"
}

func init() {
	orm.RegisterModel(new(Ajustes))
}

// AddAjustes insert a new Ajustes into database and returns
// last inserted Id on success.
func AddAjustes(m *Ajustes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAjustesById retrieves Ajustes by Id. Returns error if
// Id doesn't exist
func GetAjustesById(id int) (v *Ajustes, err error) {
	o := orm.NewOrm()
	v = &Ajustes{Id: id}
	if err = o.Read(v); err == nil {
		// CORRECCIÓN: Carga la estructura interna de la relación 'IdProducto' desde la BD
		if _, errRelated := o.LoadRelated(v, "IdProducto"); errRelated != nil {
			fmt.Println("Error al cargar la relación IdProducto:", errRelated.Error())
		}
		return v, nil
	}
	return nil, err
}

// GetAllAjustes retrieves all Ajustes matches certain condition. Returns empty list if
// no records exist
func GetAllAjustes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Ajustes))
	
	// CORRECCIÓN CRÍTICA: Hace el JOIN automático en la consulta masiva para evitar los campos vacíos
	qs = qs.RelatedSel()

	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
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
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
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

	var l []Ajustes
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
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

// UpdateAjustes updates Ajustes by Id and returns error if
// the record to be updated doesn't exist
func UpdateAjustesById(m *Ajustes) (err error) {
	o := orm.NewOrm()
	v := Ajustes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAjustes deletes Ajustes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAjustes(id int) (err error) {
	o := orm.NewOrm()
	v := Ajustes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Ajustes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}