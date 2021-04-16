package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("我是T包")
	//这个函数决定了使用哪个request结构体来请求
	vars := mux.Vars(r)
	if uid, ok := vars["uid"]; ok { //
		uid, _ := strconv.Atoi(uid)
		return UserRequest{Uid: uid}, nil
	}
	return nil,errors.New("参数错误")
}

func EncodeUserResponse(ctx context.Context,w http.ResponseWriter,response interface{}) error{
	fmt.Println("我是T2包")
	w.Header().Set("Content-type","application/json") //设置响应格式为json，这样客户端接收到的值就是json，就是把我们设置的UserResponse给json化了

	return json.NewEncoder(w).Encode(response)//判断响应格式是否正确
}