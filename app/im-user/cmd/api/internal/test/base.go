package test

import (
	"github.com/go-resty/resty/v2"
	testUtils "github.com/showurl/Zero-IM-Server/common/utils/test"
)

const (
	//Host = "http://localhost:20010"
	Host = "http://42.194.149.177:10000"
)

type option struct {
	token    string
	platform string
	uid      string
}
type funcOptions func(o *option)

func withToken(token string) funcOptions {
	return func(o *option) {
		o.token = token
	}
}
func withUid(uid string) funcOptions {
	return func(o *option) {
		o.uid = uid
	}
}
func withPlatform(platform string) funcOptions {
	return func(o *option) {
		o.platform = platform
	}
}
func post(path string, data interface{}, resp interface{}, options ...funcOptions) {
	o := &option{
		token:    "",
		platform: "IOS",
	}
	for _, option := range options {
		option(o)
	}
	testUtils.ParseResponseJsonBody(func(r *resty.Request) (*resty.Response, error) {
		return r.
			SetHeader("Content-Type", "application/json").
			SetHeader("platform", o.platform).
			SetHeader("app-version", "v0.0.1").
			SetHeader("User-Agent", "go unit test").
			SetHeader("user_id", o.uid).
			SetHeader("token", o.token).
			SetBody(data).Post(Host + path)
	}, resp)
}
func get(path string, form map[string]string, resp interface{}, options ...funcOptions) {
	o := &option{
		token:    "",
		platform: "IOS",
	}
	for _, option := range options {
		option(o)
	}
	testUtils.ParseResponseJsonBody(func(r *resty.Request) (*resty.Response, error) {
		return r.
			SetHeader("Content-Type", "application/json").
			SetHeader("platform", o.platform).
			SetHeader("app-version", "v0.0.1").
			SetHeader("User-Agent", "go unit test").
			SetHeader("token", o.token).
			SetHeader("user_id", o.uid).
			SetQueryParams(form).
			Get(Host + path)
	}, resp)
}
func put(path string, data interface{}, resp interface{}, options ...funcOptions) {
	o := &option{
		token:    "",
		platform: "IOS",
	}
	for _, option := range options {
		option(o)
	}
	testUtils.ParseResponseJsonBody(func(r *resty.Request) (*resty.Response, error) {
		return r.
			SetHeader("Content-Type", "application/json").
			SetHeader("platform", o.platform).
			SetHeader("app-version", "v0.0.1").
			SetHeader("User-Agent", "go unit test").
			SetHeader("token", o.token).
			SetHeader("user_id", o.uid).
			SetBody(data).Put(Host + path)
	}, resp)
}
func del(path string, data interface{}, resp interface{}, options ...funcOptions) {
	o := &option{
		token:    "",
		platform: "IOS",
	}
	for _, option := range options {
		option(o)
	}
	testUtils.ParseResponseJsonBody(func(r *resty.Request) (*resty.Response, error) {
		return r.
			SetHeader("Content-Type", "application/json").
			SetHeader("platform", o.platform).
			SetHeader("app-version", "v0.0.1").
			SetHeader("User-Agent", "go unit test").
			SetHeader("token", o.token).
			SetHeader("user_id", o.uid).
			SetBody(data).Delete(Host + path)
	}, resp)
}
