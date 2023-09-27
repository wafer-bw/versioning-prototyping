package c

import "net/http"

func HandleC(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("c"))
}
