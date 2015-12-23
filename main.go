package main

import (
	"github.com/codegangsta/cli"
	log "github.com/golang/glog"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"

	"github.com/micro/hailo-srv/hailo"
	"github.com/micro/hailo-srv/handler"
	api "github.com/micro/hailo-srv/proto/api"
	auth "github.com/micro/hailo-srv/proto/auth"
	drivers "github.com/micro/hailo-srv/proto/drivers"
)

func main() {
	cmd.Flags = append(cmd.Flags,
		cli.StringFlag{
			Name:   "api_token",
			EnvVar: "API_TOKEN",
			Usage:  "Hailo API token",
		},
	)

	cmd.Actions = append(cmd.Actions, func(c *cli.Context) {
		hailo.Token = c.String("api_token")
	})

	cmd.Init()

	server.Init(
		server.Name("go.micro.srv.hailo"),
	)

	api.RegisterApiHandler(server.DefaultServer, &handler.Api{})
	auth.RegisterAuthHandler(server.DefaultServer, &handler.Auth{})
	drivers.RegisterDriversHandler(server.DefaultServer, &handler.Drivers{})

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
