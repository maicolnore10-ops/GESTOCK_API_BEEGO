// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api_compras/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/proveedores",
			beego.NSInclude(
				&controllers.ProveedoresController{},
			),
		),

		beego.NSNamespace("/proveedor_productos",
			beego.NSInclude(
				&controllers.ProveedorProductosController{},
			),
		),

		beego.NSNamespace("/guias_entrada",
			beego.NSInclude(
				&controllers.GuiasEntradaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
