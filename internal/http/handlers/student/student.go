package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	typez "github.com/marsyg/studentApi/internal/types"
	"github.com/marsyg/studentApi/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student typez.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		// error validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidateError(validateErrs))
		}
		slog.Info("creating a student ")
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
		w.Write([]byte("welcome to the server"))
	}
}
