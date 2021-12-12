package route

import (
	http "net/http"

	v1 "github.com/GolangAPI/controller"
	"github.com/gorilla/mux"
)

type Route struct {
	AuthMiddleware int
	Name           string
	Method         string
	Pattern        string
	HandlerFunc    http.HandlerFunc
}
type Routes []Route

var internalRoute = Routes{
	// 0 --> default 1--admin 2--> user
	Route{0, "To Get list fom our syste", "GET", "/movies", v1.GetAllArticle},
}

func NewRoute() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range internalRoute {
		switch route.AuthMiddleware {
		case 1: // redirect this request to my middleware for checking the requst is admin or not
		default:
			router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
		}
	}
	return router
}
