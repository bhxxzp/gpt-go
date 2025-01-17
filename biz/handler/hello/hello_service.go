// Code generated by hertz generator.

package hello

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hello "gpt/biz/model/hello"
)

// Hello .
// @router /hello [GET]
func Hello(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hello.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		panic(err)
		return
	}

	resp := new(hello.HelloResp)
	resp.RespBody = c.GetString("zp")
	c.JSON(consts.StatusOK, resp)
}
