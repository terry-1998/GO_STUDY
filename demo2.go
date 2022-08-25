package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}
type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		IsSubject string `json:"is_subject"`
		Item      struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL string `json:"image_url"`
		Sitelink string `json:"sitelink"`
		ID       string `json:"id"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your word~")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("an error occured", err)
			return
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		if err != nil {
			fmt.Println("an error occured", err)
			return
			continue
		}
		if input == "quit!" {
			fmt.Println("bye-bye")
			break
		}
		request := DictRequest{TransType: "en2zh", Source: input}
		translate_query(request)
	}
}
func translate_query(request DictRequest) {
	client := &http.Client{}
	//request := DictRequest{TransType: "en2zh", Source: "good"}
	buf, err := json.Marshal(request)
	word := request.Source
	if err != nil {
		log.Fatal(err)
	}

	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "http://api。interpreter.caiyunai.com/v1/dict", data)

	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="96", "Google Chrome";v="96"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("os-type", "web")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad statuscode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictresponse DictResponse
	err = json.Unmarshal(bodyText, &dictresponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "uk:", dictresponse.Dictionary.Prons.En, "us:", dictresponse.Dictionary.Prons.EnUs)
	for _, item := range dictresponse.Dictionary.Explanations {
		fmt.Println(item)
	}
	//fmt.Printf("%#\n", dictresponse)
}
