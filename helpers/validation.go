package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

type ValidationErrorResponse struct {
}

func NewValidator(options ...validator.Option) Validator {
	return Validator{validator.New(validator.WithRequiredStructEnabled())}
}

func (v *Validator) ValidateStruct(s interface{}) error {
	if err := v.validator.Struct(s); err != nil {
		return err
	}
	return nil
}

func (v *Validator) ParseValidationError(err error) map[string]string{
	details := make(map[string]string)
	switch typedError := err.(type) {
	case validator.ValidationErrors:
		for _, e := range typedError {
			details[strings.ToLower(e.Field())] = parseFieldError(e)
		}
	case *json.UnmarshalTypeError:
		details[typedError.Field] = parseMarshalingError(*typedError)
	case *strconv.NumError:
		details["error"] = "please provide a valid number"
	default:
		if err.Error() == "EOF" {
			details["error"] = "please provide all required data"
		} else {
			details["error"] = err.Error()
		}
	}
	return details
}

func parseFieldError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "field required"
	case "min":
		return fmt.Sprintf("minimum (length/value) allowed is %s", err.Param())
	case "max":
		return fmt.Sprintf("maximum (length/value) allowed is %s", err.Param())
	case "oneof":
		return fmt.Sprintf("only allowed values are %s", err.Param())
	case "email":
		return "invalid email address"
	case "datetime":
		return "invalid date time"
	case "unique":
		return fmt.Sprintf("%s should contain only unique values", err.Field())
	case "eqcsfield":
		return fmt.Sprintf("%s must equal %s", err.Field(), err.Param())
	default:
		return err.Error()

	}
}

func parseMarshalingError(err json.UnmarshalTypeError) string {
	return fmt.Sprintf("%s must be a %s", err.Field, err.Type.String())
}
