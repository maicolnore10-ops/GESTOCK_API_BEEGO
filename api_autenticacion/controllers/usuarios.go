package controllers

import (
	"GESTOCK_API_BEEGO/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"github.com/beego/beego/v2/client/orm"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// UsuariosController operations for Usuarios
type UsuariosController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsuariosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Usuarios
// @Param	body	body	models.Usuarios	true	"body for Usuarios content"
// @Success 201 {int} models.Usuarios
// @Failure 403 body is empty
// @router / [post]
func (c *UsuariosController) Post() {
	var v models.Usuarios
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddUsuarios(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  201,
				"Message": "Usuario creado exitosamente",
				"data":    v,
			}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"Message": "Error al crear el usuario: " + err.Error(),
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Datos inválidos: " + err.Error(),
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Usuarios by id
// @Param	id	path	string	true	"The key for staticblock"
// @Success 200 {object} models.Usuarios
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsuariosController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUsuariosById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Usuario no encontrado o parámetro incorrecto",
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"Message": "Usuario encontrado",
			"data":    v,
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Usuarios
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Usuarios
// @Failure 403
// @router / [get]
func (c *UsuariosController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

  l, err := models.GetAllUsuarios(query, fields, sortby, order, offset, limit)
    if err != nil {
        c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "message": err.Error()}
    } else {
        if len(l) > 0 {
            o := orm.NewOrm()
            
            // Recorremos los usuarios e inyectamos su respectivo Rol manualmente
            for _, v := range l {
                if usuario, ok := v.(*models.Usuarios); ok {
                    // Carga la relación 'IdRol' definida en el struct de tu modelo
                    o.LoadRelated(usuario, "IdRol")
                }
            }

            c.Data["json"] = map[string]interface{}{
                "success": true,
                "status":  200,
                "message": "Registros de usuarios encontrados",
                "data":    l,
            }
        } else {
            c.Data["json"] = map[string]interface{}{
                "success": false,
                "status":  400,
                "message": "Registros de usuarios no encontrados",
            }
        }
    }
    c.ServeJSON()

}
// Put ...
// @Title Put
// @Description update the Usuarios
// @Param	id	path	string	true	"The id you want to update"
// @Param	body	body	models.Usuarios	true	"body for Usuarios content"
// @Success 200 {object} models.Usuarios
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsuariosController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Usuarios{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUsuariosById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  200,
				"Message": "Usuario actualizado exitosamente",
				"data":    v,
			}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"Message": "Error al actualizar el usuario",
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Datos inválidos",
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Usuarios
// @Param	id	path	string	true	"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsuariosController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUsuarios(id); err == nil {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"Message": "Usuario eliminado exitosamente",
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Error al eliminar el usuario: " + err.Error(),
		}
	}
	c.ServeJSON()
}
