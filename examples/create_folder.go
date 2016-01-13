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
}
