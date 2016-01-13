[![Build Status](https://drone.io/github.com/yangluoGitHub/cos-go-sdk/status.png)](https://drone.io/github.com/yangluoGitHub/cos-go-sdk/latest)
[![Coverage Status](https://coveralls.io/repos/yangluoGitHub/cos-go-sdk/badge.svg?branch=master&service=github)](https://coveralls.io/github/yangluoGitHub/cos-go-sdk?branch=master)
[![GoDoc](https://godoc.org/github.com/yangluoGitHub/cos-go-sdk?status.svg)](https://godoc.org/github.com/yangluoGitHub/cos-go-sdk)
##cos-go-sdk 简介
- cos-go-sdk 是基于[腾讯云对象存储服务 COS](http://www.qcloud.com/product/cos.html) 官方 Restful API 构建。
- 各个接口均提供了单独的文档说明，可以参考 docs 目录下的项目文档。
- 可以参考 examples 目录下的示例。

##环境
- cos-go-sdk 推荐使用 Go 1.2 及以上 Go 语言版本。
- Windows，Linux，Mac OS X

##安装
```bash
go get github.com/yangluoGitHub/cos-go-sdk
```

##API 文档
[![GoDoc](https://godoc.org/github.com/yangluoGitHub/cos-go-sdk?status.svg)](https://godoc.org/github.com/yangluoGitHub/cos-go-sdk)

##快速入门


### 创建目录
- 
```go 
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.CreateFolder("/cos-go-sdk/createFolder/test/", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nCtime:", res.Data.Ctime,
		"\nResource Path:", res.Data.ResourcePath)
```


### 文件上传

- 
```go
client := cos.NewClient(appId, secretId, secretKey)
res, err := client.UploadFile("cosdemo", "/hello/hello.txt", "/users/new.txt", "file attr")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Code:", res.Code,
    "\nMessage:", res.Message,
    "\nUrl:", res.Data.Url,
    "\nResourcePath:", res.Data.ResourcePath,
    "\nAccess Url:", res.Data.AccessUrl)
```


##完整示例

更多示例请查看 examples 目录

##项目文档

更多文档请查看 docs 目录
