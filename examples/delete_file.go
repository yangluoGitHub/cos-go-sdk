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
	res, err := bucket.Del("/cos-go-sdk/upload/test.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code, "\nMessage:", res.Message)
}
