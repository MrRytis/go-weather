package httpUtils

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetStringVariable(r *http.Request, name string) string {
	vars := mux.Vars(r)

	return vars[name]
}

func GetIntVariable(r *http.Request, name string) int {
	vars := mux.Vars(r)

	intValue, err := strconv.Atoi(vars[name])
	if err != nil {
		return 0
	}

	return intValue
}
