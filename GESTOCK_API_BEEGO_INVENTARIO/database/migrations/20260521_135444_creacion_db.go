package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreciacionDb_20260521_135444 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreciacionDb_20260521_135444{}
	m.Created = "20260521_135444"

	migration.Register("CreciacionDb_20260521_135444", m)
}

// Run the migrations
func (m *CreciacionDb_20260521_135444) Up() {
	m.SQL("INSERT INTO inventario.creciacion_db(nombre) VALUES ('Creación de base de datos');") // to make schema update

}

// Reverse the migrations
func (m *CreciacionDb_20260521_135444) Down() {
	m.SQL("DELETE FROM inventario.creciacion_db WHERE nombre = 'Creación de base de datos';") // to reverse schema update

}
