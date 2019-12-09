package todo

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/todo/foo", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("todoID", "foo")

	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	GetTodo(w, r)

	body, err := w.Body.ReadString(0)
	body = strings.TrimSpace(body)

	if err != nil && err.Error() != "EOF" {
		t.Error(err)
	}

	if body != "{\"slug\":\"foo\",\"title\":\"Hello world\",\"body\":\"Heloo world from planet earth\"}" {
		t.Error(body)
	}
}
