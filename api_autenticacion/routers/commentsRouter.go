package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	// RolesController
	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// UsuariosController
	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:UsuariosController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// ContrasenaHashController
	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:ContrasenaHashController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// RecuperacionContrasenaController
	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["GESTOCK_API_BEEGO/controllers:RecuperacionContrasenaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}
