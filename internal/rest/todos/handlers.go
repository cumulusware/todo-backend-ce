package todos

import (
	"net/http"

	"github.com/cumulusware/todo-backend-ce/internal/rest"
)

// DescribeAll handles the OPTIONS method for the /todos/ endpoint.
func DescribeAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.RespondWithOptions(w, "GET,POST,DELETE,OPTIONS")
	}
}

// Describe handles the OPTIONS method for the /todos/{key}/ endpoint.
func Describe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.RespondWithOptions(w, "GET,PATCH,DELETE,OPTIONS")
	}
}
