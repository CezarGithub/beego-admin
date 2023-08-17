package toolbar

import (
	t "quince/internal/toolbar/ajax"
//	c "quince/internal/toolbar/component"
	h "quince/internal/toolbar/html"
	m "quince/internal/toolbar/modal"
)

// XHttp-Request : Response type -> /global/response.go
func Ajax(name string) *t.Ajax {
	return &t.Ajax{Name: name}
}

// HttpRequest
func Html(name string) *h.Html {
	return &h.Html{Name: name}
}

// XHttp-Request : open modal popup
func Modal(name string) *m.Modal {
	return &m.Modal{Name: name}
}

