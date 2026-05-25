package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Inventario_20260521_130955 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Inventario_20260521_130955{}
	m.Created = "20260521_130955"

	migration.Register("Inventario_20260521_130955", m)
}

// Run the migrations
func (m *Inventario_20260521_130955) Up() {
	m.SQL("INSERT INTO inventario.inventario(nombre, fecha)VALUES ('Inventario 1', '2027-01-01');") // to make schema update

}

// Reverse the migrations
func (m *Inventario_20260521_130955) Down() {
	m.SQL("DELETE FROM inventario.inventario WHERE nombre = 'Inventario 1' AND fecha = '2027-01-01';") // to reverse schema update

}
