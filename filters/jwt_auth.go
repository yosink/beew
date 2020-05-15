package filters

import (
	"beew/utils"
	"strings"

	"github.com/astaxie/beego/context"
)

func JwtAuth(ctx *context.Context) {
	token := GetToken(ctx)
	if token != "" {
		claims, err := utils.ParseToken(token)
		if err != nil {
			ctx.Output.SetStatus(401)
			ctx.Output.JSON(map[string]interface{}{
				"code":    401,
				"message": "Unauthorized.",
			}, false, false)
		}
		if claims.Uid != 1 {
			ctx.Output.SetStatus(401)
			ctx.Output.JSON(map[string]interface{}{
				"code":    401,
				"message": "Unauthorized.",
			}, false, false)
		}
	} else {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]interface{}{
			"code":    401,
			"message": "Unauthorized.",
		}, false, false)
	}
}

func GetToken(c *context.Context) string {
	token := c.Input.Query("token")
	if token != "" {
		return token
	}
	hToken := c.Input.Header("Authorization")
	if hToken != "" && strings.HasPrefix(hToken, "Bearer ") {
		return hToken[7:]
	}
	return ""
}
