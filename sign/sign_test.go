package sign

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type SignSuite struct {
	Appid      uint
	SecretId   string
	SecretKey  string
	BucketName string
}

var _ = Suite(&SignSuite{})

// Run once when the suite starts running
func (s *SignSuite) SetUpSuite(c *C) {
	s.Appid = 10016057
	s.SecretId = "AKIDoUqGP9jLfYgqVgueaT0CneVWDha9tdUM"
	s.SecretKey = "kmabDFmmQkV9PGThSmfw8TVtrkmvrNHl"
	s.BucketName = "bucket1"
}

// Run once after all tests or benchmarks have finished running.
func (s *SignSuite) TearDownSuite(c *C) {
}

// Run before each test or benchmark starts
func (s *SignSuite) SetUpTest(c *C) {
}

// Run after each test or benchmark runs.
func (s *SignSuite) TearDownTest(c *C) {
}

func (s *SignSuite) TestAppSign(c *C) {
	var expire uint = 60
	_, err := AppSign(s.Appid, s.SecretId, s.SecretKey, s.BucketName, expire)
	c.Assert(err, IsNil)
}

func (s *SignSuite) TestAppSignOnce(c *C) {
	fileid := "442d8ddf-59a5-4dd4-b5f1-e38499fb33b4"
	_, err := AppSignOnce(s.Appid, s.SecretId, s.SecretKey, s.BucketName, fileid)
	c.Assert(err, IsNil)
}

func (s *SignSuite) TestAppSign2(c *C) {
	sign, _ := signBase(200001,
		"AKIDUfLUEUigQiXqm7CVSspKJnuaiIKtxqAv",
		"bLcPnl88WU30VY57ipRhSePfPdOfSruK",
		"newbucket",
		1436077115,
		11162,
		1438669115,
		"")
	expected := "5bIObv9KXNcITrcVNRGCLG3K6xxhPTIwMDAwMSZiPW5ld2J1Y2tldCZrPUFLSURVZkxVRVVpZ1FpWHFtN0NWU3NwS0pudWFpSUt0eHFBdiZlPTE0Mzg2NjkxMTUmdD0xNDM2MDc3MTE1JnI9MTExNjImZj0="
	c.Assert(sign, Equals, expected)
}

func (s *SignSuite) TestSignOnce2(c *C) {
	sign, _ := signBase(200001,
		"AKIDUfLUEUigQiXqm7CVSspKJnuaiIKtxqAv",
		"bLcPnl88WU30VY57ipRhSePfPdOfSruK",
		"newbucket",
		1436077115,
		11162,
		1438669115,
		"tencentyunSignTest")
	expected := "RqnnBP1zUQjcCAeDlRX0cJSDhuNhPTIwMDAwMSZiPW5ld2J1Y2tldCZrPUFLSURVZkxVRVVpZ1FpWHFtN0NWU3NwS0pudWFpSUt0eHFBdiZlPTE0Mzg2NjkxMTUmdD0xNDM2MDc3MTE1JnI9MTExNjImZj10ZW5jZW50eXVuU2lnblRlc3Q="
	c.Assert(sign, Equals, expected)
}
