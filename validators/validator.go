package validators

import (
	"regexp"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.ParamTagMap["min"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		int, _ := govalidator.ToInt(str)
		param, _ := govalidator.ToInt(params[0])
		return int >= param
	})

	govalidator.ParamTagRegexMap["min"] = regexp.MustCompile("^min\\((\\d+)\\)$")

	govalidator.ParamTagMap["max"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		int, _ := govalidator.ToInt(str)
		param, _ := govalidator.ToInt(params[0])
		return int <= param
	})

	govalidator.ParamTagRegexMap["max"] = regexp.MustCompile("^max\\((\\d+)\\)$")
}

func BindAndValidate(form interface{}) error {
	_, err := govalidator.ValidateStruct(form)
	if err != nil {
		errs := err.(govalidator.Errors).Errors()
		return errs[0]
	}
	return nil
}
