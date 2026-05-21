package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Productos_20260521_131011 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Productos_20260521_131011{}
	m.Created = "20260521_131011"

	migration.Register("Productos_20260521_131011", m)
}

// Run the migrations
func (m *Productos_20260521_131011) Up() {
	m.SQL("INSERT INTO inventario.productos(nombre, sku, precio_unitario, stock_inicial)VALUES ('Tecno Spark 10 pro', 'SKU001', 100.00, 10);") // to make schema update
	
}

// Reverse the migrations
func (m *Productos_20260521_131011) Down() {
	m.SQL("DELETE FROM inventario.productos WHERE nombre = 'Producto 1' AND sku = 'SKU001';") // to reverse schema update

}
