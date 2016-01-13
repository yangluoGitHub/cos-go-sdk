[![Build Status](https://drone.io/github.com/yangluoGitHub/cos-go-sdk/status.png)](https://drone.io/github.com/yangluoGitHub/cos-go-sdk/latest)
[![Coverage Status](https://coveralls.io/repos/yangluoGitHub/cos-go-sdk/badge.svg?branch=master&service=github)](https://coveralls.io/github/yangluoGitHub/cos-go-sdk?branch=master)
[![GoDoc](https://godoc.org/github.com/yangluoGitHub/cos-go-sdk?status.svg)](https://godoc.org/github.com/yangluoGitHub/cos-go-sdk)
##cos-go-sdk 简介
- cos-go-sdk 是[腾讯云对象存储服务 COS](http://www.qcloud.com/product/cos.html) go语言版本SDK

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

### 创建文件（完整上传）

```go
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.Upload("../test/test.jpg", "/cos-go-sdk/upload/test.jpg", "upload test")
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

### 创建文件（分片上传）

```go
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.Upload_slice("../test/data.bin", "/cos-go-sdk/upload_slice/data.bin", "upload_slice test", 3*1024*1024, "")
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


### 目录列表

```go 
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.ListFolder("/cos-go-sdk/createFolder/test/", 20, coscloud.ELISTBOTH, coscloud.Asc, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nContext:", res.Data.Context,
		"\nHasMore:", res.Data.HasMore,
		"\nDirCount:", res.Data.DirCount,
		"\nFileCount:", res.Data.FileCount,
	)
	fmt.Println("=================================")
	for _, info := range res.Data.Infos {
		fmt.Println("Name:", info.Name,
			"\nBizAttr:", info.BizAttr,
			"\nFileSize:", info.FileSize,
			"\nFileLen:", info.FileLen,
			"\nSha:", info.Sha,
			"\nCtime:", info.Ctime,
			"\nMtime:", info.Mtime,
			"\nAccess URL:", info.AccessUrl,
		)
		fmt.Println("=================================")
	}
```

### 前缀搜索

```go 
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.PrefixSearch("/cos-go-sdk", 20, coscloud.ELISTBOTH, coscloud.Asc, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nContext:", res.Data.Context,
		"\nHasMore:", res.Data.HasMore,
		"\nDirCount:", res.Data.DirCount,
		"\nFileCount:", res.Data.FileCount,
	)
	fmt.Println("=================================")
	for _, info := range res.Data.Infos {
		fmt.Println("Name:", info.Name,
			"\nBizAttr:", info.BizAttr,
			"\nFileSize:", info.FileSize,
			"\nFileLen:", info.FileLen,
			"\nSha:", info.Sha,
			"\nCtime:", info.Ctime,
			"\nMtime:", info.Mtime,
			"\nAccess URL:", info.AccessUrl,
		)
		fmt.Println("=================================")
	}
```

### 目录信息更新

```go 
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.UpdateFolder("/cos-go-sdk/createFolder/test/", "update-attr")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code, "\nMessage:", res.Message)
```

### 文件信息更新

```go 
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.Update("/cos-go-sdk/upload/test.jpg", "update-attr")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code, "\nMessage:", res.Message)
```

### 目录信息查询

```go 
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.StatFolder("/cos-go-sdk/createFolder/test/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nName:", res.Data.Name,
		"\nBizAttr:", res.Data.BizAttr,
		"\nCtime:", res.Data.Ctime,
		"\nMtime:", res.Data.Mtime)
```

### 文件信息查询

```go 
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.Stat("/cos-go-sdk/upload/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nName:", res.Data.Name,
		"\nBizAttr:", res.Data.BizAttr,
		"\nFileSize:", res.Data.FileSize,
		"\nFileLen:", res.Data.FileLen,
		"\nSha:", res.Data.Sha,
		"\nCtime:", res.Data.Ctime,
		"\nMtime:", res.Data.Mtime,
		"\nAccess Url:", res.Data.AccessUrl)
```

### 删除目录

```go 
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.DelFolder("/cos-go-sdk/createFolder/test/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code, "\nMessage:", res.Message)
```

### 删除文件

```go 
bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	res, err := bucket.Del("/cos-go-sdk/upload/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code, "\nMessage:", res.Message)
```

##完整示例

更多示例请查看 examples 目录

##项目文档

更多文档请查看 docs 目录
