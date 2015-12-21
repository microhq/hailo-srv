package hailo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/micro/go-micro/errors"
)

var (
	Token string
	Url   = "https://api.hailoapp.com/"
)

func Respond(path string, args url.Values, rsp interface{}) error {
	u := Url + path
	req, err := http.NewRequest("GET", u+"?"+args.Encode(), nil)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.hailo", err.Error())
	}
	if len(Token) > 0 {
		req.Header.Set("Authorization", "token "+Token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.hailo", err.Error())
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		if err := json.Unmarshal(b, rsp); err != nil {
			return errors.InternalServerError("go.micro.srv.hailo", err.Error())
		}
		return nil
	} else if resp.StatusCode == 403 {
		return errors.Forbidden("go.micro.srv.hailo", resp.Status)
	} else {
		return errors.InternalServerError("go.micro.srv.hailo", resp.Status)
	}

	return nil
}

func AuthTest() (string, error) {
	args := url.Values{
		"latitude":  {"51.510761"},
		"longitude": {"-0.1174437"},
	}
	u := Url + "/drivers/near"
	req, err := http.NewRequest("GET", u+"?"+args.Encode(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "token "+Token)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()
	return http.StatusText(rsp.StatusCode), nil
}

func ApiStatus() (string, error) {
	u := Url + "/status/up"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "token "+Token)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()
	return http.StatusText(rsp.StatusCode), nil
}
