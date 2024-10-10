package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type response struct {
	Data data `json:"data,omitempry"`
}
type data struct {
	Attributes attributes `json:"attributes,omitempry"`
}
type attributes struct {
	File_link_all []file_link_all `json:"file_link_all,omitempry"`
}
type file_link_all struct {
	Text    string   `json:"text,omitempry"`
	Profile string   `json:"profile,omitempry"`
	Urls    []string `json:"urls,omitempry"`
}

func download(name, url string) {
	out, err := os.Create("files/" + name + ".mp4")
	if err != nil {
		fmt.Println("download Create", err)
		return
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("download Get", err)
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("download Copy", err)
		return
	}
}
func main() {
	// list of data-uid
	// https://www.aparat.com/v/7dmXg =====>  7dmXg
	listToken := []string{
		"stm53d5",
		"jyd033q",
		"alk10g8",
		"wbljns8",
		"pyo8h24",
		"fpk58s8",
		"elr27r7",
		"agn73az",
		"okcqvyu",
		"qhh7613",
		"pmgl8zo",
		"ycz99o1",
		"bcb3feh",
		"gneff1y",
		"tcac9xa",
		"rkd5e1m",
	}

	for i := range listToken {
		fmt.Println(listToken[i])
		res, err := http.Get("https://www.aparat.com/api/fa/v1/video/video/show/videohash/" + listToken[i] + "?pr=1&mf=1")
		if err != nil {
			fmt.Println("Get", err)
			continue
		}
		defer res.Body.Close()
		var item response
		err = json.NewDecoder(res.Body).Decode(&item)
		if err != nil {
			fmt.Println("NewDecoder", err)
			continue
		}
		//
		quality := 4 // کیفیت 1 یعنی 1080 2 یعنی 720 3یعنی 480
		if len(item.Data.Attributes.File_link_all) > 0 && len(item.Data.Attributes.File_link_all[len(item.Data.Attributes.File_link_all)-1].Urls) > 0 {
			l := len(item.Data.Attributes.File_link_all)
			u := len(item.Data.Attributes.File_link_all[len(item.Data.Attributes.File_link_all)-1].Urls)
			url := item.Data.Attributes.File_link_all[l-quality].Urls[u-1]
			if len(url) > 0 {
				// It downloads the best quality
				download(listToken[i], url)
			}
		}
	}
}
