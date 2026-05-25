package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type GuiasEntrada_20260521_150328 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &GuiasEntrada_20260521_150328{}
	m.Created = "20260521_150328"

	migration.Register("GuiasEntrada_20260521_150328", m)
}

// Run the migrations
func (m *GuiasEntrada_20260521_150328) Up() {
m.SQL(`INSERT INTO compras.guias_entrada ( id_proveedor, id_bodega, estado, activo)
VALUES ('2', 1,'Recibido', true);`)
}

// Reverse the migrations
func (m *GuiasEntrada_20260521_150328) Down() {
	m.SQL(`DELETE FROM compras.guias_entrada WHERE estado = 'Recibido';`)

}
