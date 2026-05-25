package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Moneda_20260525_133442 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Moneda_20260525_133442{}
	m.Created = "20260525_133442"

	migration.Register("Moneda_20260525_133442", m)
}

// Run the migrations
func (m *Moneda_20260525_133442) Up() {
	m.SQL("INSERT INTO empresa.moneda (tipo_moneda) VALUES ('Peso')")

}

// Reverse the migrations
func (m *Moneda_20260525_133442) Down() {
	m.SQL("DELETE FROM empresa.moneda WHERE tipo_moneda = 'Peso'")

}
