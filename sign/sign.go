package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func signBase(appid uint, secretId string, secretKey string, bucket string, now int64, rdm int32, expireTime uint, fileid string) (string, error) {
	if "" == secretId || "" == secretKey {
		return "", errors.New("invalid params, secret id or key is empty")
	}

	var plainStr string
	plainStr = fmt.Sprintf("a=%d&b=%s&k=%s&e=%d&t=%d&r=%d&f=%s",
		appid,
		bucket,
		secretId,
		expireTime,
		now,
		rdm,
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
	now, rdm, expireTime := genBaseParamsToSign(expire)
	return signBase(appid, secretId, secretKey, bucket, now, rdm, expireTime, "")
}

// gen the sign binding a fileid
func AppSignOnce(appid uint, secretId string, secretKey string, bucket string, fileid string) (string, error) {
	now, rdm, expireTime := genBaseParamsToSign(0)
	return signBase(appid, secretId, secretKey, bucket, now, rdm, expireTime, fileid)
}

func genBaseParamsToSign(expire uint) (int64, int32, uint) {
	now := time.Now().Unix()
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	rdm := r.Int31()
	expireTime := expire
	if 0 != expireTime {
		expireTime += uint(now)
	}
	return now, rdm, expireTime
}
