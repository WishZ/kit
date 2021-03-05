package main

import (
	"kit/Services"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	user := Services.UserService{}
	endp := Services.GenUserEndpoint(user)

	serverHander := httptransport.NewServer(endp, Services.DecodeUserRequest, Services.EncodeUserResponse)
	//路由
	r := mux.NewRouter()
	r.Methods("GET").Path(`/user/{uid:\d+}`).Handler(serverHander)
	//http.Lis
	http.ListenAndServe(":9001", r)
}
