package coscloud

import (
	"crypto/x509"
	"fmt"
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

	//
	s.bucket.Client.Config.Pool = x509.NewCertPool()

	_, err := s.bucket.Upload_slice(slice_srcPath, slice_dstPath, "go testcase for cos sdk Upload file slice.", 3*1024*1024, "")
	c.Assert(err, NotNil)

	fmt.Println("TearDownSuite")
}

// Run before each test or benchmark starts
func (s *BucketSuite) SetUpTest(c *C) {
	s.bucket.Client.Config.Timeout = time.Second * 60
}

// Run after each test or benchmark runs.
func (s *BucketSuite) TearDownTest(c *C) {
}

var folderPath = func() string {
	now := GetNowSec()
	return fmt.Sprintf("%s_%d", "/cos-go-sdk/folder", now)
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

	res, err = s.bucket.CreateFolder("/", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Not(Equals), 0)

	res, err = s.bucket.CreateFolder("", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Not(Equals), 0)

	res, err = s.bucket.CreateFolder("//", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Not(Equals), 0)

	res, err = s.bucket.CreateFolder("cos-go-sdk/createFolder/test01", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	resDel, err = s.bucket.DelFolder("cos-go-sdk/createFolder/test01")
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)

	res, err = s.bucket.CreateFolder("/cos-go-sdk/createFolder/test02", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	resDel, err = s.bucket.DelFolder("/cos-go-sdk/createFolder/test02")
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)

	res, err = s.bucket.CreateFolder("cos-go-sdk/createFolder/test03/", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	resDel, err = s.bucket.DelFolder("cos-go-sdk/createFolder/test03/")
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)

	res, err = s.bucket.CreateFolder("/cos-go-sdk/createFolder/test04/", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	resDel, err = s.bucket.DelFolder("/cos-go-sdk/createFolder/test04/")
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)

	res, err = s.bucket.CreateFolder("/cos-go-sdk1/createFolder1/是滴ID我好滴/", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	resDel, err = s.bucket.DelFolder("/cos-go-sdk1/createFolder1/是滴ID我好滴/")
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

	res, err = s.bucket.CreateFolder(folderPath+"/01", "create")
	c.Assert(err, IsNil)
	c.Assert(res.Code, Equals, 0)

	res, err = s.bucket.CreateFolder(folderPath+"/02", "create")
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

	resSearch, err := s.bucket.PrefixSearch("/cos-go-sdk", 20, ELISTBOTH, Asc, "")
	c.Assert(err, IsNil)
	c.Assert(resSearch.Code, Equals, 0)

	resSearch, err = s.bucket.PrefixSearch("cos-go-sdk", 20, ELISTBOTH, Asc, "")
	c.Assert(err, IsNil)
	c.Assert(resSearch.Code, Equals, 0)

	resDel, err := s.bucket.DelFolder(folderPath + "/01")
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)

	resDel, err = s.bucket.DelFolder(folderPath + "/02")
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)

	resDel, err = s.bucket.DelFolder(folderPath)
	c.Assert(err, IsNil)
	c.Assert(resDel.Code, Equals, 0)
}

func (s *BucketSuite) TestUploadFile(c *C) {
	resUpload, err := s.bucket.Upload(srcPath, dstPath, "go testcase for cos sdk Upload file.")
	c.Assert(err, IsNil)
	c.Assert(resUpload.Code, Equals, 0)

	resUpload, err = s.bucket.Upload(srcPath, dstPath, "go testcase for cos sdk Upload file.")
	c.Assert(err, IsNil)
	c.Assert(resUpload.Code, Not(Equals), 0)

	//srcPath not exist
	resUpload, err = s.bucket.Upload("notExist", dstPath, "go testcase for cos sdk Upload file.")
	c.Assert(err, NotNil)

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

	resUpload, err = s.bucket.Upload_slice(slice_srcPath, slice_dstPath, "go testcase for cos sdk Upload file slice.", 4*1024*1024, "")
	c.Assert(err, IsNil)
	c.Assert(resUpload.Code, Equals, 0)

	//wrong session
	resUpload, err = s.bucket.Upload_slice(slice_srcPath, slice_dstPath, "go testcase for cos sdk Upload file slice.", 2*1024*1024, "session")
	c.Assert(err, NotNil)

	resDelFile, err := s.bucket.Del(slice_dstPath)
	c.Assert(err, IsNil)
	c.Assert(resDelFile.Code, Equals, 0)

	resUpload, err = s.bucket.Upload_slice("./test/data1.bi", slice_dstPath, "go testcase for cos sdk Upload file slice.", 3*1024*1024, "")
	c.Assert(err, NotNil)

}

// set Timeout
func (s *BucketSuite) TestUploadSlice2(c *C) {

	s.bucket.Client.Config.Timeout = time.Second * 2

	resUpload, err := s.bucket.Upload_slice(slice_srcPath, slice_dstPath, "go testcase for cos sdk Upload file slice.", 3*1024*1024, "")
	c.Assert(err, NotNil)
	// fmt.Println(err)

	s.bucket.Client.Config.Timeout = time.Second * 60
	resUpload, err = s.bucket.Upload_slice(slice_srcPath, slice_dstPath, "go testcase for cos sdk Upload file slice.", 3*1024*1024, "")
	c.Assert(err, IsNil)
	c.Assert(resUpload.Code, Equals, 0)

	resDelFile, err := s.bucket.Del(slice_dstPath)
	c.Assert(err, IsNil)
	c.Assert(resDelFile.Code, Equals, 0)

}

// try to coverage 95%
func (s *BucketSuite) TestTimeOutForPassCoverAlls(c *C) {

	s.bucket.Client.Config.Timeout = time.Millisecond * 10

	_, err := s.bucket.CreateFolder(folderPath, "create")
	c.Assert(err, NotNil)

	_, err = s.bucket.DelFolder(folderPath)
	c.Assert(err, NotNil)

	_, err = s.bucket.UpdateFolder(folderPath, "update")
	c.Assert(err, NotNil)

	_, err = s.bucket.StatFolder(folderPath)
	c.Assert(err, NotNil)

	_, err = s.bucket.ListFolder(folderPath, 20, ELISTFILEONLY, Desc, "")
	c.Assert(err, NotNil)

	_, err = s.bucket.PrefixSearch("/cos-go-sdk", 20, ELISTBOTH, Asc, "")
	c.Assert(err, NotNil)

	_, err = s.bucket.Upload(srcPath, "/cos-go-sdk/TestUpdateAndStatFile/test.jpg", "go testcase for cos sdk Upload file.")
	c.Assert(err, NotNil)

	_, err = s.bucket.Update("/cos-go-sdk/TestUpdateAndStatFile/test.jpg", "update")
	c.Assert(err, NotNil)

	_, err = s.bucket.Stat("/cos-go-sdk/TestUpdateAndStatFile/test.jpg")
	c.Assert(err, NotNil)

	_, err = s.bucket.Del("/cos-go-sdk/TestUpdateAndStatFile/test.jpg")
	c.Assert(err, NotNil)

	_, err = s.bucket.Upload_slice(slice_srcPath, slice_dstPath, "go testcase for cos sdk Upload file slice.", 2*1024*1024, "session")
	c.Assert(err, NotNil)

}

// GetNowSec returns Unix time, the number of seconds elapsed since January 1, 1970 UTC.
// 获取当前时间，从UTC开始的秒数。
func GetNowSec() int64 {
	return time.Now().Unix()
}
