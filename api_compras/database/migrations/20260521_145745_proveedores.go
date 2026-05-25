// Up: Run the migrations
func (m *ProveedorProductos_20260521_150319) Up() {
    // Creamos la tabla proveedor_productos con la relación limpia hacia proveedores
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
}

// Down: Reverse the migrations
func (m *ProveedorProductos_20260521_150319) Down() {
    m.SQL(`DROP TABLE IF EXISTS compras.proveedor_productos CASCADE;`)
}