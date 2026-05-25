// @APIVersion 1.0.0
// @Title GESTOCK API
// @Description API de autenticacion y gestion de usuarios para GESTOCK
// @Contact soporte@gestock.com
// @TermsOfServiceUrl http://gestock.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"GESTOCK_API_BEEGO/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/roles",
			beego.NSInclude(
				&controllers.RolesController{},
			),
		),

		beego.NSNamespace("/usuarios",
			beego.NSInclude(
				&controllers.UsuariosController{},
			),
		),

		beego.NSNamespace("/contrasena_hash",
			beego.NSInclude(
				&controllers.ContrasenaHashController{},
			),
		),

		beego.NSNamespace("/recuperacion_contrasena",
			beego.NSInclude(
				&controllers.RecuperacionContrasenaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
