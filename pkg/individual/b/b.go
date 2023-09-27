package b

import "net/http"

func HandleB(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("b"))
}
