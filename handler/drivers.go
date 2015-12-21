package handler

import (
	"fmt"
	"net/url"

	"github.com/micro/hailo-srv/hailo"
	drivers "github.com/micro/hailo-srv/proto/drivers"
	"golang.org/x/net/context"
)

type Drivers struct{}

func (c *Drivers) Near(ctx context.Context, req *drivers.NearRequest, rsp *drivers.NearResponse) error {
	u := url.Values{}
	u.Set("latitude", fmt.Sprintf("%.6f", req.Latitude))
	u.Set("longitude", fmt.Sprintf("%.6f", req.Longitude))
	return hailo.Respond("/drivers/near", u, rsp)
}

func (c *Drivers) Eta(ctx context.Context, req *drivers.EtaRequest, rsp *drivers.EtaResponse) error {
	u := url.Values{}
	u.Set("latitude", fmt.Sprintf("%.6f", req.Latitude))
	u.Set("longitude", fmt.Sprintf("%.6f", req.Longitude))
	return hailo.Respond("/drivers/eta", u, rsp)
}
