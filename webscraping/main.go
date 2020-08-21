package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func main() {

	url := "https://fortune.yahoo.co.jp/12astro/20190922/scorpio.html"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	det := chardet.NewTextDetector()

	detResult, err := det.DetectBest(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(detResult.Charset)

	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detResult.Charset, bReader)

	doc, _ := goquery.NewDocumentFromReader(reader)

	rslt := doc.Find("title").Text()

	fmt.Println(rslt)

}
