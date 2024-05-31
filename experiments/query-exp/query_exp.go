package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
)

func main() {
	var client *req.Client = req.C()

	resp,e:=client.R().
		Get("https://jisho.org/search/%20%23jlpt-n2%20%23words?page=2")

	if e!=nil {
		panic(e)
	}

	doc,e:=goquery.NewDocumentFromReader(resp.Body)

	if e!=nil {
		panic(e)
	}

	doc.Find("span.text").Each(func(i int, element *goquery.Selection) {
		fmt.Println(strings.Trim(element.Text()," \n"))
	})
}