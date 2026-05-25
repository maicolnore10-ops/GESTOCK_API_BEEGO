package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:AjustesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:CategoriaProductoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:InventarioProductosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:MovimientosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TiposMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO_INVENTARIO/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
