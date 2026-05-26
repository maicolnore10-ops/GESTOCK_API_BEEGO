package controllers

import (
	"api_configuracion/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// NotificacionController operations for Notificacion
type NotificacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *NotificacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Notificacion
// @Param	body		body 	models.Notificacion	true		"body for Notificacion content"
// @Success 201 {int} models.Notificacion
// @Failure 403 body is empty
// @router / [post]
func (c *NotificacionController) Post() {
	var v models.Notificacion
	if err := json.Unmarshal(getRequestBody(&c.Controller), &v); err == nil {
		if v.Titulo == "" {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"Message": "El título de la notificación es requerido",
				"status":  400,
				"success": false,
			}
			c.ServeJSON()
			return
		}

		if v.Tipo == "" {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"Message": "El tipo de notificación es requerido",
				"status":  400,
				"success": false,
			}
			c.ServeJSON()
			return
		}

		if _, err := models.AddNotificacion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{
				"Data":    v,
				"Message": "Notificación creada exitosamente",
				"status":  201,
				"success": true,
			}
		} else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"Message": "Error al crear la notificación: " + err.Error(),
				"status":  400,
				"success": false,
			}
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Message": "Datos inválidos: " + err.Error(),
			"status":  400,
			"success": false,
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Notificacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Notificacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *NotificacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetNotificacionById(id)
	if err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{
			"Message": "Notificación no encontrada",
			"status":  404,
			"success": false,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"Message": "Notificación encontrada",
			"status":  200,
			"success": true,
			"data":    v,
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Notificacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Notificacion
// @Failure 403
// @router / [get]
func (c *NotificacionController) GetAll() {
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
				c.Ctx.Output.SetStatus(400)
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllNotificacion(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Message": "Error al consultar notificaciones: " + err.Error(),
			"status":  400,
			"success": false,
		}
	} else {
		if l == nil {
			c.Data["json"] = map[string]interface{}{
				"Message": "No hay notificaciones disponibles",
				"status":  200,
				"success": true,
				"data":    []interface{}{},
			}
		} else {
			c.Data["json"] = map[string]interface{}{
				"Message": "Notificaciones encontradas",
				"status":  200,
				"success": true,
				"data":    l,
			}
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Notificacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Notificacion	true		"body for Notificacion content"
// @Success 200 {object} models.Notificacion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *NotificacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, errId := strconv.Atoi(idStr)
	if errId != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"Message": "ID inválido", "status": 400, "success": false}
		c.ServeJSON()
		return
	}
	// comprobar existencia
	_, err := models.GetNotificacionById(id)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"Message": "Notificación no encontrada", "status": 404, "success": false}
		c.ServeJSON()
		return
	}
	var v models.Notificacion
	if err := json.Unmarshal(getRequestBody(&c.Controller), &v); err == nil {
		if v.Titulo == "" {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"Message": "El título de la notificación no puede estar vacío",
				"status":  400,
				"success": false,
			}
			c.ServeJSON()
			return
		}
		v.Id = id
		if err := models.UpdateNotificacionById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{
				"Data":    v,
				"Message": "Notificación actualizada exitosamente",
				"status":  200,
				"success": true,
			}
		} else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{
				"Message": "Error al actualizar la notificación: " + err.Error(),
				"status":  400,
				"success": false,
			}
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Message": "Datos inválidos: " + err.Error(),
			"status":  400,
			"success": false,
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Notificacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *NotificacionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, errId := strconv.Atoi(idStr)
	if errId != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"Message": "ID inválido", "status": 400, "success": false}
		c.ServeJSON()
		return
	}
	// comprobar existencia
	_, err := models.GetNotificacionById(id)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"Message": "Notificación no encontrada", "status": 404, "success": false}
		c.ServeJSON()
		return
	}
	if err := models.DeleteNotificacion(id); err == nil {
		c.Data["json"] = map[string]interface{}{
			"Message": "Notificación eliminada exitosamente",
			"status":  200,
			"success": true,
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"Message": "Error al eliminar la notificación: " + err.Error(),
			"status":  400,
			"success": false,
		}
	}
	c.ServeJSON()
}
