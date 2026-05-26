# API Configuración

API REST para gestionar la configuración del sistema, seguridad, backups y notificaciones. Desarrollada con Beego y PostgreSQL.

## 📋 Tabla de Contenidos

- [Características](#características)
- [Requisitos Previos](#requisitos-previos)
- [Instalación](#instalación)
- [Configuración](#configuración)
- [Ejecución](#ejecución)
- [Endpoints](#endpoints)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Tecnologías Utilizadas](#tecnologías-utilizadas)

## ✨ Características

- ✅ Operaciones CRUD completas para Configuración del Sistema
- ✅ Operaciones CRUD completas para Seguridad
- ✅ Operaciones CRUD completas para Backups
- ✅ Operaciones CRUD completas para Notificaciones
- ✅ Validación de datos de entrada
- ✅ Manejo de errores con respuestas HTTP adecuadas
- ✅ Soporte para CORS
- ✅ Paginación y filtrado de resultados
- ✅ Variables de entorno para configuración segura

## 📦 Requisitos Previos

Asegúrate de tener instalados:

- **Go** 1.24.2 o superior
- **PostgreSQL** 12 o superior
- **Git** para control de versiones

## 🚀 Instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/tu-usuario/api_configuracion.git
cd api_configuracion
```

### 2. Descargar dependencias

```bash
go mod download
go mod tidy
```

### 3. Verificar la estructura de carpetas

```
api_configuracion/
├── conf/
│   └── app.conf
├── controllers/
│   ├── sistema.go
│   ├── seguridad.go
│   ├── backup.go
│   └── notificacion.go
├── models/
│   ├── sistema.go
│   ├── seguridad.go
│   ├── backup.go
│   └── notificacion.go
├── routers/
│   └── router.go
├── main.go
├── go.mod
├── .env.example
├── .gitignore
└── README.md
```

## ⚙️ Configuración

### 1. Crear archivo `.env`

Copia el archivo `.env.example` y renómbralo como `.env`:

```bash
cp .env.example .env
```

### 2. Configurar variables de entorno

Edita el archivo `.env` con tus credenciales de base de datos:

```env
# Configuración de la API
api_configuracion_HTTP_PORT=8082
api_configuracion_RUN_MODE=dev

# Configuración de PostgreSQL
api_configuracion_PG_USER=tu_usuario
api_configuracion_PG_PASS=tu_contraseña
api_configuracion_PGHOST=localhost
api_configuracion_PGPORT=5432
api_configuracion_PGDB=nombre_base_datos
api_configuracion_PGSCHEMA=configuracion
```

**Variables importantes:**

| Variable | Descripción | Ejemplo |
|----------|-------------|---------|
| `api_configuracion_HTTP_PORT` | Puerto en el que corre la API | 8082 |
| `api_configuracion_RUN_MODE` | Modo de ejecución (dev/prod) | dev |
| `api_configuracion_PG_USER` | Usuario de PostgreSQL | postgres |
| `api_configuracion_PG_PASS` | Contraseña de PostgreSQL | password123 |
| `api_configuracion_PGHOST` | Host de la base de datos | localhost |
| `api_configuracion_PGPORT` | Puerto de PostgreSQL | 5432 |
| `api_configuracion_PGDB` | Nombre de la base de datos | mi_tienda |
| `api_configuracion_PGSCHEMA` | Esquema PostgreSQL | configuracion |

### 3. Crear base de datos y esquema

Conéctate a PostgreSQL y ejecuta:

```sql
-- Crear base de datos
CREATE DATABASE nombre_base_datos;

-- Crear esquema
CREATE SCHEMA configuracion;

-- Crear tablas
CREATE TABLE configuracion.sistema (
    id_sistema SERIAL PRIMARY KEY,
    nombre_empresa VARCHAR(255) NOT NULL,
    descripcion TEXT,
    url VARCHAR(255),
    telefono VARCHAR(20),
    correo VARCHAR(255),
    logo_url VARCHAR(255),
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE configuracion.seguridad (
    id_seguridad SERIAL PRIMARY KEY,
    tipo_autenticacion VARCHAR(100) NOT NULL,
    requiere_autenticacion BOOLEAN DEFAULT true,
    intentos_maximos_fallidos INTEGER,
    tiempo_bloqueo_minutos INTEGER,
    requiere_two_factor BOOLEAN DEFAULT false,
    encriptacion_datos BOOLEAN DEFAULT true,
    sesion_maxima_horas INTEGER,
    descripcion TEXT,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE configuracion.backup (
    id_backup SERIAL PRIMARY KEY,
    nombre_backup VARCHAR(255) NOT NULL,
    tipo_backup VARCHAR(100) NOT NULL,
    fecha_ejecucion TIMESTAMP NOT NULL,
    fecha_proximo_backup TIMESTAMP,
    frecuencia VARCHAR(100),
    ubicacion_guardado VARCHAR(500),
    tamano_mb DECIMAL(10,2),
    estado VARCHAR(50) NOT NULL,
    descripcion TEXT,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE configuracion.notificacion (
    id_notificacion SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    descripcion TEXT NOT NULL,
    tipo VARCHAR(100) NOT NULL,
    canal VARCHAR(100) NOT NULL,
    destinatario VARCHAR(255),
    fecha_envio TIMESTAMP,
    fecha_programada TIMESTAMP,
    estado VARCHAR(50) NOT NULL,
    intentos_envio INTEGER,
    mensaje_error TEXT,
    prioridad VARCHAR(50),
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## ▶️ Ejecución

### Desarrollo

```bash
go run main.go
```

La API estará disponible en: `http://localhost:8082`

### Compilación para producción

```bash
go build -o api_configuracion
./api_configuracion
```

## 📡 Endpoints

### Sistema

#### GET - Obtener todas las configuraciones
```
GET /v1/sistema
```

#### GET - Obtener configuración por ID
```
GET /v1/sistema/:id
```

#### POST - Crear nueva configuración
```
POST /v1/sistema
Content-Type: application/json
```

**Body:**
```json
{
  "NombreEmpresa": "Mi Empresa",
  "Descripcion": "Descripción de la empresa",
  "Url": "https://miempresa.com",
  "Telefono": "3123456789",
  "Correo": "info@miempresa.com",
  "LogoUrl": "https://ejemplo.com/logo.png",
  "Activo": true
}
```

#### PUT - Actualizar configuración
```
PUT /v1/sistema/:id
```

#### DELETE - Eliminar configuración
```
DELETE /v1/sistema/:id
```

---

### Seguridad

#### GET - Obtener todas las configuraciones de seguridad
```
GET /v1/seguridad
```

#### GET - Obtener configuración de seguridad por ID
```
GET /v1/seguridad/:id
```

#### POST - Crear configuración de seguridad
```
POST /v1/seguridad
Content-Type: application/json
```

**Body:**
```json
{
  "TipoAutenticacion": "OAuth2",
  "RequiereAutenticacion": true,
  "IntentosMaximosFallidos": 5,
  "TiempoBloqueoMinutos": 15,
  "RequiereTwoFactor": true,
  "EncriptacionDatos": true,
  "SesionMaximaHoras": 8,
  "Activo": true
}
```

#### PUT - Actualizar configuración de seguridad
```
PUT /v1/seguridad/:id
```

#### DELETE - Eliminar configuración de seguridad
```
DELETE /v1/seguridad/:id
```

---

### Backup

#### GET - Obtener todos los backups
```
GET /v1/backup
```

#### GET - Obtener backup por ID
```
GET /v1/backup/:id
```

#### POST - Crear configuración de backup
```
POST /v1/backup
Content-Type: application/json
```

**Body:**
```json
{
  "NombreBackup": "Backup Diario",
  "TipoBackup": "Completo",
  "FechaEjecucion": "2026-05-20T10:00:00Z",
  "FechaProximoBackup": "2026-05-21T10:00:00Z",
  "Frecuencia": "Diaria",
  "UbicacionGuardado": "/backups/completo",
  "TamanoMB": 1024.50,
  "Estado": "Exitoso",
  "Activo": true
}
```

#### PUT - Actualizar configuración de backup
```
PUT /v1/backup/:id
```

#### DELETE - Eliminar configuración de backup
```
DELETE /v1/backup/:id
```

---

### Notificación

#### GET - Obtener todas las notificaciones
```
GET /v1/notificacion
```

#### GET - Obtener notificación por ID
```
GET /v1/notificacion/:id
```

#### POST - Crear notificación
```
POST /v1/notificacion
Content-Type: application/json
```

**Body:**
```json
{
  "Titulo": "Backup completado",
  "Descripcion": "El backup diario se completó exitosamente",
  "Tipo": "Informativa",
  "Canal": "Email",
  "Destinatario": "admin@empresa.com",
  "Estado": "Pendiente",
  "Prioridad": "Alta",
  "Activo": true
}
```

#### PUT - Actualizar notificación
```
PUT /v1/notificacion/:id
```

#### DELETE - Eliminar notificación
```
DELETE /v1/notificacion/:id
```

## 📁 Estructura del Proyecto

```
api_configuracion/
├── conf/
│   └── app.conf              # Configuración de la aplicación
├── controllers/
│   ├── sistema.go            # Controlador de Sistema
│   ├── seguridad.go          # Controlador de Seguridad
│   ├── backup.go             # Controlador de Backup
│   └── notificacion.go       # Controlador de Notificación
├── models/
│   ├── sistema.go            # Modelo de Sistema
│   ├── seguridad.go          # Modelo de Seguridad
│   ├── backup.go             # Modelo de Backup
│   └── notificacion.go       # Modelo de Notificación
├── routers/
│   └── router.go             # Definición de rutas
├── main.go                   # Punto de entrada
├── go.mod                    # Dependencias de Go
├── .env.example              # Variables de entorno (ejemplo)
├── .gitignore                # Archivos ignorados por Git
└── README.md                 # Este archivo
```

## 🛠️ Tecnologías Utilizadas

- **Beego 2.3.10** - Framework web de alto rendimiento
- **PostgreSQL 12+** - Base de datos relacional
- **Go 1.24.2** - Lenguaje de programación
- **pgx** - Driver PostgreSQL

## 🔒 Seguridad

### Protecciones implementadas:

1. **Variables de entorno** - Las credenciales se cargan desde `.env`
2. **CORS** - Control de acceso entre orígenes
3. **Validación de entrada** - Validación en controladores
4. **Errores manejados** - Respuestas HTTP apropiadas

### Recomendaciones:

- ⚠️ **NUNCA** commitear el archivo `.env` al repositorio
- ✅ Usar `.env.example` como referencia
- ✅ Proteger el acceso a la base de datos
- ✅ Usar HTTPS en producción

## 📝 Ejemplo completo de uso

```bash
# 1. Iniciar la API
go run main.go

# 2. En otra terminal, crear una configuración del sistema
curl -X POST http://localhost:8082/v1/sistema \
  -H "Content-Type: application/json" \
  -d '{
    "NombreEmpresa": "Mi Empresa",
    "Url": "https://miempresa.com",
    "Activo": true
  }'

# 3. Obtener todas las configuraciones del sistema
curl http://localhost:8082/v1/sistema

# 4. Crear una configuración de seguridad
curl -X POST http://localhost:8082/v1/seguridad \
  -H "Content-Type: application/json" \
  -d '{
    "TipoAutenticacion": "OAuth2",
    "RequiereAutenticacion": true,
    "Activo": true
  }'

# 5. Crear una notificación
curl -X POST http://localhost:8082/v1/notificacion \
  -H "Content-Type: application/json" \
  -d '{
    "Titulo": "Bienvenido",
    "Descripcion": "Bienvenido al sistema",
    "Tipo": "Informativa",
    "Canal": "Email",
    "Estado": "Pendiente"
  }'
```

## 🐛 Solución de Problemas

### Error: "Error cargando archivo .env"
- Asegúrate de que existe el archivo `.env` en la raíz del proyecto
- Verifica que las variables estén correctamente formateadas

### Error: "Connection refused" a PostgreSQL
- Verifica que PostgreSQL esté corriendo
- Comprueba las credenciales en `.env`
- Verifica el host y puerto

### Error: "Schema no existe"
- Verifica que el esquema `configuracion` exista en tu base de datos
- Ejecuta el script SQL de creación de tablas

## 📞 Soporte

Para reportar problemas o sugerencias, crea un issue en el repositorio.

## 📄 Licencia

Este proyecto está bajo la licencia Apache 2.0.

---

**Última actualización:** Mayo 2026
