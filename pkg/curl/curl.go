package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//发送GET请求
//url:请求地址
//response:请求返回的内容
func ChainGet(url string) (response string) {
	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		log.Fatalf("地址请求失败 err: %v", url)
	}
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := res.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(url, "地址请求失败")
		}
	}
	defer res.Body.Close()
	response = result.String()
	return
}

// ChainPost 发送POST请求 param: url:请求地址		data:POST请求提交的数据		contentType:请求体格式，如：application/json
func ChainPost(url string, data interface{}, contentType string) (content string, err error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if req != nil {
		req.Header.Add("content-type", contentType)
	}
	if err != nil {
		return url + " 地址请求失败" + err.Error(), err
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 10 * time.Second}
	resp, errs := client.Do(req)
	if errs != nil {
		fmt.Println(errs.Error())
		return url + " 地址请求超时", err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}
