package test

import (
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/types"
	"testing"
)

func TestRegister(t *testing.T) {
	resp := &types.RegisterResp{}
	post("/v1/white/account/register", &types.RegisterReq{
		Username: "user17",
		Password: "123456",
	}, resp)
	t.Log(resp.Token)
	t.Logf("%+v", resp.UserModel)
}

func TestIsUsernameExistLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	get("/v1/white/account/exist", map[string]string{
		"username": "showurl3",
	}, resp)
	t.Logf("%+v", resp)
}

func TestLogin(t *testing.T) {
	resp := &types.LoginResp{}
	post("/v1/white/account/login", &types.LoginReq{
		Username: "showurl3",
		Password: "123456",
	}, resp)
	t.Logf("%+v", resp)
}
func TestGetSelfInfoLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	get("/v1/user/selfinfo", map[string]string{}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}

func TestUpdateSelfInfoLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	put("/v1/user/selfinfo", &types.ModifySelfInfoReq{
		Nickname: "1改昵称咯",
		Sign:     "2改签名咯",
		Avatar:   "",
		Province: "北京市",
		City:     "北京市",
		District: "通州区",
		Birthday: "2000-01-01",
		IsMale:   false,
	}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}
