package controllers

import (
	"io"

	beego "github.com/beego/beego/v2/server/web"
)

func getRequestBody(c *beego.Controller) []byte {
	body := c.Ctx.Input.RequestBody
	if len(body) == 0 && c.Ctx.Request != nil && c.Ctx.Request.Body != nil {
		if b, err := io.ReadAll(c.Ctx.Request.Body); err == nil {
			body = b
		}
	}
	return body
}
