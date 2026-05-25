CREATE SCHEMA IF NOT EXISTS empresa;



CREATE TABLE empresa.unidades_medida (
    id_unidad_medida SERIAL PRIMARY KEY,
    nombre VARCHAR(50),
    sigla VARCHAR(10),
    activo BOOLEAN NOT NULL DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE empresa.bodegas (
    id_bodegas SERIAL PRIMARY KEY,
    id_empresa INT,
    nombre VARCHAR(100),
    direccion TEXT,
    capacidad_maxima INT,
    activo BOOLEAN NOT NULL DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE empresa.moneda (
    id_moneda SERIAL PRIMARY KEY,
    tipo_moneda VARCHAR(50),
    activo BOOLEAN NOT NULL DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
    
    
CREATE TABLE empresa.empresas (
    id_empresa SERIAL PRIMARY KEY,
    nombre_empresa VARCHAR(100) NOT NULL,
    correo VARCHAR(100),
    activo BOOLEAN NOT NULL DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);