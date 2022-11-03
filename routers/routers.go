package routers

import (
	"fmt"
	"golang-redis/controller"
	"net/http"

	"github.com/gorilla/mux"
)

// router modelo de rota
type router struct {
	URI      string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
}

// userRouters variavel que contem um array do modelo de rotas
var userRouters = []router{
	{
		URI:      "/session",
		Method:   http.MethodGet,
		Function: controller.Session,
	},
	{
		URI:      "/requests",
		Method:   http.MethodPost,
		Function: controller.Requests,
	},
	{
		URI:      "/timeleft",
		Method:   http.MethodPost,
		Function: controller.TimeLeft,
	},
}

// R variavel que carrega o objeto de que lida com as rotas
var R = mux.NewRouter()

// HandleRouters itera sobre as o slice de rotas, jogando o usuario para dentro da rota solicitada
func HandleRouters() {
	for _, route := range userRouters {
		fmt.Println(route)
		R.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

}
