// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

// this is a demo for qcloud go sdk
package main

import (
	"fmt"
	"qcloud/cos-go-sdk"
	// "io/ioutil"
	// "os"
)

func main() {
	// var appid uint = 10000001
	// sid := "AKIDNZwDVhbRtdGkMZQfWgl2Gnn1dhXs95C0"
	// skey := "ZDdyyRLCLv1TkeYOl5OCMLbyH4sJ40wp"
	// bucket := "testb"

	// var appid uint = 10016191
	// sid := "AKID87beD5ugjJUdj7CvlxpX4sG0PT18fZJi"
	// skey := "2RN7eMlzI8ZqZqYmT5NM8mqZrqgv8aSf"
	// bucket := "space1"

	var appid uint = 10016057
	sid := "AKIDoUqGP9jLfYgqVgueaT0CneVWDha9tdUM"
	skey := "kmabDFmmQkV9PGThSmfw8TVtrkmvrNHl"
	bucket := "bucket1"

	cloud := coscloud.CosCloud{appid, sid, skey, bucket}
	fmt.Println("================CreateFolder=========================")
	json, err := cloud.CreateFolder("为立科技2", "test")
	if err != nil {
		fmt.Printf("cos createFolder failed, err = %s\n", err.Error())
	} else {
		fmt.Println("cos createFolder success")
	}
	fmt.Println(json)

	fmt.Println("================ListFolder=========================")
	json2, err := cloud.ListFolder("/", 20, "eListBoth", 0, "")
	if err != nil {
		fmt.Printf("cos ListFolder failed, err = %s\n", err.Error())
	} else {
		fmt.Println("cos ListFolder success")
	}
	fmt.Println(json2)

	fmt.Println("================UpdateFolder=========================")
	json3, err := cloud.UpdateFolder("为立科技2", "testUpdateFaa")
	if err != nil {
		fmt.Printf("cos UpdateFolder failed, err = %s\n", err.Error())
	} else {
		fmt.Println("cos UpdateFolder success")
	}
	fmt.Println(json3)

	fmt.Println("==================StatFolder=======================")
	json4, err := cloud.StatFolder("为立科技2")
	if err != nil {
		fmt.Printf("cos StatFolder failed, err = %s\n", err.Error())
	} else {
		fmt.Println("cos StatFolder success")
	}
	fmt.Println(json4)

	// fmt.Println("===============DelFolder==========================")
	// json5, err := cloud.DelFolder("为立科技2")
	// if err != nil {
	// 	fmt.Printf("cos DelFolder failed, err = %s\n", err.Error())
	// } else {
	// 	fmt.Println("cos DelFolder success")
	// }
	// fmt.Println(json5)

	// fmt.Println("================PrefixSearch=========================")
	// json6, err := cloud.PrefixSearch("/aaTest", 20, "eListBoth", 0, "")
	// if err != nil {
	// 	fmt.Printf("cos PrefixSearch failed, err = %s\n", err.Error())
	// } else {
	// 	fmt.Println("cos PrefixSearch success")
	// }
	// fmt.Println(json6)

	// fmt.Println("================Upload=========================")
	// json7, err := cloud.Upload("./pic/food.jpg", "/aaTest/food_test.jpg", "biz_attr")
	// if err != nil {
	// 	fmt.Printf("cos Upload failed, err = %s\n", err.Error())
	// } else {
	// 	fmt.Println("cos Upload success")
	// }
	// fmt.Println(json7)

	fmt.Println("=================Upload_slice========================")
	json8, err := cloud.Upload_slice("/Users/renxiaoqing/Downloads/apache-activemq-5.12.1-bin.tar.gz", "/aaTest/test001_apache-activemq-5.12.1-bin.tar.gz", "", 1024*1024, "")
	if err != nil {
		fmt.Printf("cos Upload_slice failed, err = %s\n", err.Error())
	} else {
		fmt.Println("cos Upload_slice success")
	}
	fmt.Println(json8)

	// fmt.Println("===================Update======================")
	// json9, err := cloud.Update("/aaTest/food_test.jpg", "testUpdateFaa")
	// if err != nil {
	// 	fmt.Printf("cos Update failed, err = %s\n", err.Error())
	// } else {
	// 	fmt.Println("cos Update success")
	// }
	// fmt.Println(json9)

	// fmt.Println("=====================Stat====================")
	// json10, err := cloud.Stat("/aaTest/food_test.jpg")
	// if err != nil {
	// 	fmt.Printf("cos Stat failed, err = %s\n", err.Error())
	// } else {
	// 	fmt.Println("cos Stat success")
	// }
	// fmt.Println(json10)

	// fmt.Println("=================Del========================")
	// json11, err := cloud.Del("/aaTest/food_test.jpg")
	// if err != nil {
	// 	fmt.Printf("cos Del failed, err = %s\n", err.Error())
	// } else {
	// 	fmt.Println("cos Del success")
	// }
	// fmt.Println(json11)
}
