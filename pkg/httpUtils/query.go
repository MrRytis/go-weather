package httpUtils

import (
	"net/http"
	"strconv"
)

func GetStringQueryParam(r *http.Request, name string, def string) string {
	p := r.URL.Query().Get(name)

	if p == "" {
		return def
	}

	return p
}

func GetIntQueryParam(r *http.Request, name string, def int) int {
	p := r.URL.Query().Get(name)

	if p == "" {
		return def
	}

	intValue, err := strconv.Atoi(p)
	if err != nil {
		return 0
	}

	return intValue
}

func GetFloatQueryParam(r *http.Request, name string, def float64) float64 {
	p := r.URL.Query().Get(name)

	if p == "" {
		return def
	}

	floatValue, err := strconv.ParseFloat(p, 64)
	if err != nil {
		return 0.0
	}

	return floatValue
}
