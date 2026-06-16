// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"APICRUD_GESTOCK/EMPRESA/empresa/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/unidades_medida",
			beego.NSInclude(
				&controllers.UnidadesMedidaController{},
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

		beego.NSNamespace("/empresas",
			beego.NSInclude(
				&controllers.EmpresasController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
