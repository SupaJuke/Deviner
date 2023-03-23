package guess

import (
	"net/http"
)

func Guess(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("uwu"))
	if err != nil {
		return
	}
	// TODO: implement this
}
