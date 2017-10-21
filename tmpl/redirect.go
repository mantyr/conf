package tmpl

import (
	"net/http"
)

func Redirect(w http.ResponseWriter, address string) {
	if address == "" {
		address = "/"
	}
	w.Header().Set("Location", address)
	w.Header().Set("Cache-Control", "private, no-store, max-age=0, must-revalidate")
	w.WriteHeader(303) // see https://ru.wikipedia.org/wiki/Список_кодов_состояния_HTTP#303
}
