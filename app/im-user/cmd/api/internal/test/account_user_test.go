package test

import (
	"fmt"
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/api/internal/types"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	resp := &types.RegisterResp{}
	post("/v1/white/account/register", &types.RegisterReq{
		Username: "user52",
		Password: "123456",
	}, resp)
	t.Log(resp.Token)
	t.Logf("%+v", resp.UserModel)
}
func TestRegisters(t *testing.T) {
	for i := 0; i < 1000; i++ {
		post("/v1/white/account/register", &types.RegisterReq{
			Username: fmt.Sprintf("trs-1-%d", i),
			Password: "123456",
		}, nil)
	}
}

func BenchmarkRegister(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp := &types.RegisterResp{}
		post("/v1/white/account/register", &types.RegisterReq{
			Username: fmt.Sprintf("bm-1-%d", i),
			Password: "123456",
		}, resp)
	}
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
func TestGetSelfInfos(t *testing.T) {
	for i := 0; i < 20; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				get("/v1/user/selfinfo", map[string]string{}, nil,
					withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
					withUid("21f5ca36a0c77ed38dabddafa7d39445"),
				)
			}
		}()
	}
	time.Sleep(time.Second * 20)
}

func TestGetUserInfoByIdLogic(t *testing.T) {
	resp := &types.GetUserByIdResp{}
	get("/v1/user/getinfobyid", map[string]string{
		"id": "26338dd192cccc5ec7c656bdd588ef67",
		//"Id": "",
	}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}
func TestGetUserInfoByUsernameLogic(t *testing.T) {
	resp := &types.GetUserByIdResp{}
	get("/v1/user/getinfobyusername", map[string]string{
		"username": "21f5ca36a0c77ed38dabddafa7d39445",
		//"Id": "",
	}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}

func TestUpdateSelfInfoLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	put("/v1/user/selfinfo", &types.ModifySelfInfoReq{
		Nickname: "1????????????",
		Sign:     "2????????????",
		Avatar:   "",
		Province: "?????????",
		City:     "?????????",
		District: "?????????",
		Birthday: "2000-01-01",
		IsMale:   false,
	}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}
