package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Bodegas_20260525_133412 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Bodegas_20260525_133412{}
	m.Created = "20260525_133412"

	migration.Register("Bodegas_20260525_133412", m)
}

// Run the migrations
func (m *Bodegas_20260525_133412) Up() {
m.SQL("INSERT INTO empresa.bodegas (nombre_bodega, ubicacion) VALUES ('Bodega Principal Norte', 'Zona Industrial')")
}
// Reverse the migrations
func (m *Bodegas_20260525_133412) Down() {
	m.SQL("DELETE FROM empresa.bodegas WHERE nombre_bodega = 'Bodega Principal Norte'")

}
