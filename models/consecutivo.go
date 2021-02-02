package models

import (
	"github.com/astaxie/beego/orm"
)

type Consecutivo struct {
	Id                int     `orm:"column(id);pk;auto"`
	ContextoId        int     `orm:"column(contexto_id)"`
	Year              float64 `orm:"column(year)"`
	Descripcion       string  `orm:"column(descripcion);null"`
	Activo            bool    `orm:"column(activo)"`
	FechaCreacion     string  `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string  `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *Consecutivo) TableName() string {
	return "consecutivo"
}

func init() {
	orm.RegisterModel(new(Consecutivo))
}

// AddConsecutivo insert a new Consecutivo into database and returns
// last inserted Id on success.
func AddConsecutivo(m *Consecutivo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
