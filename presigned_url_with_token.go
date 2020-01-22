package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type URLToken struct {
	SessionToken string `url:"x-cos-security-token,omitempty" header:"-"`
}

func main() {
    // 替换成您的临时密钥
	tak := "<tmp_secretid>"
	tsk := "<tmp_secretkey>"
	token := &URLToken{
		SessionToken: "<token>",
	}
	u, _ := url.Parse("https://test-1259654469.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{})

	name := "exampleobject"
	ctx := context.Background()

	// Get presigned
	presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodGet, name, tak, tsk, time.Hour, token)
	if err != nil {
		panic(err)
	}
	// Get object by presinged url
	_, err = http.Get(presignedURL.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(presignedURL.String())
}
