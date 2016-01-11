// copyright : tencent
// author : weili
// github :

// Package Bucket implements go sdk for cos
package coscloud

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"qcloud/cos-go-sdk/sign"
)

const (
	SDK_VERSION     = "1.0.0"
	COSCLOUD_DOMAIN = "web.file.myqcloud.com"
	SIGN_ONCE       = "sign_once"
	SIGN            = "sign"

	ELISTBOTH     = "eListBoth"
	ELISTDIRONLY  = "eListDirOnly"
	ELISTFILEONLY = "eListFileOnly"

	COSAPI_ILLEGAL_SLICE_SIZE_ERROR = -4
)

type Bucket struct {
	BucketName string
	Client     Client
}

type Client struct {
	Config *Config
	Conn   *Conn
}

/*
	Bucket 构造方法
	@param uint 	appId		授权appid
	@param string 	secretId	授权secret id
	@param string 	secretKey	授权secret key
	@param string 	bucketName	bucket名称
*/
func New(appid uint, secretId, secretKey, bucketName string) (*Bucket, error) {
	config := getDefaultCosConfig()
	config.Appid = appid
	config.SecretId = secretId
	config.SecretKey = secretKey

	bucketUrl := buildBucketUrl(config.Endpoint, appid, bucketName)
	conn := &Conn{config, bucketUrl}

	client := Client{
		config,
		conn,
	}
	bucket := &Bucket{
		bucketName,
		client,
	}

	return bucket, nil
}

/*
	构建访问bucket URL
	@param string  	endpoint
	@param uint  	appId	 		授权appid
	@param string 	bucketName	  	bucket name
*/
func buildBucketUrl(endpoint string, appid uint, bucketName string) string {

	if match, _ := regexp.MatchString(`\/$`, endpoint); !match {
		endpoint += "/"
	}

	return fmt.Sprintf("%s%d/%s", endpoint, appid, bucketName)

}

//GET params
type Params map[string]string

// Encode encodes the Params into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (p Params) encode() string {
	if p == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := p[k]
		prefix := k + "="
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(prefix)
		buf.WriteString(url.QueryEscape(vs))
	}
	return buf.String()
}

//格式化文件路径， 以`/`开始
func formatFilePath(path string) string {
	if match, _ := regexp.MatchString(`^\/`, path); !match {
		path = "/" + path
	}

	return path
}

//格式化目录路径，以`/`开始，且`/`结束
func formatFolderPath(path string) string {
	if match, _ := regexp.MatchString(`^\/`, path); !match {
		path = "/" + path
	}
	if match, _ := regexp.MatchString(`\/$`, path); !match {
		path = path + "/"
	}

	return path
}

//url编码
func cosUrlEncode(path string) string {
	return strings.Replace(url.QueryEscape(path), "%2F", "/", -1)
}

//获取文件sha1
//param path string 文件路径
func getFileSha1(path string) (string, error) {
	if "" == path {
		err := errors.New("invalid srcPath")
		return "", err
	}

	fi, err := os.Open(path)
	if nil != err {
		return "", err
	}
	defer fi.Close()

	h := sha1.New()
	_, erro := io.Copy(h, fi)
	if erro != nil {
		return "", erro
	}

	hmacStr := fmt.Sprintf("%x", h.Sum(nil))
	return hmacStr, nil

}

//获取文件sha1
//param bytes []byte
func getbytesSha1(bytes []byte) (string, error) {

	h := sha1.New()
	_, erro := h.Write(bytes)
	if erro != nil {
		return "", erro
	}

	hmacStr := fmt.Sprintf("%x", h.Sum(nil))
	return hmacStr, nil

}

//获取文件内容
func getFileContents(filePath string) ([]byte, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	filecontent, err := ioutil.ReadAll(file)
	if nil != err {
		return nil, err
	}
	return filecontent, nil

}

//获取文件分片数据
func getFileSliceCntents(srcPath string, offset int64, sliceSize int) ([]byte, error) {

	file, err := os.Open(srcPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sr := io.NewSectionReader(file, offset, int64(sliceSize))
	buf := make([]byte, sliceSize)
	n, err := sr.Read(buf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return buf[:n], nil
}

//解析response
//return json数据
func (buc Bucket) parseRsp(rsp []byte) (js *simplejson.Json, err error) {
	//fmt.Printf("http rsp : %s\r\n", string(rsp))
	js, err = simplejson.NewJson(rsp)
	if nil != err {
		return
	}
	code, err := js.Get("code").Int()
	if nil != err {
		return
	}
	message, err := js.Get("message").String()
	if nil != err {
		return
	}

	if code != 0 {
		desc := fmt.Sprintf("rsp error, code=%d, message=%s", code, message)
		err = errors.New(desc)
		return
	}

	return
}

//post 请求数据json编码，返回请求body， headers
func jsonReqData(reqData map[string]string) (io.Reader, map[string]string, error) {
	d, err := json.Marshal(reqData)
	if nil != err {
		fmt.Printf("json.Marshal error, err=%s", err.Error())
		return nil, nil, err
	}
	body := bytes.NewBuffer([]byte(d))
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	return body, headers, nil
}

//multipart/form-data 请求body，headers
func multipartReqData(reqData map[string]string, filecontent []byte, boundary string) (io.Reader, map[string]string, error) {

	body := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(body)
	bodyWriter.SetBoundary(boundary)

	if nil != filecontent {
		fileWriter, err := bodyWriter.CreateFormField("filecontent")
		if nil != err {
			return nil, nil, err
		}
		_, err = fileWriter.Write(filecontent)
		if nil != err {
			return nil, nil, err
		}
	}

	for key, val := range reqData {
		_ = bodyWriter.WriteField(key, val)
	}
	err := bodyWriter.Close()
	if err != nil {
		return nil, nil, err
	}

	headers := map[string]string{
		"Content-Type": "multipart/form-data; boundary=" + boundary,
	}

	return body, headers, nil
}

/*
	创建目录
	@param string path     目录路径，sdk会补齐末尾的 '/'
	@param string bizAttr  目录属性
*/
func (buc Bucket) CreateFolder(path, bizAttr string) (js *simplejson.Json, err error) {

	path = formatFolderPath(path)

	reqData := map[string]string{
		"op":       "create",
		"biz_attr": bizAttr,
	}

	body, headers, err := jsonReqData(reqData)
	if nil != err {
		return
	}

	data, err := buc.do("POST", path, nil, headers, body, SIGN)
	if nil != err {
		return
	}

	return buc.parseRsp(data)

}

// Private
/*
	do request
	@param string  				method  		请求方法
	@param string  				path    		format后的目录或文件路径
	@param Params  				params  		‘GET’ 方法请求参数
	@param map[string]string  	headers  	 	请求头
	@param io.Reader  			data  		 	请求body
	@param string  				signType 		签名方式，SIGN_ONCE， SIGN

*/
func (buc Bucket) do(method, path string, params Params,
	headers map[string]string, data io.Reader, signType string) ([]byte, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	path = cosUrlEncode(path)
	var signHead string
	if signType == SIGN_ONCE {
		sign, err := buc.signOnce(path)
		if nil != err {
			fmt.Printf("SignOnce error, err=%s", err.Error())
			return nil, err
		}
		signHead = sign

	} else {
		sign, err := buc.sign(buc.Client.Config.SignExpiredSeconds)
		if nil != err {
			fmt.Printf("Sign error, err=%s", err.Error())
			return nil, err
		}
		signHead = sign
	}
	headers["Authorization"] = signHead

	urlParams := ""
	if params != nil {
		urlParams = params.encode()
	}

	return buc.Client.Conn.Do(method, path, urlParams, headers, data)
}

/*
	目录信息 更新
	@param string path     目录路径，sdk会补齐末尾的 '/'
	@param string bizAttr  更新信息
*/
func (buc Bucket) UpdateFolder(path, bizAttr string) (js *simplejson.Json, err error) {
	path = formatFolderPath(path)
	return buc.updateBase(path, bizAttr)
}

/*
	文件信息 更新
	@param string path     文件路径
	@param string bizAttr  更新信息
*/
func (buc Bucket) Update(path, bizAttr string) (js *simplejson.Json, err error) {
	path = formatFilePath(path)
	return buc.updateBase(path, bizAttr)
}

//Private
func (buc Bucket) updateBase(path, bizAttr string) (js *simplejson.Json, err error) {

	reqData := map[string]string{
		"op":       "update",
		"biz_attr": bizAttr,
	}

	body, headers, err := jsonReqData(reqData)
	if nil != err {
		return
	}

	data, err := buc.do("POST", path, nil, headers, body, SIGN_ONCE)
	if nil != err {
		return
	}

	return buc.parseRsp(data)
}

/*
	目录信息 查询
	@param string  path  目录路径
*/
func (buc Bucket) StatFolder(path string) (js *simplejson.Json, err error) {

	path = formatFolderPath(path)
	return buc.statBase(path)
}

/*
	文件信息 查询
	@param string  path 文件路径
*/
func (buc Bucket) Stat(path string) (js *simplejson.Json, err error) {

	path = formatFilePath(path)
	return buc.statBase(path)
}

//Private
func (buc Bucket) statBase(path string) (js *simplejson.Json, err error) {

	params := Params{
		"op": "stat",
	}

	data, err := buc.do("GET", path, params, nil, nil, SIGN)
	if nil != err {
		return
	}

	return buc.parseRsp(data)
}

/*
	删除目录
	@param string path  目录路径
*/
func (buc Bucket) DelFolder(path string) (js *simplejson.Json, err error) {

	path = formatFolderPath(path)
	return buc.delBase(path)

}

/*
	删除文件
	@param string path  文件路径
*/
func (buc Bucket) Del(path string) (js *simplejson.Json, err error) {

	path = formatFilePath(path)
	return buc.delBase(path)

}

//Private
func (buc Bucket) delBase(path string) (js *simplejson.Json, err error) {

	reqData := map[string]string{
		"op": "delete",
	}

	body, headers, err := jsonReqData(reqData)
	if nil != err {
		return
	}

	data, err := buc.do("POST", path, nil, headers, body, SIGN_ONCE)
	if nil != err {
		return
	}

	return buc.parseRsp(data)
}

/*
	目录列表
    @param  string  path      	目录路径，以"/"开头,以"/"结尾，api会补齐
    @param  int     num      	拉取的总数
    @param  string  pattern  	eListBoth,eListDirOnly,eListFileOnly  默认eListBoth
    @param  int     order       默认正序(=0), 填1为反序,
    @param  string  context     透传字段，查看第一页，则传空字符串。若需要翻页，需要将前一页返回值中的context透传到参数中。order用于指定翻页顺序。若order填0，则从当前页正序/往下翻页；若order填1，则从当前页倒序/往上翻页
*/
func (buc Bucket) ListFolder(path string, num int,
	pattern string, order int, context string) (*simplejson.Json, error) {

	path = formatFolderPath(path)
	return buc.listBase(path, num, pattern, order, context)
}

/*
	前缀搜索
    @param  string  prefix      列出含此前缀的所有文件
    @param  int     num      	拉取的总数
    @param  string  pattern  	eListBoth,eListDirOnly,eListFileOnly  默认eListBoth
    @param  int     order       默认正序(=0), 填1为反序,
    @param  string  context     透传字段，查看第一页，则传空字符串。若需要翻页，需要将前一页返回值中的context透传到参数中。order用于指定翻页顺序。若order填0，则从当前页正序/往下翻页；若order填1，则从当前页倒序/往上翻页
*/
func (buc Bucket) PrefixSearch(prefix string, num int,
	pattern string, order int, context string) (*simplejson.Json, error) {

	if match, _ := regexp.MatchString(`^\/`, prefix); !match {
		prefix = "/" + prefix
	}
	return buc.listBase(prefix, num, pattern, order, context)
}

//Private
func (buc Bucket) listBase(path string, num int,
	pattern string, order int, context string) (js *simplejson.Json, err error) {

	params := Params{
		"op":      "list",
		"num":     strconv.Itoa(num),
		"pattern": pattern,
		"order":   strconv.Itoa(order),
		"context": context,
	}

	data, err := buc.do("GET", path, params, nil, nil, SIGN)
	if nil != err {
		return
	}

	return buc.parseRsp(data)
}

/*
	上传文件 (一般小于8MB)的上传
    @param  string  srcPath      本地文件路径
    @param  string  dstPath      上传的文件路径
    @param  string  bizAttr      文件属性
*/
func (buc Bucket) Upload(srcPath, dstPath, bizAttr string) (js *simplejson.Json, err error) {

	//file sha1
	sha, err := getFileSha1(srcPath)
	if nil != err {
		fmt.Printf("getFileSha1 error, err=%s", err.Error())
		return
	}

	reqData := map[string]string{
		"op":       "upload",
		"sha":      sha,
		"biz_attr": bizAttr,
	}

	boundary := "-------------------------abcdefg1234567"

	filecontent, err := getFileContents(srcPath)
	if nil != err {
		return
	}

	body, headers, err := multipartReqData(reqData, filecontent, boundary)
	if nil != err {
		return
	}

	data, err := buc.do("POST", dstPath, nil, headers, body, SIGN)
	if nil != err {
		return
	}

	return buc.parseRsp(data)
}

/*
	分片上传
    @param  string  srcPath      本地文件路径
    @param  string  dstPath      上传的文件路径
    @param  string  bizAttr      文件属性
    @param  int  	sliceSize    分片大小，字节数。比如100 * 1024为每片100KB。
    @param  string  session      文件传输过程的id
*/
func (buc Bucket) Upload_slice(srcPath, dstPath, bizAttr string, sliceSize int, session string) (js *simplejson.Json, err error) {

	filemode, err := os.Stat(srcPath)
	if nil != err {
		fmt.Println("os.Stat error", err)
	}

	fileSize := filemode.Size()

	if fileSize < buc.Client.Config.MinSliceFileSize {
		return buc.Upload(srcPath, dstPath, bizAttr)
	}

	//file sha1
	sha1, err := getFileSha1(srcPath)
	if nil != err {
		fmt.Printf("getFileSha1 error, err=%s", err.Error())
		return
	}

	reqData := map[string]string{
		"op":       "upload_slice",
		"filesize": strconv.Itoa(int(fileSize)),
		"sha":      sha1,
	}

	if "" != bizAttr {
		reqData["biz_attr"] = bizAttr
	}
	if "" != session {
		reqData["session"] = session
	}

	if sliceSize > 0 {
		if sliceSize <= buc.Client.Config.DefaultSliceSize {
			reqData["slice_size"] = strconv.Itoa(sliceSize)
		} else {
			reqData["slice_size"] = strconv.Itoa(buc.Client.Config.DefaultSliceSize)
		}
	}

	boundary := "-------------------------abcdefg1234567"

	body, headers, err := multipartReqData(reqData, nil, boundary)
	if nil != err {
		return
	}

	data, err := buc.do("POST", dstPath, nil, headers, body, SIGN)
	if nil != err {
		return
	}

	js, err = buc.parseRsp(data)
	if nil != err {
		return
	}

	_, err = js.Get("data").Get("access_url").String()
	if nil == err { //秒传命中，直接返回了url
		return
	}

	sliceSize, _ = js.Get("data").Get("slice_size").Int()
	// fmt.Pribucketntf("1111111111sliceSize ==%d\n", sliceSize)
	if sliceSize > buc.Client.Config.DefaultSliceSize || sliceSize <= 0 {
		js.Set("code", COSAPI_ILLEGAL_SLICE_SIZE_ERROR)
		js.Set("message", "illegal slice size")
		return
	}

	session, _ = js.Get("data").Get("session").String()
	offset, _ := js.Get("data").Get("offset").Int64()

	fmt.Println(js)

	return buc.upload_data(fileSize, sliceSize, dstPath, srcPath, offset, session)

}

//Private
func (buc Bucket) upload_data(fileSize int64, sliceSize int, dstPath, srcPath string,
	offset int64, session string) (js *simplejson.Json, err error) {

	boundary := "-------------------------abcdefg1234567"
	var retry_times uint = 0
	for fileSize > offset {

		fmt.Printf("offset:%d \n", offset)

		if (offset + int64(sliceSize)) > fileSize {
			sliceSize = int(fileSize - offset)
		}

		filecontent, err1 := getFileSliceCntents(srcPath, offset, sliceSize)
		if nil != err1 {
			fmt.Printf("[upload_data]:getFileSliceCntents error, err=%s", err1.Error())
			err = err1
			return
		}

		//file sha1
		// sha, err1 := getbytesSha1(filecontent)
		// if nil != err1 {
		// 	fmt.Printf("[upload_data]:getFileSha1 error, err=%s", err1.Error())
		// 	err = err1
		// 	return
		// }

		reqData := map[string]string{
			"op": "upload_slice",
			// "sha":     sha,
			"session": session,
			"offset":  strconv.Itoa(int(offset)),
		}

		body, headers, err1 := multipartReqData(reqData, filecontent, boundary)
		if nil != err1 {
			err = err1
			return
		}

		data, err1 := buc.do("POST", dstPath, nil, headers, body, SIGN)

		if nil != err1 {
			fmt.Printf("=========err= %s ================", err1.Error())
			err = err1
			if retry_times < buc.Client.Config.RetryTimes {
				retry_times++
				fmt.Println(retry_times)
				continue

			} else {
				return
			}

		}
		js, err = buc.parseRsp(data)
		if nil != err {
			if retry_times < buc.Client.Config.RetryTimes {
				retry_times++
				continue
			} else {
				return
			}

		}

		session, _ = js.Get("data").Get("session").String()
		offset += int64(sliceSize)
		retry_times = 0
	}

	return
}

// 多次有效签名
// @param  uint  expire   签名过期时间，单位秒
func (buc Bucket) sign(expire uint) (string, error) {
	return sign.AppSign(buc.Client.Config.Appid, buc.Client.Config.SecretId, buc.Client.Config.SecretKey, buc.BucketName, expire)
}

//单次有效签名
// @param  string  path   目录/文件路径
func (buc Bucket) signOnce(path string) (string, error) {
	fileId := fmt.Sprintf("/%d/%s%s", buc.Client.Config.Appid, buc.BucketName, path)
	return sign.AppSignOnce(buc.Client.Config.Appid, buc.Client.Config.SecretId, buc.Client.Config.SecretKey, buc.BucketName, fileId)
}
