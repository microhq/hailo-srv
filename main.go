package main

import (
	"log"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"

	"github.com/micro/hailo-srv/hailo"
	"github.com/micro/hailo-srv/handler"
	api "github.com/micro/hailo-srv/proto/api"
	auth "github.com/micro/hailo-srv/proto/auth"
	drivers "github.com/micro/hailo-srv/proto/drivers"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.hailo"),
		micro.Flags(
			cli.StringFlag{
				Name:   "api_token",
				EnvVar: "API_TOKEN",
				Usage:  "Hailo API token",
			},
		),
		micro.Action(func(c *cli.Context) {
			hailo.Token = c.String("api_token")
		}),
	)

	service.Init()

	api.RegisterApiHandler(service.Server(), &handler.Api{})
	auth.RegisterAuthHandler(service.Server(), &handler.Auth{})
	drivers.RegisterDriversHandler(service.Server(), &handler.Drivers{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
