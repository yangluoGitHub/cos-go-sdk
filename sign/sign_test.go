// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

package sign

import (
	"fmt"
	"testing"
)

const APPID = 200941
const SECRET_ID = "AKIDh51wIFHJ13Mbc5AWd37z6WmQwIdTghBu"
const SECRET_KEY = "SU4Qn0GoK0YRNS97p0l5rAsxwxcN6Il3"

func TestAppSign(t *testing.T) {
	var expire uint = 3600 * 24 * 7
	sign, err := AppSign(APPID, SECRET_ID, SECRET_KEY, expire)
	if err != nil {
		t.Errorf("gen sign failed, err = %s\n", err.Error())
	} else {
		fmt.Printf("gen sign success, sign = %s\n", sign)
	}
}

func TestAppSignOnce(t *testing.T) {
	fileid := "442d8ddf-59a5-4dd4-b5f1-e38499fb33b4"
	sign, err := AppSignOnce(APPID, SECRET_ID, SECRET_KEY, fileid)
	if err != nil {
		t.Errorf("gen sign failed, err = %s\n", err.Error())
	} else {
		fmt.Printf("gen sign success, sign = %s\n", sign)
	}
}

func TestDecode(t *testing.T) {
	//test1
	sign := "gh8WN5lyExipeQ5SAfzif13LWEJhPTIwMDk0MSZrPUFLSURoNTF3SUZISjEzTWJjNUFXZDM3ejZXbVF3SWRUZ2hCdSZlPTE0MzMxNjc0MDUmdD0xNDMzMTYzODA1JnI9MTkwNjQ3MzUwMyZ1PTEyMzQ1NiZmPQ=="
	expire, fileid, bucket, err := Decode(sign, APPID, SECRET_ID, SECRET_KEY)
	if err != nil {
		t.Error("decode error, err=%s\n", err.Error())
	} else if expire == 0 {
		t.Error("decode expire info error, expire=0\n")
	} else if fileid != "" {
		t.Error("decode fileid info error, fileid must be empty\n")
	} else if bucket != "" {
		t.Error("decode bucket info error, bucket must be empty\n")
	}
}

func TestDecode2(t *testing.T) {
	//test2
	sign := "ROAtPTf9pbN5vRMFpoKCMjI5gDFhPTIwMDk0MSZrPUFLSURoNTF3SUZISjEzTWJjNUFXZDM3ejZXbVF3SWRUZ2hCdSZlPTAmdD0xNDMzMTYzODA1JnI9MTkwNjQ3MzUwMyZ1PTEyMzQ1NiZmPTQ0MmQ4ZGRmLTU5YTUtNGRkNC1iNWYxLWUzODQ5OWZiMzNiNA=="
	expire, fileid, bucket, err := Decode(sign, APPID, SECRET_ID, SECRET_KEY)
	if err != nil {
		t.Error("decode error, err=%s\n", err.Error())
	} else if expire != 0 {
		t.Error("decode expire info error, expire must be 0\n")
	} else if fileid == "" {
		t.Error("decode fileid info error, fileid is empty\n")
	} else if bucket != "" {
		t.Error("decode bucket info error, bucket must be empty\n")
	}
}

func TestDecode3(t *testing.T) {
	//wrong sign
	sign := "gh8WN5lyExipeQ5SBfzif13LWEJhPTIwMDk0MSZrPUFLSURoNTF3SUZISjEzTWJjNUFXZDM3ejZXbVF3SWRUZ2hCdSZlPTE0MzMxNjc0MDUmdD0xNDMzMTYzODA1JnI9MTkwNjQ3MzUwMyZ1PTEyMzQ1NiZmPQ=="
	_, _, _, err := Decode(sign, APPID, SECRET_ID, SECRET_KEY)
	if err == nil {
		t.Error("decode error, this sign is wrong!\n")
	}
	//wrong appid
	sign = "76n8W8B0Y+fp1ClLLjX8vsRBkFNhPTIwMDk0MCZrPUFLSURoNTF3SUZISjEzTWJjNUFXZDM3ejZXbVF3SWRUZ2hCdSZlPTE0MzMxNjgzNTYmdD0xNDMzMTY0NzU2JnI9NDk0MTY2NjMwJnU9MTIzNDU2JmY9"
	_, _, _, err = Decode(sign, APPID, SECRET_ID, SECRET_KEY)
	if err == nil {
		t.Error("decode error, this sign is wrong!\n")
	}
}

func TestDecode4(t *testing.T) {
	//test2
	sign := "ROAtPTf=8pbN5vRMFpoKCMjI5gDFhPTIwMDk0MSZrPUFLSURoNTF3SUZISjEzTWJjNUFXZDM3ejZXbVF3SWRUZ2hCdSZlPTAmdD0xNDMzMTYzODA1JnI9MTkwNjQ3MzUwMyZ1PTEyMzQ1NiZmPTQ0MmQ4ZGRmLTU5YTUtNGRkNC1iNWYxLWUzODQ5OWZiMzNiNA=="
	_, _, _, err := Decode(sign, APPID, SECRET_ID, SECRET_KEY)
	if err == nil {
		t.Error("decode error, this sign is wrong!\n")
	}
	//wrong appid
	sign = "NGu6Vr0av2DNYNcDLInDFC/dWl9hPTIwMDk0MCZrPUFLSURoNTF3SUZISjEzTWJjNUFXZDM3ejZXbVF3SWRUZ2hCdSZlPTAmdD0xNDMzMTY0NzU2JnI9NDk0MTY2NjMwJnU9MTIzNDU2JmY9NDQyZDhkZGYtNTlhNS00ZGQ0LWI1ZjEtZTM4NDk5ZmIzM2I0"
	_, _, _, err = Decode(sign, APPID, SECRET_ID, SECRET_KEY)
	if err == nil {
		t.Error("decode error, this sign is wrong!\n")
	}
}
