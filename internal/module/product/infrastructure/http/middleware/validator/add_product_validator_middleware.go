package validator

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AddProductValidatorInputDto struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description"`
	Price       int    `json:"price" validate:"required,min=1,max=1000000"`
	Quantity    int    `json:"quantity" validate:"required,min=1,max=1000000"`
	Status      bool   `json:"status" validate:"boolean"`
}

func AddProductValidatorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input AddProductValidatorInputDto

		// Ler o corpo da requisição em um novo reader
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Criar um novo reader para o corpo da requisição
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		if err := validator.New().Struct(input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// Passar o corpo da requisição adiante
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)
	})
}
