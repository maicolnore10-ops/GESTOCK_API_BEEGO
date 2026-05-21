package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Ajustes_20260521_130914 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Ajustes_20260521_130914{}
	m.Created = "20260521_130914"

	migration.Register("Ajustes_20260521_130914", m)
}

// Run the migrations
func (m *Ajustes_20260521_130914) Up() {
m.SQL("INSERT INTO inventario.ajustes(cantidad_esperada, cantidad_real, diferencia, motivo )VALUES ('39', '35', '4', 'Inventario inicial');") // to make schema update

}

// Reverse the migrations
func (m *Ajustes_20260521_130914) Down() {
m.SQL("DELETE FROM inventario.ajustes WHERE motivo = 'Inventario inicial';")  // to reverse schema update

}
