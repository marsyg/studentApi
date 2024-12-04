package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Resposne struct {
	Status string
	Error  string
}

const (
	StatusOk    = "Ok"
	StatusError = "Error"
)

// error written after the func args -  shows that this func can even return error
// or you can handle the error the  explicitly so that func will not return error
func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
func GeneralError(err error) Resposne {
	return Resposne{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidateError(errs validator.ValidationErrors) Resposne {
	var errsMsg []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errsMsg = append(errsMsg, fmt.Sprintf("field %s is required feild", err.Field()))
		default:
			errsMsg = append(errsMsg, fmt.Sprintf("field %s is required feild", err.Field()))

		}
	}
	return Resposne{
		Status: StatusError,
		Error:  strings.Join(errsMsg, ", "),
	}
}
