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

	govalidator.TagMap["mobile"] = govalidator.Validator(func(str string) bool {
		regx := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
		compile := regexp.MustCompile(regx)
		return compile.MatchString(str)
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
