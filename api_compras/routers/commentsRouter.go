package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"] = append(beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"] = append(beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"] = append(beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"] = append(beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"] = append(beego.GlobalControllerRouter["api_compras/controllers:GuiasEntradaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedorProductosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"] = append(beego.GlobalControllerRouter["api_compras/controllers:ProveedoresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
