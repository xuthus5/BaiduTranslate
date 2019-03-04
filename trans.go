package BaiduTranslate

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//百度翻译开放平台信息
type BaiduInfo struct {
	AppID     string
	Salt      string
	SecretKey string
	From      string
	To        string
	Text      string
}

//返回结果
type TransResult struct {
	From   string    `json:"from"`
	To     string    `json:"to"`
	Result [1]Result `json:"trans_result"`
}
type Result struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

//自动生盐
//入口参数为盐的长度
func Salt(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成32位MD5
func Sign(bi *BaiduInfo) string {
	text := bi.AppID + bi.Text + bi.Salt + bi.SecretKey
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//翻译 传入需要翻译的语句
func (bi *BaiduInfo) Translate() string{
	url := "http://api.fanyi.baidu.com/api/trans/vip/translate?q=" + bi.Text + "&from=" + bi.From + "&to=" + bi.To + "&appid=" + bi.AppID + "&salt=" + bi.Salt + "&sign=" + Sign(bi)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("网络异常")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var ts TransResult
	_ = json.Unmarshal(body, &ts)
	return ts.Result[0].Dst
}
