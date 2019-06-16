package justgo

import (
	"net/http"
	"testing"
)

func TestAddRoute(t *testing.T) {
	AddRoute(http.MethodGet, "/", handler, middleware1, middleware2)
}

func middleware2(i http.Handler) http.Handler {
	return i
}

func middleware1(i http.Handler) http.Handler {
	return i
}

func handler(writer http.ResponseWriter, request *http.Request) {

}
