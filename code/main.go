package main

import (
	"context"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
	fc.StartHttp(HandleHttpRequest)
}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	/*
		在底下写业务逻辑。
		记得删除空返回值,和status200。
	*/

	return nil
}
