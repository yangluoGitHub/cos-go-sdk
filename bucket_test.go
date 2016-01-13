package coscloud

import (
	"fmt"
	// "math/rand"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type BucketSuite struct {
	bucket *Bucket
}

var _ = Suite(&BucketSuite{})

const (
	Appid      = 10016057
	SecretId   = "AKIDoUqGP9jLfYgqVgueaT0CneVWDha9tdUM"
	SecretKey  = "kmabDFmmQkV9PGThSmfw8TVtrkmvrNHl"
	BucketName = "bucket1"
)

// Run once when the suite starts running
func (s *BucketSuite) SetUpSuite(c *C) {
	bucket, err := New(Appid, SecretId, SecretKey, BucketName)
	c.Assert(err, IsNil)
	s.bucket = bucket

	fmt.Println("SetUpSuite")
}

// Run once after all tests or benchmarks have finished running.
func (s *BucketSuite) TearDownSuite(c *C) {
	// Delete Objects
	// lor, err := s.bucket.ListObjects()
	// c.Assert(err, IsNil)

	// for _, object := range lor.Objects {
	// 	err = s.bucket.DeleteObject(object.Key)
	// 	c.Assert(err, IsNil)
	// }

	fmt.Println("TearDownSuite")
}

// Run before each test or benchmark starts
func (s *BucketSuite) SetUpTest(c *C) {
}

// Run after each test or benchmark runs.
func (s *BucketSuite) TearDownTest(c *C) {
}

var folderPath = func() string {
	now := GetNowSec()
	return fmt.Sprintf("%s_%d", "/cos-go-sdk/createFolder/folder", now)
}()

var srcPath string = "./test/test.jpg"
var dstPath = func() string {
	now := GetNowSec()
	return fmt.Sprintf("%s_%d.jpg", "/cos-go-sdk/Upload/test", now)
}()

var slice_srcPath string = "./test/data.bin"
var slice_dstPath string = "/cos-go-sdk/upload_slice/data.bin"

func (s *BucketSuite) TestCreateAndDelFolder(c *C) {

	res, err := s.bucket.CreateFolder(folderPath, "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	resDel, err := s.bucket.DelFolder(folderPath)
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)
}

func (s *BucketSuite) TestUpdateAndStatFolder(c *C) {

	res, err := s.bucket.CreateFolder(folderPath, "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	resUpdate, err := s.bucket.UpdateFolder(folderPath, "update")
	c.Assert(err, IsNil)
	c.Assert(resUpdate.Code, Equals, 0)

	resStat, err := s.bucket.StatFolder(folderPath)
	c.Assert(err, IsNil)
	c.Assert(resStat.Code, Equals, 0)
	c.Assert(resStat.Data.BizAttr, Equals, "update")

	resStat, err = s.bucket.StatFolder("")
	c.Assert(err, IsNil)
	c.Assert(resStat.Code, Equals, 0)

	resDel, err := s.bucket.DelFolder(folderPath)
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)
}

func (s *BucketSuite) TestListFolder(c *C) {
	res, err := s.bucket.CreateFolder(folderPath, "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	resList, err := s.bucket.ListFolder(folderPath, 20, ELISTBOTH, Asc, "")
	c.Assert(err, IsNil)
	c.Assert(resList.Code, Equals, 0)

	resList, err = s.bucket.ListFolder("/", 20, ELISTBOTH, Asc, "")
	c.Assert(err, IsNil)
	c.Assert(resList.Code, Equals, 0)

	resList, err = s.bucket.ListFolder(folderPath, 20, ELISTDIRONLY, Asc, "")
	c.Assert(err, IsNil)
	c.Assert(resList.Code, Equals, 0)

	resList, err = s.bucket.ListFolder(folderPath, 20, ELISTDIRONLY, Desc, "")
	c.Assert(err, IsNil)
	c.Assert(resList.Code, Equals, 0)

	resList, err = s.bucket.ListFolder(folderPath, 20, ELISTFILEONLY, Asc, "")
	c.Assert(err, IsNil)
	c.Assert(resList.Code, Equals, 0)

	resList, err = s.bucket.ListFolder(folderPath, 20, ELISTFILEONLY, Desc, "")
	c.Assert(err, IsNil)
	c.Assert(resList.Code, Equals, 0)

	resSearch, err := s.bucket.PrefixSearch("/cos-go-sdk/u", 20, ELISTBOTH, Asc, "")
	c.Assert(err, IsNil)
	c.Assert(resSearch.Code, Equals, 0)

	resDel, err := s.bucket.DelFolder(folderPath)
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)
}

func (s *BucketSuite) TestUploadFile(c *C) {
	resUpload, err := s.bucket.Upload(srcPath, dstPath, "go testcase for cos sdk Upload file.")
	c.Assert(err, IsNil)
	c.Assert(resUpload.Code, Equals, 0)

	resUpload, err = s.bucket.Upload(srcPath, dstPath, "go testcase for cos sdk Upload file.")
	c.Assert(err, IsNil)

	resDelFile, err := s.bucket.Del(dstPath)
	c.Assert(err, IsNil)
	c.Assert(resDelFile.Code, Equals, 0)

}

func (s *BucketSuite) TestUpdateAndStatFile(c *C) {

	res, err := s.bucket.CreateFolder("/cos-go-sdk/TestUpdateAndStatFile/", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	resUpload, err := s.bucket.Upload(srcPath, "/cos-go-sdk/TestUpdateAndStatFile/test.jpg", "go testcase for cos sdk Upload file.")
	c.Assert(err, IsNil)
	c.Assert(resUpload.Code, Equals, 0)

	resUpdate, err := s.bucket.Update("/cos-go-sdk/TestUpdateAndStatFile/test.jpg", "update")
	c.Assert(err, IsNil)
	c.Assert(resUpdate.Code, Equals, 0)

	resStat, err := s.bucket.Stat("/cos-go-sdk/TestUpdateAndStatFile/test.jpg")
	c.Assert(err, IsNil)
	c.Assert(resStat.Code, Equals, 0)
	c.Assert(resStat.Data.BizAttr, Equals, "update")

	resStat, err = s.bucket.Stat("")
	c.Assert(err, IsNil)
	c.Assert(resStat.Code, Equals, 0)

	resDelFile, err := s.bucket.Del("/cos-go-sdk/TestUpdateAndStatFile/test.jpg")
	c.Assert(err, IsNil)
	c.Assert(resDelFile.Code, Equals, 0)

	resDelFolder, err := s.bucket.DelFolder("/cos-go-sdk/TestUpdateAndStatFile/")
	c.Assert(err, IsNil)
	c.Assert(resDelFolder.Code, Equals, 0)
}

func (s *BucketSuite) TestUploadSlice(c *C) {

	resUpload, err := s.bucket.Upload_slice(slice_srcPath, slice_dstPath, "go testcase for cos sdk Upload file slice.", 3*1024*1024, "")
	c.Assert(err, IsNil)
	c.Assert(resUpload.Code, Equals, 0)

	resUpload, err = s.bucket.Upload_slice(slice_srcPath, slice_dstPath, "go testcase for cos sdk Upload file slice.", 3*1024*1024, "")
	c.Assert(err, IsNil)
	c.Assert(resUpload.Code, Equals, 0)

	resDelFile, err := s.bucket.Del(slice_dstPath)
	c.Assert(err, IsNil)
	c.Assert(resDelFile.Code, Equals, 0)

	resUpload, err = s.bucket.Upload_slice(slice_srcPath, "./test/data1.bin", "go testcase for cos sdk Upload file slice.", 3*1024*1024, "")
	c.Assert(err, NotNil)

}

// func TestCreateFolder(t *testing.T) {
// 	response, err := bucket.CreateFolder(folderPath, "")
// 	if err != nil {
// 		t.Errorf("cos createFolder path=%s, failed, err=%s\n", folderPath, err.Error())
// 	} else {
// 		fmt.Printf("cos createFolder path=%s success\n", folderPath)
// 	}
// 	fmt.Println(response)

// }

// func TestUpdateFolder(t *testing.T) {

// 	bizAttr := "update"
// 	response, err := bucket.UpdateFolder(folderPath, bizAttr)
// 	if err != nil {
// 		t.Errorf("cos updateFolder path=%s, failed, err=%s\n", folderPath, err.Error())
// 	} else {
// 		fmt.Printf("cos updateFolder path=%s success\n", folderPath)
// 	}
// 	fmt.Println(response)
// }

// func TestStatFolder(t *testing.T) {
// 	json, err := bucket.StatFolder(folderPath)
// 	if err != nil {
// 		t.Errorf("cos statFolder path=%s, failed, err=%s\n", folderPath, err.Error())
// 	} else {
// 		fmt.Printf("cos statFolder path=%s success\n", folderPath)
// 		fmt.Println(json)
// 	}
// }

// func TestListFolder(t *testing.T) {
// 	json, err := bucket.ListFolder(folderPath, 20, "eListBoth", 0, "")
// 	if err != nil {
// 		t.Errorf("cos listFolder path=%s, failed, err=%s\n", folderPath, err.Error())
// 	} else {
// 		fmt.Println("cos listFolder path=%s success", folderPath)
// 		fmt.Println(json)
// 	}

// }

// func TestPrefixSearch(t *testing.T) {
// 	json, err := bucket.PrefixSearch(folderPath, 20, "eListBoth", 0, "")
// 	if err != nil {
// 		t.Errorf("cos prefixSearch path=%s, failed, err=%s\n", folderPath, err.Error())
// 	} else {
// 		fmt.Println("cos prefixSearch path=%s success", folderPath)
// 		fmt.Println(json)
// 	}

// }

// func TestDelFolder(t *testing.T) {
// 	json, err := bucket.DelFolder(folderPath)
// 	if err != nil {
// 		t.Errorf("cos delFolder path=%s, failed, err=%s\n", folderPath, err.Error())
// 	} else {
// 		fmt.Println("cos delFolder path=%s success", folderPath)
// 	}
// 	fmt.Println(json)
// }

// func TestUpload(t *testing.T) {

// 	_, err := bucket.Upload(srcPath, dstPath, "")
// 	if err != nil {
// 		t.Errorf("cos upload failed, srcPath=%s, err=%s\n", srcPath, err.Error())
// 	} else {
// 		fmt.Printf("cos upload srcPath=%s, dstPath=%s success\n", srcPath, dstPath)
// 	}
// }

// func TestStat(t *testing.T) {
// 	json, err := bucket.Stat(dstPath)
// 	if err != nil {
// 		t.Errorf("cos stat file=%s, failed, err=%s\n", dstPath, err.Error())
// 	} else {
// 		fmt.Printf("cos stat file=%s success\n", dstPath)
// 		fmt.Println(json)
// 	}
// }

// func TestUpdate(t *testing.T) {
// 	json, err := bucket.Update(dstPath, "update")
// 	if err != nil {
// 		t.Errorf("cos update file=%s, failed, err=%s\n", dstPath, err.Error())
// 	} else {
// 		fmt.Printf("cos update file=%s success\n", dstPath)
// 	}
// 	fmt.Println(json)
// }

// func TestDel(t *testing.T) {
// 	json, err := bucket.Del(dstPath)
// 	if err != nil {
// 		t.Errorf("cos del file=%s, failed, err=%s\n", dstPath, err.Error())
// 	} else {
// 		fmt.Println("cos del file=%s success", dstPath)
// 	}
// 	fmt.Println(json)
// }

// func TestUpload_slice(t *testing.T) {

// 	json, err := bucket.Upload_slice(slice_srcPath, slice_dstPath, "", 0, "")
// 	if err != nil {
// 		t.Errorf("cos upload_slice file=%s, failed, err = %s\n", slice_srcPath, err.Error())
// 	} else {
// 		fmt.Printf("cos Upload_slice file=%s dstPath=%s success", slice_srcPath, slice_dstPath)
// 	}
// 	fmt.Println(json)
// }

// func TestDel2(t *testing.T) {
// 	json, err := bucket.Del(slice_dstPath)
// 	if err != nil {
// 		t.Errorf("cos del file=%s, failed, err=%s\n", slice_dstPath, err.Error())
// 	} else {
// 		fmt.Println("cos del file=%s success", slice_dstPath)
// 	}
// 	fmt.Println(json)
// }

// func BenchmarkUpload(b *testing.B) {
// 	now := GetNowSec()
// 	for i := 0; i < 10; i++ {
// 		dstPath := fmt.Sprintf("%s_%d_%d.jpg", "/cos-go-sdk/test", now, i)
// 		_, _ = bucket.Upload(srcPath, dstPath, "")
// 	}
// }

// GetNowSec returns Unix time, the number of seconds elapsed since January 1, 1970 UTC.
// 获取当前时间，从UTC开始的秒数。
func GetNowSec() int64 {
	return time.Now().Unix()
}
