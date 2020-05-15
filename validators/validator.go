package validators

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.ParamTagMap["min"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		int, _ := govalidator.ToInt(str)
		param, _ := govalidator.ToInt(params[0])
		return int >= param
	})

	govalidator.ParamTagMap["max"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		int, _ := govalidator.ToInt(str)
		param, _ := govalidator.ToInt(params[0])
		return int <= param
	})
}

func BindAndValidate(form interface{}) error {
	_, err := govalidator.ValidateStruct(form)
	if err != nil {
		errs := err.(govalidator.Errors).Errors()
		return errs[0]
	}
	return nil
}
