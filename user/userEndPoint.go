package user

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	//封装User请求结构体
	Uid int `json:"uid"`
}

type UserResponse struct {
	Resule string `json:"resule"`
}

func  GetUserEnPoint(usersevicons IuserService)endpoint.Endpoint {
     return func(ctx context.Context, request interface{}) (response interface{}, err error) {
     	fmt.Println("我调用了Servicon层")
		 r := request.(UserRequest)
		result := usersevicons.GetName(r.Uid)
		return UserResponse{Resule:result},nil
	 }
}