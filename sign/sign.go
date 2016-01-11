// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

// Package sign implements sign for qcloud sdk
package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const HMAC_LENGTH = 20

func SignBase(appid uint, secretId string, secretKey string, bucket string, expire uint, fileid string) (string, error) {
	if "" == secretId || "" == secretKey {
		return "", errors.New("invalid params, secret id or key is empty")
	}

	now := time.Now().Unix()
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	rdm := r.Int31()
	expireTime := expire
	if 0 != expireTime {
		expireTime += uint(now)
	}

	var plainStr string
	plainStr = fmt.Sprintf("a=%d&b=%s&k=%s&e=%d&t=%d&r=%d&u=%s&f=%s",
		appid,
		bucket,
		secretId,
		expireTime,
		now,
		rdm,
		"0",
		fileid)

	cryptoStr := []byte(plainStr)
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write(cryptoStr)
	hmacStr := h.Sum(nil)

	cryptoStr = append(hmacStr, cryptoStr...)
	sign := base64.StdEncoding.EncodeToString(cryptoStr)
	// fmt.Println(sign)
	return sign, nil
}

// gen the sign with a expire time.
func AppSign(appid uint, secretId string, secretKey string, bucket string, expire uint) (string, error) {
	return SignBase(appid, secretId, secretKey, bucket, expire, "")
}

// gen the sign binding a fileid(pic resource)
func AppSignOnce(appid uint, secretId string, secretKey string, bucket string, fileid string) (string, error) {
	return SignBase(appid, secretId, secretKey, bucket, 0, fileid)
}

// decode a sign
func Decode(sign string, appid uint, secretId string, secretKey string) (expire uint, fileid string, bucket string, e error) {
	if "" == sign {
		e = errors.New("invalid sign string")
		return
	}

	cryptoStr, e := base64.StdEncoding.DecodeString(sign)
	if nil != e {
		return
	} else if len(cryptoStr) <= HMAC_LENGTH {
		e = errors.New("sign is too short")
		return
	}

	hmacStr := cryptoStr[0:HMAC_LENGTH]
	cryptoStr = cryptoStr[HMAC_LENGTH:]

	//check hmac str
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write(cryptoStr)
	hmacStr2 := h.Sum(nil)
	if len(hmacStr) != len(hmacStr2) {
		desc := fmt.Sprintf("hmac check failed, hmac1=%s, hmac2=%s", hmacStr, hmacStr2)
		e = errors.New(desc)
		return
	}

	for i := range hmacStr {
		if hmacStr[i] != hmacStr2[i] {
			desc := fmt.Sprintf("hmac check failed, hmac1=%s, hmac2=%s", hmacStr, hmacStr2)
			e = errors.New(desc)
			return
		}
	}

	//check cryto string
	fields := strings.Split(string(cryptoStr), "&")
	cnt := 0
	//check appid
	if fields[cnt] != ("a=" + strconv.Itoa(int(appid))) {
		desc := fmt.Sprintf("invalid appid, appid=%d, sign=%s", appid, fields[0])
		e = errors.New(desc)
		return
	}
	cnt++
	//check skey
	if strings.HasPrefix(fields[cnt], "b=") {
		//v2
		bucket = strings.TrimLeft(fields[cnt], "b=")
		cnt++
	}
	if fields[cnt] != ("k=" + secretId) {
		desc := fmt.Sprintf("invalid secret_id, sid=%s, sign=%s", secretId, fields[1])
		e = errors.New(desc)
		return
	}
	cnt++
	//check time
	//[3] is create time
	//[2] is expire time
	tmp, e := strconv.Atoi(strings.TrimLeft(fields[cnt], "e="))
	if nil != e {
		return
	}
	expire = uint(tmp)
	cnt += 4
	//check fileid
	fileid = strings.TrimLeft(fields[cnt], "f=")
	/////
	if (expire == 0 && fileid == "") ||
		(expire != 0 && fileid != "") {
		desc := fmt.Sprintf("invalid expire time or fileid, expire=%s, fileid=%s", expire, fileid)
		e = errors.New(desc)
		return
	}

	e = nil
	return
}
