package mns

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

var (
	AccessKeyId, AccessKeySecret string
	Url                          string
	Version                      string = "2015-06-06"
	ContentType                  string = "text/xml"
)

func GetHeader(method, resource string, exheader map[string]string) (header map[string]string) {
	if exheader == nil {
		exheader = make(map[string]string)
	}
	// 默认版本号
	exheader["x-mns-version"] = Version

	// 时区
	// now := time.Now().UTC()
	// loc, err := time.LoadLocation("GMT")
	// if err == nil {
	// 	now = now.In(loc)
	// }

	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	auth := getAuthorization(method, ContentType, date, getMNSHeaderStr(exheader), resource)

	murl := Url
	purl, _ := url.Parse(murl)
	host := purl.Host

	header = exheader
	header["Authorization"] = auth
	header["Date"] = date
	header["Host"] = host
	header["Content-Type"] = ContentType
	return
}

/*
Authorization = base64(hmac-sha1(HTTP_METHOD + "\n"
                + CONTENT-MD5 + "\n"
                + CONTENT-TYPE + "\n"
                + DATE + "\n"
                + CanonicalizedMNSHeaders
                + CanonicalizedResource))
*/
func getAuthorization(Method, ContentType, Date, CanonicalizedMNSHeaders, CanonicalizedResource string) (auth string) {
	auth = fmt.Sprintf("%s\n%s\n%s\n%s\n%s%s",
		Method,
		"",
		ContentType,
		Date,
		CanonicalizedMNSHeaders,
		CanonicalizedResource)

	//hmac-sha1
	key := []byte(AccessKeySecret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(auth))
	// fmt.Printf("%x\n", mac.Sum(nil))

	//base64
	auth = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	auth = "MNS " + AccessKeyId + ":" + auth
	fmt.Println(Date, auth)
	return
}

/*
注意：CanonicalizedMNSHeaders（即x-mns-开头的head）在签名验证前需要符合以下规范:
head的名字需要变成小写；
head自小到大排序；
分割head name和value的冒号前后不能有空格；
每个Head之后都有一个\n，如果没有以x-mns-开头的head，则在签名时CanonicalizedMNSHeaders就设置为空。
*/
func getMNSHeaderStr(CanonicalizedMNSHeaders map[string]string) (hstr string) {
	if CanonicalizedMNSHeaders == nil {
		return
	}
	headers := make(map[string]string)
	var keys []string
	for k, v := range CanonicalizedMNSHeaders {
		key := strings.ToLower(k)
		if strings.HasPrefix(key, "x-mns-") {
			headers[key] = v
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	for _, v := range keys {
		hstr += fmt.Sprintf("%s:%s\n", v, headers[v])
	}
	return
}
