package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"kit/Services"
	"net/http"
)

func main() {
	user := Services.UserService{}
	endp := Services.GenUserEndpoint(user)

	serverHander := httptransport.NewServer(endp, Services.DecodeUserRequest, Services.EncodeUserResponse)

	//http.Lis
	http.ListenAndServe(":9001", serverHander)
}
