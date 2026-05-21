// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"GESTOCK_API_BEEGO_INVENTARIO/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/categoria_producto",
			beego.NSInclude(
				&controllers.CategoriaProductoController{},
			),
		),

		beego.NSNamespace("/productos",
			beego.NSInclude(
				&controllers.ProductosController{},
			),
		),

		beego.NSNamespace("/inventario",
			beego.NSInclude(
				&controllers.InventarioController{},
			),
		),

		beego.NSNamespace("/inventario_productos",
			beego.NSInclude(
				&controllers.InventarioProductosController{},
			),
		),

		beego.NSNamespace("/tipos_movimiento",
			beego.NSInclude(
				&controllers.TiposMovimientoController{},
			),
		),

		beego.NSNamespace("/movimientos",
			beego.NSInclude(
				&controllers.MovimientosController{},
			),
		),

		beego.NSNamespace("/ajustes",
			beego.NSInclude(
				&controllers.AjustesController{},
			),
		),

		beego.NSNamespace("/transferencia",
			beego.NSInclude(
				&controllers.TransferenciaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
