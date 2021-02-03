package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquemaConsecutivos_20210203_084841 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquemaConsecutivos_20210203_084841{}
	m.Created = "20210203_084841"

	migration.Register("CrearEsquemaConsecutivos_20210203_084841", m)
}

// Run the migrations
func (m *CrearEsquemaConsecutivos_20210203_084841) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210203_084841_crear_esquema_consecutivos_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *CrearEsquemaConsecutivos_20210203_084841) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210203_084841_crear_esquema_consecutivos_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
