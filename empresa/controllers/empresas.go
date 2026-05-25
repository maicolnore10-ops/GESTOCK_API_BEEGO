package controllers

import (
	"EMPRESA/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// EmpresasController operations for Empresas
type EmpresasController struct {
	beego.Controller
}

// URLMapping ...
func (c *EmpresasController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Empresas
// @Param	body	body	models.Empresas	true	"body for Empresas content"
// @Success 201 {int} models.Empresas
// @Failure 403 body is empty
// @router / [post]
func (c *EmpresasController) Post() {
	var v models.Empresas
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddEmpresas(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  201,
				"Message": "Empresa creada exitosamente",
				"data":    v,
			}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"Message": "Error al crear la empresa: " + err.Error(),
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
// @Description get Empresas by id
// @Param	id	path	string	true	"The key for staticblock"
// @Success 200 {object} models.Empresas
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EmpresasController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEmpresasById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Empresa no encontrada o parámetro incorrecto",
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"Message": "Empresa encontrada",
			"data":    v,
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Empresas
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Empresas
// @Failure 403
// @router / [get]
func (c *EmpresasController) GetAll() {
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

	l, err := models.GetAllEmpresas(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": err.Error(),
		}
	} else {
		if l == nil {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"Message": "Empresas no encontradas",
			}
		} else {
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  200,
				"Message": "Empresas encontradas",
				"data":    l,
			}
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Empresas
// @Param	id	path	string	true	"The id you want to update"
// @Param	body	body	models.Empresas	true	"body for Empresas content"
// @Success 200 {object} models.Empresas
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EmpresasController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Empresas{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateEmpresasById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  200,
				"Message": "Empresa actualizada exitosamente",
				"data":    v,
			}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"Message": "Error al actualizar la empresa",
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
// @Description delete the Empresas
// @Param	id	path	string	true	"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EmpresasController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEmpresas(id); err == nil {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"Message": "Empresa eliminada exitosamente",
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Error al eliminar la empresa: " + err.Error(),
		}
	}
	c.ServeJSON()
}
