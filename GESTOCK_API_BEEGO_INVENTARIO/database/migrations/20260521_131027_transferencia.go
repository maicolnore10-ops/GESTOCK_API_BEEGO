package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Transferencia_20260521_131027 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Transferencia_20260521_131027{}
	m.Created = "20260521_131027"

	migration.Register("Transferencia_20260521_131027", m)
}

// Run the migrations
func (m *Transferencia_20260521_131027) Up() {
m.SQL("INSERT INTO inventario.transferencia(cantidad) VALUES (5);") // to make schema update

}

// Reverse the migrations
func (m *Transferencia_20260521_131027) Down() {
m.SQL("DELETE FROM inventario.transferencia WHERE cantidad = 5;") // to reverse schema update

}
