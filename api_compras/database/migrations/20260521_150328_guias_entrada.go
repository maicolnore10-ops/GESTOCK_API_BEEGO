// Up: Run the migrations
func (m *GuiasEntrada_20260521_150328) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS compras.guias_entrada (
		id_guia SERIAL PRIMARY KEY,
		id_proveedor INT,
		id_bodega INT,
		estado VARCHAR(50),
		activo BOOLEAN NOT NULL DEFAULT TRUE,
		CONSTRAINT fk_guias_entrada_proveedor
			FOREIGN KEY (id_proveedor)
			REFERENCES compras.proveedores(id)
			ON DELETE RESTRICT
	);`)

	m.SQL(`INSERT INTO compras.guias_entrada (id_proveedor, id_bodega, estado, activo) 
		   VALUES (1, 1, 'Recibido', true);`)
}

func (m *GuiasEntrada_20260521_150328) Down() {
	m.SQL(`DROP TABLE IF EXISTS compras.guias_entrada CASCADE;`)
}