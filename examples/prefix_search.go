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

	res, err := bucket.PrefixSearch("/cos-go-sdk/", 20, coscloud.ELISTBOTH, 0, "")
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

}
