package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Movimientos_20260521_131003 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Movimientos_20260521_131003{}
	m.Created = "20260521_131003"

	migration.Register("Movimientos_20260521_131003", m)
}

// Run the migrations
func (m *Movimientos_20260521_131003) Up() {
m.SQL("INSERT INTO inventario.movimientos(cantidad, motivo, fecha) VALUES (5, 'Ingreso de mercancía', '2027-01-01');") // to make schema update

}

// Reverse the migrations
func (m *Movimientos_20260521_131003) Down() {
m.SQL("DELETE FROM inventario.movimientos WHERE cantidad = 5 AND motivo = 'Ingreso de mercancía' AND fecha = '2027-01-01';") // to reverse schema update

}
