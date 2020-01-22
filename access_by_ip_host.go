package main

import (
	"context"
	"net/url"
	"os"
	"strings"
	"net/http"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func main() {
    // 通过 IP 和 PORT 访问 COS
    u, _ := url.Parse("http://ip:port")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
            Transport: &debug.DebugRequestTransport{
                RequestHeader: true,
                RequestBody:    false,
                ResponseHeader: true,
                ResponseBody:   true,
            },
		},
	})
    // 设置 Host 信息
    c.Host = "test-1259654469.cos.ap-guangzhou.myqcloud.com"

    name := "example"
    f := strings.NewReader("test")
    _, err := c.Object.Put(context.Background(), name, f, nil)
    if err != nil {
        panic(err)
    }
}
