package filters

import (
	"beew/repositories"
	"beew/services"
	"beew/utils"
	"net/http"
	"strings"

	"github.com/astaxie/beego/context"
)

func JwtAuth(ctx *context.Context) {
	app := utils.B{
		C: ctx,
	}
	token := GetToken(ctx)
	if token != "" {
		claims, err := utils.ParseToken(token)
		if err != nil {
			app.JsonResponse(http.StatusUnauthorized, 401, "Unauthorized.", nil)
			return
		}
		repo := repositories.NewUserRepository()
		userService := services.NewUserService(repo)
		exists, err := userService.ExistsByID(claims.Uid)
		if !exists {
			app.JsonResponse(http.StatusUnauthorized, 401, "Unauthorized.", nil)
			return
		}
	} else {
		app.JsonResponse(http.StatusUnauthorized, 401, "Unauthorized.", nil)
		return
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
