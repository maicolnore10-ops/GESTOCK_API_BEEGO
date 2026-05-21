package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InventarioProductos_20260521_130947 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InventarioProductos_20260521_130947{}
	m.Created = "20260521_130947"

	migration.Register("InventarioProductos_20260521_130947", m)
}

// Run the migrations
func (m *InventarioProductos_20260521_130947) Up() {
m.SQL("INSERT INTO inventario.inventario_productos(cantidad) VALUES (6);") // to reverse schema update


}

// Reverse the migrations
func (m *InventarioProductos_20260521_130947) Down() {
m.SQL("DELETE FROM inventario.inventario_productos WHERE cantidad = 6;") // to reverse schema update

}
