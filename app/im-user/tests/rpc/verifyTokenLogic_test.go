package rpc

import (
	"context"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	"testing"
)

func TestVerifyTokenLogic(t *testing.T) {
	resp, err := imUserService.VerifyToken(context.Background(), &pb.VerifyTokenReq{
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiIyZTVmNDAxZjljZmFiNmIxNDc1YzZlYmNkOTAxZTg1MiIsIlBsYXRmb3JtIjoiSU9TIiwiZXhwIjo0ODA2NTYwNjM0LCJuYmYiOjE2NTI5NjA2MzQsImlhdCI6MTY1Mjk2MDYzNH0.VXMJDhY8kexcDQ0m1FwsXZqy9Kez-zK4imTryPLtigU",
		Platform: "IOS",
	})

	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}
