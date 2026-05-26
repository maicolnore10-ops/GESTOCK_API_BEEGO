package main

// @Title API Configuración
// @Description API REST para gestionar la configuración del sistema, seguridad, backups y notificaciones
// @Contact.name Soporte API
// @Contact.email soporte@example.com
// @TermsOfServiceUrl http://example.com/terms
// @License.name MIT
// @License.url http://opensource.org/licenses/MIT
// @APIVersion 1.0.0
// @Host localhost:8082
// @BasePath /
// @Schemes http https
// @Accept json
// @Produce json

import (
	"fmt"
	"net/url"

	_ "api_configuracion/routers" // Importa tus rutas
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// 1. Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		fmt.Println("Advertencia: No se pudo cargar el archivo .env")
	}

	// 2. Cargar configuración de Beego
	beego.LoadAppConfig("ini", "conf/app.conf")

	// 3. Obtener credenciales
	pgUser := beego.AppConfig.DefaultString("PGuser", "postgres")
	pgPass := beego.AppConfig.DefaultString("PGpass", "postgres")
	pgHost := beego.AppConfig.DefaultString("PGhost", "localhost")
	pgPort := beego.AppConfig.DefaultString("PGport", "5432")
	pgDb := beego.AppConfig.DefaultString("PGdb", "Gestock_db")
	pgSchema := beego.AppConfig.DefaultString("PGschema", "configuracion")

	// 4. Construir DSN de forma segura
	safePass := url.QueryEscape(pgPass)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		pgUser, safePass, pgHost, pgPort, pgDb, pgSchema)

	// 5. Registrar Base de Datos
	if err := orm.RegisterDataBase("default", "postgres", dsn); err != nil {
		panic("Error registrando base de datos: " + err.Error())
	}

	// 6. Configuración CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
	}))

	// 7. Configuración Swagger (Beego 2.x)
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.SetStaticPath("/swagger", "swagger")
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	beego.Run()
}