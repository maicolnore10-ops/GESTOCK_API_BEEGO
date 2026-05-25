// Run the migrations
func (m *ProveedorProductos_20260521_150319) Up() {
    // 1. Creamos la tabla proveedor_productos con su relación correcta
    m.SQL(`CREATE TABLE IF NOT EXISTS compras.proveedor_productos (
        id_proveedor_producto SERIAL PRIMARY KEY,
        id_proveedor INT,
        id_producto INT,
        precio NUMERIC(10,2),
        activo BOOLEAN NOT NULL DEFAULT TRUE,
        fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_productos_proveedor 
            FOREIGN KEY (id_proveedor) 
            REFERENCES compras.proveedores(id) 
            ON DELETE RESTRICT
    );`)

    // 2. Insertamos el producto de prueba (Quitamos la inserción rota de proveedores)
    m.SQL(`INSERT INTO compras.proveedor_productos (id_proveedor, id_producto, precio) 
           VALUES (1, 10, 250000) ON CONFLICT DO NOTHING;`)
}

// Reverse the migrations
func (m *ProveedorProductos_20260521_150319) Down() {
    // Borrado seguro de la tabla en caso de revertir la migración
    m.SQL(`DROP TABLE IF EXISTS compras.proveedor_productos CASCADE;`)
}