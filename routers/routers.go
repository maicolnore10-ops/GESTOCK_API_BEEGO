package routers

import (
	"api_configuracion/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// Registro manual: el más estable y profesional
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/sistema",
			beego.NSRouter("/", &controllers.SistemaController{}, "post:Post;get:GetAll"),
			beego.NSRouter("/:id", &controllers.SistemaController{}, "get:GetOne;put:Put;delete:Delete"),
		),
		// Registra aquí tus otros controladores (ejemplo):
		// beego.NSNamespace("/notificacion",
		// 	beego.NSRouter("/:id", &controllers.NotificacionController{}, "get:GetOne"),
		// ),
	)
	beego.AddNamespace(ns)
}