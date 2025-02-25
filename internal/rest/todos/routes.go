package todos

import (
	"net/http"
)

// AddRoutes adds the todos related sub-routes.
func AddRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.Handle("OPTIONS /", DescribeAll())
	r.Handle("OPTIONS /{key}", Describe())
	r.Handle("GET /", ReadAll())
	r.Handle("POST /", Create())
	r.Handle("DELETE /", DeleteAll())
	return r
}
