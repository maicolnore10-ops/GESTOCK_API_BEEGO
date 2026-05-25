package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type TipoMovimiento_20260521_131019 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TipoMovimiento_20260521_131019{}
	m.Created = "20260521_131019"

	migration.Register("TipoMovimiento_20260521_131019", m)
}

// Run the migrations
func (m *TipoMovimiento_20260521_131019) Up() {
m.SQL("INSERT INTO inventario.tipos_movimiento(nombre) VALUES ('Ingreso');") // to make schema update

}

// Reverse the migrations
func (m *TipoMovimiento_20260521_131019) Down() {
m.SQL("DELETE FROM inventario.tipos_movimiento WHERE nombre = 'Ingreso';") // to reverse schema update

}
