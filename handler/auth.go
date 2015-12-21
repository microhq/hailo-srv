package handler

import (
	"github.com/micro/go-micro/errors"
	"github.com/micro/hailo-srv/hailo"
	auth "github.com/micro/hailo-srv/proto/auth"
	"golang.org/x/net/context"
)

type Auth struct{}

func (c *Auth) Test(ctx context.Context, req *auth.TestRequest, rsp *auth.TestResponse) error {
	status, err := hailo.AuthTest()
	if err != nil {
		return errors.InternalServerError("go.micro.srv.hailo.Auth.Test", err.Error())
	}
	rsp.Status = status
	return nil
}
