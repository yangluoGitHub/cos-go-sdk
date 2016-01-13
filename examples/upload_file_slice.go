package main

import (
	"fmt"
	"github.com/yangluoGitHub/cos-go-sdk"
)

func main() {

	var appId uint = 10016057
	secretId := "AKIDoUqGP9jLfYgqVgueaT0CneVWDha9tdUM"
	secretKey := "kmabDFmmQkV9PGThSmfw8TVtrkmvrNHl"
	bucketName := "bucket1"

	//new bucket Object
	bucket, _ := coscloud.New(appId, secretId, secretKey, bucketName)
	//用户指定分片大小来分片上传
	res, err := bucket.Upload_slice("../test/data.bin", "/cos-go-sdk/upload_slice/data.bin", "upload_slice test", 3*1024*1024, "")
	//上传失败，重新上传，不论是否指定session，都可以实现断点续传
	// res, err := bucket.Upload_slice("../test/data.bin", "/cos-go-sdk/upload_slice/data.bin", "upload_slice test", 3*1024*1024, "48d44422-3188-4c6c-b122-6f780742f125+CpzDLtEHAA==")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nUrl:", res.Data.Url,
		"\nResourcePath:", res.Data.ResourcePath,
		"\nAccess Url:", res.Data.AccessUrl)
}
