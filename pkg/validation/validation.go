package validation

import (
	validationDictionary "beta/pkg/dictionary"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func GetErrors(errorList validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)
	dictionary := new(validationDictionary.ValidationDictionary)
	dictionary = dictionary.Init()

	for _, f := range errorList {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Error())
		}
		errs[f.Field()] = dictionary.Get(err, "ddddd")
	}

	return errs
}
