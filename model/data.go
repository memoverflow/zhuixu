package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const downloadPath = "/Users/lucas/Downloads/zhuixu"
const dataFile = "data/data.json"

type Play struct {
	Url       string `json:"url"`
	File      string `json:"file"`
	TrackName string `json:"trackName"`
	Pid       int    `json:"pid"`
}

type Site struct {
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Announcer string  `json:"announcer"`
	Limit     int     `json:"limit"`
	PlayList  []*Play `json:"playlist"`
}

func RetrieveBooks() (*Site, error) {
	file, err := os.Open(dataFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var site *Site

	err = json.NewDecoder(file).Decode(&site)

	site.Announcer = decodeUnicode(site.Announcer)
	site.Author = decodeUnicode(site.Author)
	site.Title = decodeUnicode(site.Title)

	for _, p := range site.PlayList {
		p.TrackName = decodeUnicode(p.TrackName)
		p.File = decodeUrl((p.File))
	}

	return site, err
}

func decodeUnicode(encodeText string) string {
	return fmt.Sprintf(encodeText)
}

func decodeUrl(url string) string {
	var result = ""
	strs := strings.Split(url, "*")

	for _, s := range strs {
		i, _ := strconv.Atoi(s)
		dec := fmt.Sprintf("%c", i)
		result += dec
	}
	return result
}

func DownloadFile(p *Play) error {

	fmt.Println("开始下载第" + strconv.Itoa(p.Pid) + "个: " + p.TrackName)

	exts := strings.Split(p.File, ".")
	fmt.Println(p.File)

	extension := exts[len(exts)-1]

	fileName := downloadPath + "/第" + strconv.Itoa(p.Pid) + "章 " + p.TrackName + "." + extension

	req, err := http.NewRequest("GET", p.File, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Create the file
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Println("完成第" + strconv.Itoa(p.Pid) + "个: " + p.TrackName + "的下载")
	return err

}
