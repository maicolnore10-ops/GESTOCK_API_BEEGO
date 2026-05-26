package controllers

import (
	"api_configuracion/models"
	"encoding/json"
	"strconv"
	beego "github.com/beego/beego/v2/server/web"
)

type SistemaController struct {
	beego.Controller
}

// @router / [post]
func (c *SistemaController) Post() {
	var v models.Sistema
	if err := json.Unmarshal(getRequestBody(&c.Controller), &v); err == nil {
		if _, err := models.AddSistema(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"success": true, "data": v}
		} else {
			c.Data["json"] = map[string]interface{}{"success": false, "message": err.Error()}
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "message": err.Error()}
	}
	c.ServeJSON()
}

// @router / [get]
func (c *SistemaController) GetAll() {
	l, _ := models.GetAllSistema(nil, nil, nil, nil, 0, 100)
	c.Data["json"] = map[string]interface{}{"success": true, "data": l}
	c.ServeJSON()
}

// @router /:id [get]
func (c *SistemaController) GetOne() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if v, err := models.GetSistemaById(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "data": v}
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "No encontrado"}
	}
	c.ServeJSON()
}

// @router /:id [put]
func (c *SistemaController) Put() {
	id, errId := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if errId != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	_, err := models.GetSistemaById(id)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "No encontrado"}
		c.ServeJSON()
		return
	}
	var v models.Sistema
	if err := json.Unmarshal(getRequestBody(&c.Controller), &v); err == nil {
		v.Id = int64(id)
		if err := models.UpdateSistemaById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"success": true, "message": "Actualizado"}
		} else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{"success": false, "message": err.Error()}
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "message": err.Error()}
	}
	c.ServeJSON()
}

// @router /:id [delete]
func (c *SistemaController) Delete() {
	id, errId := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if errId != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	// comprobar existencia
	_, err := models.GetSistemaById(id)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "No encontrado"}
		c.ServeJSON()
		return
	}
	if err := models.DeleteSistema(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "message": "Eliminado"}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "message": err.Error()}
	}
	c.ServeJSON()
}