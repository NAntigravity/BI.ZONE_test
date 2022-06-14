package utils

import "net/http"

func SetTotalCountHeader(w http.ResponseWriter, count string) {
	w.Header().Add("Access-Control-Expose-Headers", "X-Total-Count")
	w.Header().Add("X-Total-Count", count)
}
