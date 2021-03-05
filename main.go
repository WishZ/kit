package main

import (
	"kit/services"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	user := services.UserService{}
	endp := services.GenUserEndpoint(user)

	serverHander := httptransport.NewServer(endp, services.DecodeUserRequest, services.EncodeUserResponse)
	//路由
	r := mux.NewRouter()
	r.Methods("GET").Path(`/user/{uid:\d+}`).Handler(serverHander)
	//http.Lis
	http.ListenAndServe(":9001", r)
}
