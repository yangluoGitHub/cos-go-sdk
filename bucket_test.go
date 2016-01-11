package coscloud

import (
	"fmt"
	// "math/rand"
	"testing"
	"time"
)

const (
	Appid      = 10016057
	SecretId   = "AKIDoUqGP9jLfYgqVgueaT0CneVWDha9tdUM"
	SecretKey  = "kmabDFmmQkV9PGThSmfw8TVtrkmvrNHl"
	BucketName = "bucket1"
)

var bucket = func() *Bucket {
	buc, _ := New(Appid, SecretId, SecretKey, BucketName)
	return buc
}()

var folderPath = func() string {
	now := GetNowSec()
	return fmt.Sprintf("%s_%d", "/cos-go-sdk/createFolder/folder", now)
}()

var srcPath string = "./test/pic/test.jpg"
var dstPath = func() string {
	now := GetNowSec()
	return fmt.Sprintf("%s_%d.jpg", "/cos-go-sdk/Upload/test", now)
}()

var slice_srcPath string = "/Users/renxiaoqing/Downloads/apache-activemq-5.12.1-bin.tar.gz"
var slice_dstPath string = "/cos-go-sdk/upload_slice/test001_apache-activemq-5.12.1-bin.tar.gz"

// var slice_srcPath string = "/Users/renxiaoqing/Downloads/9.18OK.7z"
// var slice_dstPath string = "/cos-go-sdk/upload_slice/test001_9.18OK.7z"

func TestCreateFolder(t *testing.T) {
	_, err := bucket.CreateFolder(folderPath, "")
	if err != nil {
		t.Errorf("cos createFolder path=%s, failed, err=%s\n", folderPath, err.Error())
	} else {
		fmt.Printf("cos createFolder path=%s success\n", folderPath)
	}

}

func TestUpdateFolder(t *testing.T) {

	bizAttr := "update"
	_, err := bucket.UpdateFolder(folderPath, bizAttr)
	if err != nil {
		t.Errorf("cos updateFolder path=%s, failed, err=%s\n", folderPath, err.Error())
	} else {
		fmt.Printf("cos updateFolder path=%s success\n", folderPath)
	}
}

func TestStatFolder(t *testing.T) {
	json, err := bucket.StatFolder(folderPath)
	if err != nil {
		t.Errorf("cos statFolder path=%s, failed, err=%s\n", folderPath, err.Error())
	} else {
		fmt.Printf("cos statFolder path=%s success\n", folderPath)
		fmt.Println(json)
	}
}

func TestListFolder(t *testing.T) {
	json, err := bucket.ListFolder(folderPath, 20, "eListBoth", 0, "")
	if err != nil {
		t.Errorf("cos listFolder path=%s, failed, err=%s\n", folderPath, err.Error())
	} else {
		fmt.Println("cos listFolder path=%s success", folderPath)
		fmt.Println(json)
	}

}

func TestPrefixSearch(t *testing.T) {
	json, err := bucket.PrefixSearch(folderPath, 20, "eListBoth", 0, "")
	if err != nil {
		t.Errorf("cos prefixSearch path=%s, failed, err=%s\n", folderPath, err.Error())
	} else {
		fmt.Println("cos prefixSearch path=%s success", folderPath)
		fmt.Println(json)
	}

}

func TestDelFolder(t *testing.T) {
	json, err := bucket.DelFolder(folderPath)
	if err != nil {
		t.Errorf("cos delFolder path=%s, failed, err=%s\n", folderPath, err.Error())
	} else {
		fmt.Println("cos delFolder path=%s success", folderPath)
	}
	fmt.Println(json)
}

func TestUpload(t *testing.T) {

	_, err := bucket.Upload(srcPath, dstPath, "")
	if err != nil {
		t.Errorf("cos upload failed, srcPath=%s, err=%s\n", srcPath, err.Error())
	} else {
		fmt.Printf("cos upload srcPath=%s, dstPath=%s success\n", srcPath, dstPath)
	}
}

func TestStat(t *testing.T) {
	json, err := bucket.Stat(dstPath)
	if err != nil {
		t.Errorf("cos stat file=%s, failed, err=%s\n", dstPath, err.Error())
	} else {
		fmt.Printf("cos stat file=%s success\n", dstPath)
		fmt.Println(json)
	}
}

func TestUpdate(t *testing.T) {
	_, err := bucket.Update(dstPath, "update")
	if err != nil {
		t.Errorf("cos update file=%s, failed, err=%s\n", dstPath, err.Error())
	} else {
		fmt.Printf("cos update file=%s success\n", dstPath)
	}
}

func TestDel(t *testing.T) {
	json, err := bucket.Del(dstPath)
	if err != nil {
		t.Errorf("cos del file=%s, failed, err=%s\n", dstPath, err.Error())
	} else {
		fmt.Println("cos del file=%s success", dstPath)
	}
	fmt.Println(json)
}

func TestUpload_slice(t *testing.T) {

	json, err := bucket.Upload_slice(slice_srcPath, slice_dstPath, "", 0, "")
	if err != nil {
		t.Errorf("cos upload_slice file=%s, failed, err = %s\n", slice_srcPath, err.Error())
	} else {
		fmt.Printf("cos Upload_slice file=%s dstPath=%s success", slice_srcPath, slice_dstPath)
	}
	fmt.Println(json)
}

func BenchmarkUpload(b *testing.B) {
	now := GetNowSec()
	for i := 0; i < 10; i++ {
		dstPath := fmt.Sprintf("%s_%d_%d.jpg", "/cos-go-sdk/test", now, i)
		_, _ = bucket.Upload(srcPath, dstPath, "")
	}
}

// GetNowSec returns Unix time, the number of seconds elapsed since January 1, 1970 UTC.
// 获取当前时间，从UTC开始的秒数。
func GetNowSec() int64 {
	return time.Now().Unix()
}
