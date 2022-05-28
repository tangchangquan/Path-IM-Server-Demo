package test

import (
	"github.com/Path-IM/Path-IM-Server-Demo/app/im-user/cmd/api/internal/types"
	"testing"
)

func TestApplyFriendLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	post("/v1/relation/friend/apply", &types.ApplyFriendReq{
		UserId:  "21f5ca36a0c77ed38dabddafa7d39445",
		Message: "我就",
	}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("6553ce51d4310130cbe5a2c28764c200"),
	)
	t.Logf("%+v", resp)
}

func TestGetFriendApplyList(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	get("/v1/relation/friend/apply/list", map[string]string{
		"page":     "1",
		"pageSize": "10",
	}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}

func TestAgreeFriend(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	post("/v1/relation/friend/agree", &types.AgreeFriendReq{ApplyId: "51ee3c8a2e6b1d8ea523e1248e71c50f"}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}

func TestGetFriendListLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	get("/v1/relation/friend/list", map[string]string{}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}

func TestDeleteFriendLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	post("/v1/relation/friend/delete", &types.DeleteFriendReq{FriendId: "6553ce51d4310130cbe5a2c28764c200"}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}

func TestBlackUserLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	post("/v1/relation/black/add", &types.BlackUserReq{UserId: "6553ce51d4310130cbe5a2c28764c200"}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}

func TestUnBlackUserLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	post("/v1/relation/black/delete", &types.UnBlackUserReq{UserId: "6553ce51d4310130cbe5a2c28764c200"}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}

func TestGetBlackListLogic(t *testing.T) {
	resp := &types.IsUsernameExistResp{}
	get("/v1/relation/black/list", map[string]string{
		"page":     "1",
		"pageSize": "20",
	}, resp,
		withToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiI2NTUzY2U1MWQ0MzEwMTMwY2JlNWEyYzI4NzY0YzIwMCIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NzM4MzY0LCJuYmYiOjE2NTMxMzgzNjQsImlhdCI6MTY1MzEzODM2NH0.Jye6_oySFM9pDDuc7s3LczLTLuxJrJgbO5d9zTXbL_E"),
		withUid("21f5ca36a0c77ed38dabddafa7d39445"),
	)
	t.Logf("%+v", resp)
}
