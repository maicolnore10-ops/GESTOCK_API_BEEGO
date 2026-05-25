package controllers

import (
	"GESTOCK_API_BEEGO_INVENTARIO/models"
	"encoding/json"
	"errors" // Importación utilizada para comparar los errores del ORM
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// AjustesController operations for Ajustes
type AjustesController struct {
	beego.Controller
}

// URLMapping ...
func (c *AjustesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Ajustes
// @Param   body        body    models.Ajustes  true        "body for Ajustes content"
// @Success 201 {int} models.Ajustes
// @Failure 403 body is empty
// @router / [post]
func (c *AjustesController) Post() {
	var v models.Ajustes
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddAjustes(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  201,
				"Message": "Ajuste creado exitosamente",
				"data":    v,
			}
		} else {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"Message": "Error al insertar el ajuste: " + err.Error(),
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Error en el formato de la petición: " + err.Error(),
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Ajustes by id
// @Param   id      path    string  true        "The key for staticblock"
// @Success 200 {object} models.Ajustes
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AjustesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAjustesById(id)
	
	if err != nil {
		// Uso de errors.Is para verificar si el registro no existe en la base de datos
		if errors.Is(err, orm.ErrNoRows) {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  404,
				"Message": "Error en el servidor GetOne: El registro con el ID solicitado no existe",
			}
		} else {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"Message": "Error en la solicitud: " + err.Error(),
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"Message": "Ajuste encontrado con éxito",
			"data":    v,
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Ajustes
// @Param   query   query   string  false   "Filter. e.g. col1:v1,col2:v2 ..."
// @Param   fields  query   string  false   "Fields returned. e.g. col1,col2 ..."
// @Param   sortby  query   string  false   "Sorted-by fields. e.g. col1,col2 ..."
// @Param   order   query   string  false   "Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param   limit   query   string  false   "Limit the size of result set. Must be an integer"
// @Param   offset  query   string  false   "Start position of result set. Must be an integer"
// @Success 200 {object} models.Ajustes
// @Failure 403
// @router / [get]
func (c *AjustesController) GetAll() {
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
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"status":  400,
					"Message": "Error: invalid query key/value pair",
				}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAjustes(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Error al obtener la lista de ajustes: " + err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"Message": "Usuarios encontrados",
			"data":    l,
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Ajustes
// @Param   id      path    string  true        "The id you want to update"
// @Param   body        body    models.Ajustes  true        "body for Ajustes content"
// @Success 200 {object} models.Ajustes
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AjustesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Ajustes{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAjustesById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  200,
				"Message": "Ajuste actualizado exitosamente",
				"data":    v,
			}
		} else {
			// Uso de errors.Is para controlar si el ID a modificar no existe en BD
			if errors.Is(err, orm.ErrNoRows) {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"status":  404,
					"Message": "Error al actualizar: El registro que intenta modificar no existe",
				}
			} else {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"status":  400,
					"Message": "Error al actualizar el ajuste: " + err.Error(),
				}
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Error en el formato de la petición: " + err.Error(),
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Ajustes
// @Param   id      path    string  true        "The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AjustesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	
	if err := models.DeleteAjustes(id); err == nil {
		c.Data["json"] = map[string]interface{}{
			"success": true, 
			"status":  200, 
			"Message": "Peticion existosa", 
			"data":    id,
		}
	} else {
		// Aquí usamos errors.Is para validar si falló porque el registro no existía en la BD
		if errors.Is(err, orm.ErrNoRows) {
			c.Data["json"] = map[string]interface{}{
				"success": false, 
				"status":  400, 
				"Message": "Error en el servidor delete: no existe el dato que desea eliminar",
			}
		} else {
			c.Data["json"] = map[string]interface{}{
				"success": false, 
				"status":  500, 
				"Message": "Error interno al intentar eliminar: " + err.Error(),
			}
		}
	}
	c.ServeJSON()
}