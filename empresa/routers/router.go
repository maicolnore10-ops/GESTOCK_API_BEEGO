// @APIVersion 1.0.0
// @Title EMPRESA API
// @Description API de gestión de empresa, bodegas, moneda y unidades de medida
// @Contact soporte@empresa.com
// @TermsOfServiceUrl http://empresa.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"EMPRESA/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/empresas",
			beego.NSInclude(
				&controllers.EmpresasController{},
			),
		),

		beego.NSNamespace("/bodegas",
			beego.NSInclude(
				&controllers.BodegasController{},
			),
		),

		beego.NSNamespace("/moneda",
			beego.NSInclude(
				&controllers.MonedaController{},
			),
		),

		beego.NSNamespace("/unidades_medida",
			beego.NSInclude(
				&controllers.UnidadesMedidaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
