package a

import "net/http"

func HandleA(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("a!!"))
}
