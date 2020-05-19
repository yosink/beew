package controllers

import (
	"beew/repositories"
	"beew/services"
	"beew/utils"
	"beew/validators"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
	service services.IUserService
}

func (ctl *AuthController) Prepare() {
	userRepo := repositories.NewUserRepository()
	ctl.service = services.NewUserService(userRepo)
}

func (ctl *AuthController) Token() {
	var (
		auth validators.AuthUser
		app  = utils.B{
			C: ctl.Ctx,
		}
	)
	err := ctl.ParseForm(&auth)
	if err != nil {
		app.JsonResponse(http.StatusBadRequest, 400, "parse error", nil)
		return
	}
	err = validators.BindAndValidate(auth)
	if err != nil {
		app.JsonResponse(http.StatusBadRequest, 400, "invalid params", nil)
		return
	}

	user, err := ctl.service.GetUserByPhone(auth.Account)
	if err != nil {
		fmt.Println(err)
		app.JsonResponse(http.StatusOK, 404, "user not found", nil)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
	if err != nil {
		app.JsonResponse(http.StatusOK, 401, "password not match", nil)
		return
	}
	data, err := utils.GenerateJwtToken(int(user.ID))
	if err != nil {
		fmt.Println()
		fmt.Printf("%v", err)
		app.JsonResponse(http.StatusOK, 404, "user not found", nil)
		return
	}
	app.JsonResponse(http.StatusOK, 200, "", data)
	return
}
