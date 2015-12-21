package handler

import (
	"github.com/micro/go-micro/errors"
	"github.com/micro/hailo-srv/hailo"
	api "github.com/micro/hailo-srv/proto/api"
	"golang.org/x/net/context"
)

type Api struct{}

func (c *Api) Status(ctx context.Context, req *api.StatusRequest, rsp *api.StatusResponse) error {
	status, err := hailo.ApiStatus()
	if err != nil {
		return errors.InternalServerError("go.micro.srv.hailo.Api.Status", err.Error())
	}
	rsp.Status = status
	return nil
}
