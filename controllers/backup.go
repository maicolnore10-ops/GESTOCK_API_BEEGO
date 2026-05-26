package controllers

import (
	"api_configuracion/models"
	"encoding/json"
	"strconv"
	"time"
	beego "github.com/beego/beego/v2/server/web"
)

type BackupController struct {
	beego.Controller
}

// @router / [post]
func (c *BackupController) Post() {
	var v models.Backup
	if err := json.Unmarshal(getRequestBody(&c.Controller), &v); err == nil {
		// Validaciones básicas
		if v.NombreBackup == "" || v.TipoBackup == "" {
			c.CustomAbort(400, "Nombre y Tipo de Backup son requeridos")
		}
		// Si no se proporcionó FechaEjecucion, usar ahora
		if v.FechaEjecucion.IsZero() {
			v.FechaEjecucion = time.Now()
		}
		if _, err := models.AddBackup(&v); err == nil {
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

// @router /:id [get]
func (c *BackupController) GetOne() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if v, err := models.GetBackupById(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "data": v}
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "No encontrado"}
	}
	c.ServeJSON()
}

// @router / [get]
func (c *BackupController) GetAll() {
	// Simplificado: remueve la lógica compleja si no es necesaria
	l, _ := models.GetAllBackup(nil, nil, nil, nil, 0, 100)
	c.Data["json"] = map[string]interface{}{"success": true, "data": l}
	c.ServeJSON()
}

// @router /:id [put]
func (c *BackupController) Put() {
	id, errId := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if errId != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	_, err := models.GetBackupById(id)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "No encontrado"}
		c.ServeJSON()
		return
	}
	var v models.Backup
	if err := json.Unmarshal(getRequestBody(&c.Controller), &v); err == nil {
		v.Id = id
		if err := models.UpdateBackupById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"success": true}
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
func (c *BackupController) Delete() {
	id, errId := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if errId != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "ID inválido"}
		c.ServeJSON()
		return
	}
	// comprobar existencia
	_, err := models.GetBackupById(id)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]interface{}{"success": false, "message": "No encontrado"}
		c.ServeJSON()
		return
	}
	if err := models.DeleteBackup(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"success": false, "message": err.Error()}
	}
	c.ServeJSON()
}