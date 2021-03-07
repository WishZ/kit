package main

import (
	"kit/services"
	"kit/util"
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
	{
		r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHander)
		r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-type", "application/json")
			writer.Write([]byte(`{"status":"ok"}`))
		})
	}
	//http.Lis
	util.RegService() //注册服务
	http.ListenAndServe(":10125", r)
}
