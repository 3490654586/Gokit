package main

import (
	user2 "demo/user"
	"fmt"
	myhttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)





func main() {
	userss := &user2.User{}
	fmt.Println(userss)
	endp := user2.GetUserEnPoint(userss)
	fmt.Println(endp)

	serverHandel := myhttp.NewServer(endp, user2.DecodeUserRequest, user2.EncodeUserResponse) //使用go kit创建server传入我们之前定义的两个解析函数

	r := mux.NewRouter()
	r.Methods("GET").Path(`/user/{uid:\d+}`).Handler(serverHandel)


	http.ListenAndServe(":8050",r)
}
