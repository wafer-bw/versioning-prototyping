package d

import "net/http"

func HandleD(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("d"))
}
