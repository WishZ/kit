package util

import (
	"github.com/hashicorp/consul/api"
	"log"
)

func RegService() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"

	reg := api.AgentServiceRegistration{}
	reg.ID = "userservice"
	reg.Name = "userservice"
	reg.Address = "192.168.50.93"
	reg.Port = 10125
	reg.Tags = []string{"primary"}

	check := api.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://192.168.50.93:10125/health"

	reg.Check = &check

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	client.Agent().ServiceRegister(&reg)

}
