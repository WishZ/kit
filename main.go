package main

import (
	"fmt"
	"kit/services"
	"kit/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	errChan := make(chan error)
	go (func() {
		util.RegService() //注册服务
		err := http.ListenAndServe(":10125", r)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	})()

	go (func() {
		signChan := make(chan os.Signal)
		signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-signChan)
	})()
	getErr := <-errChan
	util.UnRegService()
	log.Println(getErr)
}
