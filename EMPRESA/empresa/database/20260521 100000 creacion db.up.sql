-- Creacion de Schema
CREATE SCHEMA IF NOT EXISTS empresa;

-- Creacion de Tabla Empresas
CREATE TABLE empresa.empresas (
    id_empresa          SERIAL          PRIMARY KEY,
    nombre_empresa      VARCHAR(100)    NOT NULL,
    correo              VARCHAR(100)    NULL,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);

-- Creacion de Tabla Unidades de Medida
CREATE TABLE empresa.unidades_medida (
    id_unidad_medida    SERIAL          PRIMARY KEY,
    nombre              VARCHAR(50)     NULL,
    sigla               VARCHAR(10)     NULL,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);

-- Creacion de Tabla Bodegas
CREATE TABLE empresa.bodegas (
    id_bodegas          SERIAL          PRIMARY KEY,
    id_empresa          INTEGER         NULL,
    nombre              VARCHAR(100)    NULL,
    direccion           TEXT            NULL,
    capacidad_maxima    INTEGER         NULL,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_bodegas_empresa
        FOREIGN KEY (id_empresa)
        REFERENCES empresa.empresas(id_empresa)
        ON DELETE RESTRICT
);

-- Creacion de Tabla Moneda
CREATE TABLE empresa.moneda (
    id_moneda           SERIAL          PRIMARY KEY,
    tipo_moneda         VARCHAR(50)     NULL,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);