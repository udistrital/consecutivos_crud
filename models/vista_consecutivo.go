package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type VistaConsecutivo struct {
	Id                int     `orm:"column(id);pk;auto"`
	ContextoId        int     `orm:"column(contexto_id)"`
	Year              float64 `orm:"column(year)"`
	Consecutivo       int     `orm:"column(consecutivo)"`
	Descripcion       string  `orm:"column(descripcion);null"`
	Activo            bool    `orm:"column(activo)"`
	FechaCreacion     string  `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string  `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *VistaConsecutivo) ViewName() string {
	return "vista_consecutivo"
}

func init() {
	orm.RegisterModel(new(VistaConsecutivo))
}

// GetVistaConsecutivoById retrieves VistaConsecutivo by Id. Returns error if
// Id doesn't exist
func GetVistaConsecutivoById(id int) (v *VistaConsecutivo, err error) {
	o := orm.NewOrm()
	v = &VistaConsecutivo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllVistaConsecutivo retrieves all VistaConsecutivo matches certain condition. Returns empty list if
// no records exist
func GetAllVistaConsecutivo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(VistaConsecutivo))
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

	var l []VistaConsecutivo
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
