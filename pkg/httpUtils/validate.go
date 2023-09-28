package httpUtils

import (
	"fmt"
	"github.com/gookit/validate"
	"net/http"
)

func ValidateStruct(w http.ResponseWriter, req interface{}) error {
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
	})

	v := validate.Struct(req)

	if !v.Validate() {
		var errors []Error
		for e := range v.Errors.All() {
			for _, v := range v.Errors.Field(e) {
				errors = append(errors, Error{
					Field:   e,
					Message: v,
				})
			}
		}

		ErrorJSON(w, "Invalid request body", http.StatusBadRequest, errors...)
		return fmt.Errorf("invalid request body")
	}

	return nil
}
